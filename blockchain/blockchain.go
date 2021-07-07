package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type BlockChain interface {
	CreateBlock(proof int64, previousHash string) *block
	ProofOfWork(prevProof int64) int64
	HashBlock(bl *block) string
	GetLastBlock() *block
	GetChain() []*block
}

type block struct {
	Index        int64  `json:"index"`
	Nonce        int64  `json:"nonce"`
	Timestamp    int64  `json:"timestamp"`
	PreviousHash string `json:"previousHash"`
}

func NewBlockChain() BlockChain {
	b := &blockChain{}
	b.CreateBlock(1, "0")

	return b
}

type blockChain struct {
	chain []*block
}

func (b *blockChain) CreateBlock(proof int64, previousHash string) *block {
	b.chain = append(b.chain, &block{
		Index:        int64(len(b.chain) + 1),
		Timestamp:    time.Now().Unix(),
		Nonce:        proof,
		PreviousHash: previousHash,
	})

	return b.GetLastBlock()
}

func (b *blockChain) ProofOfWork(prevProof int64) int64 {
	for newProof := int64(1); ; newProof++ {
		str := fmt.Sprintf("%d", (newProof*newProof)-(prevProof*prevProof))
		hasher := sha256.New()
		hasher.Write([]byte(str))
		hash := hex.EncodeToString(hasher.Sum(nil))
		if strings.HasPrefix(hash, "0000") {
			return newProof
		}
	}
}

func (b *blockChain) HashBlock(bl *block) string {
	hasher := sha256.New()
	bytes, _ := json.Marshal(bl)
	hasher.Write(bytes)

	return hex.EncodeToString(hasher.Sum(nil))
}

func (b *blockChain) GetChain() []*block {
	return b.chain
}

func (b *blockChain) GetLastBlock() *block {
	return b.chain[len(b.chain)-1]
}

package blockchain

import "time"

type BlockChain interface {
	CreateBlock(proof, previousHash string) *block
	GetLastBlock() *block
	GetChain() []*block
}

type block struct {
	index        int64
	nonce        string
	timestamp    int64
	previousHash string
}

func NewBlockChain() BlockChain {
	b := &blockChain{}
	b.CreateBlock("1", "0")

	return b
}

type blockChain struct {
	chain []*block
}

func (b *blockChain) CreateBlock(proof, previousHash string) *block {
	b.chain = append(b.chain, &block{
		index:        int64(len(b.chain) + 1),
		timestamp:    time.Now().Unix(),
		nonce:        proof,
		previousHash: previousHash,
	})

	return b.GetLastBlock()
}

func (b *blockChain) GetChain() []*block {
	return b.chain
}

func (b *blockChain) GetLastBlock() *block {
	return b.chain[len(b.chain)-1]
}

package handler

import (
	"encoding/json"
	"net/http"

	"basic-blockchain/blockchain"
)

func MineBlock(b blockchain.BlockChain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lastBlock := b.GetLastBlock()
		block := b.CreateBlock(b.ProofOfWork(lastBlock.Nonce), b.HashBlock(lastBlock))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(block)
	}
}

func GetChain(b blockchain.BlockChain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(b.GetChain())
	}
}

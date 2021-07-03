package main

import (
	"fmt"

	"basic-blockchain/blockchain"
)

func main() {
	blockChain := blockchain.NewBlockChain()
	fmt.Printf("%+v", blockChain.GetLastBlock())
}
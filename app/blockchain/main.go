package main

import (
	"fmt"
	"net/http"

	"basic-blockchain/blockchain"
	"basic-blockchain/router"
)

func main() {
	port := 9000
	blockChain := blockchain.NewBlockChain()
	r := router.NewRouter(blockChain)

	fmt.Printf("Running server on port: %d\n", port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

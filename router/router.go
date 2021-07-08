package router

import (
	"github.com/gorilla/mux"

	"basic-blockchain/blockchain"
	"basic-blockchain/handler"
)

func NewRouter(b blockchain.BlockChain) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/mine", handler.MineBlock(b)).Methods("GET")
	router.HandleFunc("/chain", handler.GetChain(b)).Methods("GET")

	return router
}

package main

import (
	"io"
	"log"
	"mad-blocks/block"
	"mad-blocks/wallet"
	"net/http"
	"strconv"
)

type BlockchainServer struct {
	port uint16
}

var cache map[string]*block.BlockChain = make(map[string]*block.BlockChain)

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

func (bcs *BlockchainServer) GetBlockchain() *block.BlockChain {
	bc, ok := cache["blockchain"]
	if !ok {
		minerWallet := wallet.NewWallet()
		bc = block.NewBlockChain(minerWallet.Address(), bcs.Port())
		cache["blockchain"] = bc
	}

	return bc
}

func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		bc := bcs.GetBlockchain()
		m, _ := bc.MarshalJSON()
		// fmt.Println(string(m[:]))
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(m[:]))
	default:
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", bcs.GetChain)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa((int(bcs.Port()))), nil))
}

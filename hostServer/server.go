package main

import (
	"encoding/json"
	"io"
	"log"
	"mad-blocks/block"
	"mad-blocks/utils"
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

func (bcs *BlockchainServer) Transaction(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		transactions := bc.TransactionPool()
		m, _ := json.Marshal(struct {
			Transactions []*block.Transaction `json:"transactions"`
			Length       int                  `json:"length"`
		}{
			Transactions: transactions,
			Length:       len(transactions),
		})
		io.WriteString(w, string(m[:]))

	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t block.TransactionRequest
		if err := decoder.Decode(&t); err != nil {
			log.Println("ERROR: Cannot Decode TransactionRequest")
			io.WriteString(w, utils.JsonStatus("fail"))
			return
		}
		if !t.Validate() {
			log.Println("ERROR: Invalid TransactionRequest")
			io.WriteString(w, utils.JsonStatus("fail"))
			return
		}
		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		signature := utils.SignatureFromString(*t.Signature)
		bc := bcs.GetBlockchain()
		isCreated := bc.CreateTransaction(*t.SenderAddress, *t.RecipientAddress, *t.Value, publicKey, signature)
		w.Header().Add("Content-Type", "application/json")
		var m []byte
		if !isCreated {
			w.WriteHeader(http.StatusBadRequest)
			m = []byte(utils.JsonStatus("fail"))
		} else {
			w.WriteHeader(http.StatusCreated)
			m = []byte(utils.JsonStatus("success"))
		}
		io.WriteString(w, string(m[:]))

	default:
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (bcs *BlockchainServer) Mine(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		bc := bcs.GetBlockchain()
		isMined := bc.Mining()

		var m []byte
		if !isMined {
			w.WriteHeader(http.StatusBadRequest)
			m = []byte(utils.JsonStatus("fail"))
		} else {
			w.WriteHeader(http.StatusCreated)
			m = []byte(utils.JsonStatus("success"))
		}
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(m[:]))

	default:
		log.Println("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (bcs *BlockchainServer) StartMining(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		bc := bcs.GetBlockchain()
		bc.StartMining()

		w.WriteHeader(http.StatusCreated)
		m := utils.JsonStatus("success")
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(m[:]))

	default:
		log.Println("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (bcs *BlockchainServer) Amount(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		address := req.URL.Query().Get("address")
		amount := bcs.GetBlockchain().BalanceOf(address)
		ar := &block.AmountResponse{Amount: amount}
		m, _ := ar.MarshalJSON()
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(m[:]))

	default:
		log.Println("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", bcs.GetChain)
	http.HandleFunc("/mine", bcs.Mine)
	http.HandleFunc("/amount", bcs.Amount)
	http.HandleFunc("/mine/start", bcs.StartMining)
	http.HandleFunc("/transactions", bcs.Transaction)
	log.Fatal(http.ListenAndServe("127.0.0.1:"+strconv.Itoa((int(bcs.Port()))), nil))
}

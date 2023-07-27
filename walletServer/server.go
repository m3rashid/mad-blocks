package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"mad-blocks/block"
	"mad-blocks/utils"
	"mad-blocks/wallet"
	"net/http"
	"path"
	"strconv"
)

const (
	// TEMP_DIR = "walletServer/templates"
	TEMP_DIR = "templates"
)

type WalletServer struct {
	port    uint16
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(TEMP_DIR, "index.html"))
		t.Execute(w, "")
	default:
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := wallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t wallet.TransactionRequest
		if err := decoder.Decode(&t); err != nil {
			log.Println("ERROR: Cannot Decode TransactionRequest")
			io.WriteString(w, utils.JsonStatus("Bad Request"))
			return
		}
		if !t.Validate() {
			log.Println("ERROR: Invalid TransactionRequest")
			io.WriteString(w, utils.JsonStatus("fail"))
			return
		}

		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		privatekey := utils.PrivateKeyFromString(*t.SenderPrivateKey, publicKey)
		value, err := strconv.ParseFloat(*t.Value, 32)
		if err != nil {
			log.Println("ERROR: Cannot ParseFloat")
			io.WriteString(w, utils.JsonStatus("fail"))
			return
		}
		value32 := float32(value)
		w.Header().Add("Content-Type", "application/json")
		transaction := wallet.NewTransaction(privatekey, publicKey, *t.SenderAddress, *t.RecipientAddress, value32)
		signature := transaction.GenerateSignature()
		signatureStr := signature.String()

		bt := &block.TransactionRequest{
			SenderAddress:    t.SenderAddress,
			RecipientAddress: t.RecipientAddress,
			SenderPublicKey:  t.SenderPublicKey,
			Signature:        &signatureStr,
			Value:            &value32,
		}
		m, _ := json.Marshal(bt)
		buffer := bytes.NewBuffer(m)
		resp, err := http.Post(ws.Gateway()+"/transactions", "application/json", buffer)
		if err != nil {
			log.Println("ERROR: Cannot Post Transaction due to API error")
			io.WriteString(w, utils.JsonStatus("fail"))
			return
		}

		if resp.StatusCode == http.StatusCreated {
			io.WriteString(w, utils.JsonStatus("success"))
			return
		} else {
			fmt.Println(resp.Body)
			log.Println("ERROR: Cannot Create Transaction on API")
			io.WriteString(w, utils.JsonStatus("fail"))
			return
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}

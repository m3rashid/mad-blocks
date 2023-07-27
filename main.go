package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Nonce        int      `json:"nonce"`
	PreviousHash [32]byte `json:"previousHash"`
	Transactions []string `json:"transactions"`
	Timestamp    int64    `json:"timestamp"`
}

func newBlock(nonce int, previousHash [32]byte) *Block {
	return &Block{
		Nonce:        nonce,
		PreviousHash: previousHash,
		Transactions: []string{},
		Timestamp:    time.Now().UnixNano(),
	}
}

func (b *Block) hash() [32]byte {
	m, _ := json.Marshal(b)

	return sha256.Sum256(m)
}

func (b *Block) Print() {
	fmt.Printf("nonce:\t\t%d\n", b.Nonce)
	fmt.Printf("previousHash:\t%s\n", fmt.Sprintf("%x", b.PreviousHash))
	fmt.Printf("transactions:\t%s\n", b.Transactions)
	fmt.Printf("timestamp:\t%d\n", b.Timestamp)
}

type BlockChain struct {
	Chain           []*Block
	TransactionPool []string
}

func (bc *BlockChain) createBlock(nonce int, previousHash [32]byte) *Block {
	b := newBlock(nonce, previousHash)
	bc.Chain = append(bc.Chain, b)
	return b
}

func (bc *BlockChain) Print() {
	for i, block := range bc.Chain {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 20), i+1, strings.Repeat("=", 20))
		block.Print()
	}
	fmt.Println()
	fmt.Println()
}

func NewBlockChain() *BlockChain {
	block := &Block{}
	blockChain := &BlockChain{}
	blockChain.createBlock(1, block.hash())
	return blockChain
}

func (bc *BlockChain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

type Transaction struct {
	SenderAddress    string  `json:"senderAddress"`
	RecipientAddress string  `json:"recipientAddress"`
	Value            float32 `json:"value"`
}

func main() {
	blockChain := NewBlockChain()
	blockChain.Print()

	previousHash := blockChain.LastBlock().hash()
	blockChain.createBlock(2, previousHash)
	blockChain.Print()
}

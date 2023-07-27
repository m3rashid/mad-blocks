package main

import (
	"fmt"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	transactions []string
	timestamp    int64
}

func newBlock(nonce int, previousHash string) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		transactions: []string{},
		timestamp:    time.Now().UnixNano(),
	}
}

func (b *Block) Print() {
	fmt.Printf("nonce:\t\t%d\n", b.nonce)
	fmt.Printf("previousHash:\t%s\n", b.previousHash)
	fmt.Printf("transactions:\t%s\n", b.transactions)
	fmt.Printf("timestamp:\t%d\n", b.timestamp)
}

type BlockChain struct {
	chain           []*Block
	transactionPool []string
}

func (bc *BlockChain) createBlock(nonce int, previousHash string) *Block {
	b := newBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *BlockChain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 20), i, strings.Repeat("=", 20))
		block.Print()
	}
	fmt.Println()
	fmt.Println()
}

func NewBlockChain() *BlockChain {
	bc := &BlockChain{}
	bc.createBlock(1, "init_hash")
	return bc
}

func main() {
	blockChain := NewBlockChain()
	blockChain.Print()
	blockChain.createBlock(2, "hash_1")
	blockChain.Print()
}

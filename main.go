package main

import (
	"fmt"
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
	fmt.Printf("nonce:\t%d\n", b.nonce)
	fmt.Printf("previousHash:\t%s\n", b.previousHash)
	fmt.Printf("transactions:\t%s\n", b.transactions)
	fmt.Printf("timestamp:\t%d\n", b.timestamp)
}

func main() {
	b := newBlock(1, "hash")
	b.Print()
}

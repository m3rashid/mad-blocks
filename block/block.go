package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Nonce        int            `json:"nonce"`
	PreviousHash [32]byte       `json:"previousHash"`
	Transactions []*Transaction `json:"transactions"`
	Timestamp    int64          `json:"timestamp"`
}

func newBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		Nonce:        nonce,
		PreviousHash: previousHash,
		Transactions: transactions,
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
	fmt.Printf("timestamp:\t%d\n", b.Timestamp)
	// fmt.Printf("transactions:\t%s\n", b.Transactions)
	for _, transaction := range b.Transactions {
		transaction.Print()
	}
}

package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	transactions []*Transaction
	timestamp    int64
}

func newBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		transactions: transactions,
		timestamp:    time.Now().UnixNano(),
	}
}

func (b *Block) Print() {
	fmt.Printf("nonce:\t\t%d\n", b.nonce)
	fmt.Printf("previousHash:\t%s\n", fmt.Sprintf("%x", b.previousHash))
	fmt.Printf("timestamp:\t%d\n", b.timestamp)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := b.MarshalJSON()
	return sha256.Sum256(m)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash string         `json:"previousHash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: fmt.Sprintf("%x", b.previousHash),
		Transactions: b.transactions,
	})
}

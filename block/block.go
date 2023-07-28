package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

func newBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	var b Block
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions

	return &b
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

func (b *Block) UnMarshalJSON(data []byte) error {
	var previousHash string
	x := &struct {
		TimeStamp    *int64          `json:"timestamp"`
		Nonce        *int            `json:"nonce"`
		PreviousHash *string         `json:"previousHash"`
		Transactions *[]*Transaction `json:"transactions"`
	}{
		TimeStamp:    &b.timestamp,
		Nonce:        &b.nonce,
		PreviousHash: &previousHash,
		Transactions: &b.transactions,
	}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	ph, _ := hex.DecodeString(*x.PreviousHash)
	copy(b.previousHash[:], ph[:32])
	return nil
}

func (b *Block) PreviousHash() [32]byte {
	return b.previousHash
}

func (b *Block) Nonce() int {
	return b.nonce
}

func (b *Block) Transactions() []*Transaction {
	return b.transactions
}

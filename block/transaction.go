package block

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"mad-blocks/utils"
	"strings"
)

type Transaction struct {
	sender    string
	recipient string
	value     float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (bc *BlockChain) CreateTransaction(
	sender string,
	recipient string,
	value float32,
	publicKey *ecdsa.PublicKey,
	s *utils.Signature,
) bool {
	isTransactionValid := bc.AddTransaction(sender, recipient, value, publicKey, s)
	return isTransactionValid
}

func (bc *BlockChain) AddTransaction(
	sender string,
	recipient string,
	value float32,
	publicKey *ecdsa.PublicKey,
	s *utils.Signature,
) bool {
	t := NewTransaction(sender, recipient, value)

	if sender == utils.MINING_SENDER {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	}

	if bc.BalanceOf(sender) < value {
		log.Println("Insufficient Balance")
		return false
	}

	if bc.VerifyTransactionSignature(publicKey, s, t) {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	} else {
		log.Println("ERROR: Cannot Verify Transaction")
	}

	return false
}

func (t *Transaction) Print() {
	fmt.Println(strings.Repeat("-", 102))
	fmt.Printf("Sender: %s, ", t.sender)
	fmt.Printf("Recipient: %s, ", t.recipient)
	fmt.Printf("Value: %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender"`
		Recipient string  `json:"recipient"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.sender,
		Recipient: t.recipient,
		Value:     t.value,
	})
}

func (t *Transaction) UnMarshalJSON(data []byte) error {
	x := &struct {
		Sender    *string  `json:"sender"`
		Recipient *string  `json:"recipient"`
		Value     *float32 `json:"value"`
	}{
		Sender:    &t.sender,
		Recipient: &t.recipient,
		Value:     &t.value,
	}

	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	return nil
}

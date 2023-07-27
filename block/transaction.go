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

func (bc *BlockChain) AddTransaction(
	sender string,
	recipient string,
	value float32,
	senderPublicKey *ecdsa.PublicKey,
	s *utils.Signature,
) bool {
	t := NewTransaction(sender, recipient, value)

	if sender == utils.MINING_SENDER {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	}

	// if bc.BalanceOf(sender) < value {
	// 	log.Println("Insufficient Balance")
	// 	return false
	// }

	if bc.VerifyTransactionSignature(senderPublicKey, s, t) {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	} else {
		log.Println("ERROR: Cannot Verify Transaction")
	}

	return false
}

func (t *Transaction) Print() {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("Sender: %s, ", t.sender)
	fmt.Printf("Recipient: %s, ", t.recipient)
	fmt.Printf("Value: %.1f\n", t.value)
}

func (tr *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    tr.sender,
		Recipient: tr.recipient,
		Value:     tr.value,
	})
}

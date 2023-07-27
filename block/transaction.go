package block

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"mad-blocks/utils"
	"strings"
)

type Transaction struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Value     float32 `json:"value"`
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{
		Sender:    sender,
		Recipient: recipient,
		Value:     value,
	}
}

func (bc *BlockChain) AddTransaction(
	sender string,
	recipient string,
	value float32,
	senderPublicKey *ecdsa.PublicKey,
	signature *utils.Signature,
) bool {
	t := NewTransaction(sender, recipient, value)

	if sender == utils.MINING_SENDER {
		bc.TransactionPool = append(bc.TransactionPool, t)
		return true
	}

	// if bc.BalanceOf(sender) < value {
	// 	log.Println("ERROR: No balance")
	// 	return false
	// }

	if bc.VerifyTransactionSignature(senderPublicKey, signature, t) {
		bc.TransactionPool = append(bc.TransactionPool, t)
		return true
	} else {
		log.Println("Cannot Verify Transaction")
		return false
	}
}

func (t *Transaction) Print() {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("Sender: %s, ", t.Sender)
	fmt.Printf("Recipient: %s, ", t.Recipient)
	fmt.Printf("Value: %.1f\n", t.Value)
}

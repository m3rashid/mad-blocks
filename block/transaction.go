package block

import (
	"fmt"
	"strings"
)

type Transaction struct {
	Sender    string  `json:"senderAddress"`
	Recipient string  `json:"recipientAddress"`
	Value     float32 `json:"value"`
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{
		Sender:    sender,
		Recipient: recipient,
		Value:     value,
	}
}

func (t *Transaction) Print() {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("Sender: %s, ", t.Sender)
	fmt.Printf("Recipient: %s, ", t.Recipient)
	fmt.Printf("Value: %f\n", t.Value)
}

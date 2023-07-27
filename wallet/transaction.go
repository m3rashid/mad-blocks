package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"mad-blocks/utils"
)

type Transaction struct {
	SenderPrivateKey *ecdsa.PrivateKey
	SenderPublicKey  *ecdsa.PublicKey
	Sender           string  `json:"sender"`
	Recipient        string  `json:"recipient"`
	Value            float32 `json:"value"`
}

func (t *Transaction) GenerateSignature() *utils.Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256(m)
	r, s, _ := ecdsa.Sign(rand.Reader, t.SenderPrivateKey, h[:])
	return &utils.Signature{
		R: r,
		S: s,
	}
}

func NewTransaction(
	senderPrivateKey *ecdsa.PrivateKey,
	senderPublicKey *ecdsa.PublicKey,
	sender string,
	recipient string,
	value float32,
) *Transaction {
	return &Transaction{
		SenderPrivateKey: senderPrivateKey,
		SenderPublicKey:  senderPublicKey,
		Sender:           sender,
		Recipient:        recipient,
		Value:            value,
	}
}

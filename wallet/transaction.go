package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"mad-blocks/utils"
)

type Transaction struct {
	senderPrivateKey *ecdsa.PrivateKey
	senderPublicKey  *ecdsa.PublicKey
	sender           string
	recipient        string
	value            float32
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

func (t *Transaction) GenerateSignature() *utils.Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256(m)
	r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
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
		senderPrivateKey: senderPrivateKey,
		senderPublicKey:  senderPublicKey,
		sender:           sender,
		recipient:        recipient,
		value:            value,
	}
}

package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"mad-blocks/utils"
	"strings"
)

type Transaction struct {
	SenderPrivateKey *ecdsa.PrivateKey
	SenderPublicKey  *ecdsa.PublicKey
	Sender           string  `json:"sender"`
	Recipient        string  `json:"recipient"`
	Value            float32 `json:"value"`
}

func (t *Transaction) Print() {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("Sender: %s, ", t.Sender)
	fmt.Printf("Recipient: %s, ", t.Recipient)
	fmt.Printf("Value: %f\n", t.Value)
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

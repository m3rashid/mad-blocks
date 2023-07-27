package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    string
}

func NewWallet() *Wallet {
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.PrivateKey = privateKey

	w.PublicKey = &w.PrivateKey.PublicKey

	// 2. Perform SHA256 hashing on the public key (32 bytes)
	h2 := sha256.New()
	h2.Write(w.PublicKey.X.Bytes())
	h2.Write(w.PublicKey.Y.Bytes())
	digest2 := h2.Sum(nil)

	// 3. Perform RIPEMD160 hashing on the result of SHA256 (20 bytes)
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)

	// 4. Add version byte in front of RIPEMD-160 hash (0x00 for Main Network)
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])

	// 5. Perform SHA256 hash on the extended RIPEMD-160 result
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)

	// 6. Perform SHA256 hash on the result of the previous SHA256 hash
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)

	// 7. Take the first 4 bytes of the second SHA256 hash. This is the address checksum
	checksum := digest6[:4]

	// 8. Add the 4 checksum bytes from point 7 at the end of extended RIPEMD-160 hash from point 4.
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], checksum[:])

	// 9. Convert the string from byte to base58
	address := base58.Encode(dc8)
	w.Address = address

	return w
}

func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.PublicKey.X.Bytes(), w.PublicKey.Y.Bytes())
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.PrivateKey.D.Bytes())
}

func (w *Wallet) NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, sender string, recipient string, value float32) *Transaction {
	return &Transaction{
		SenderPrivateKey: privateKey,
		SenderPublicKey:  publicKey,
		Sender:           sender,
		Recipient:        recipient,
		Value:            value,
	}
}

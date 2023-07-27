package main

import (
	"fmt"
	"log"
	"mad-blocks/wallet"
)

func init() {
	log.SetPrefix("MadBlocks: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Printf("Public Key: %s\n", w.PublicKeyStr())
	fmt.Printf("Private Key: %s\n", w.PrivateKeyStr())

	fmt.Println("Address: ", w.Address)
	t := w.NewTransaction(w.PrivateKey, w.PublicKey, w.Address, "recipient", 1.0)
	fmt.Printf("Signature %s\n", t.GenerateSignature())
}

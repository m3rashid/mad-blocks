package main

import (
	"fmt"
	"log"
	"mad-blocks/utils"
	"mad-blocks/wallet"
)

func init() {
	log.SetPrefix("MadBlocks: ")
}

func main() {
	params := utils.DefaultFuncParams{
		Verbose: false,
	}

	w := wallet.NewWallet()
	fmt.Printf("Public Key: %s\n", w.PublicKeyStr())
	fmt.Printf("Private Key: %s\n", w.PrivateKeyStr())

	fmt.Println("Address: ", w.Address)

	fmt.Println(params)
}

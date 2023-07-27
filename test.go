package main

import (
	"fmt"
	"log"
	"mad-blocks/block"
	"mad-blocks/wallet"
)

func init() {
	log.SetPrefix("MadBlocks: ")
}

func main() {
	userA := wallet.NewWallet()
	userB := wallet.NewWallet()
	miner := wallet.NewWallet()

	t := wallet.NewTransaction(userA.PrivateKey(), userA.PublicKey(), userA.Address(), userB.Address(), 1.0)

	blockchain := block.NewBlockChain(miner.Address(), 8080)
	isAdded := blockchain.AddTransaction(userA.Address(), userB.Address(), 1.0, userA.PublicKey(), t.GenerateSignature())
	fmt.Println("Added: ", isAdded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Println("UserA: ", blockchain.BalanceOf(userA.Address()), "\t", userA.Address())
	fmt.Println("UserB: ", blockchain.BalanceOf(userB.Address()), "\t", userB.Address())
	fmt.Println("Miner: ", blockchain.BalanceOf(miner.Address()), "\t", miner.Address())
}

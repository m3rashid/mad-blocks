package main

import (
	"fmt"
	"log"
	"mad-blocks/block"
	"mad-blocks/utils"
	"mad-blocks/wallet"
)

func init() {
	log.SetPrefix("MadBlocks: ")
}

func main() {
	userA := wallet.NewWallet()
	userB := wallet.NewWallet()
	miner := wallet.NewWallet()

	t := wallet.NewTransaction(userA.PrivateKey, userA.PublicKey, userA.Address, userB.Address, 1.0)

	blockchain := block.NewBlockChain(miner.Address)
	isAdded := blockchain.AddTransaction(userA.Address, userB.Address, 1.0, userA.PublicKey, t.GenerateSignature())
	fmt.Println("Added: ", isAdded)

	blockchain.Mining(utils.DefaultFuncParams)
	blockchain.Print()

	fmt.Println("UserA: ", blockchain.BalanceOf(userA.Address))
	fmt.Println("UserB: ", blockchain.BalanceOf(userB.Address))
	fmt.Println("Miner: ", blockchain.BalanceOf(miner.Address))
}

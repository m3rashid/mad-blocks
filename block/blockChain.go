package block

import (
	"fmt"
	"mad-blocks/utils"
	"strings"
)

type BlockChain struct {
	Chain           []*Block       `json:"chain"`
	TransactionPool []*Transaction `json:"transactionPool"`
	Address         string         `json:"address"`
}

func (bc *BlockChain) createBlock(nonce int, previousHash [32]byte) *Block {
	b := newBlock(nonce, previousHash, bc.TransactionPool)
	bc.Chain = append(bc.Chain, b)
	bc.TransactionPool = []*Transaction{}
	return b
}

func (bc *BlockChain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.TransactionPool = append(bc.TransactionPool, t)
}

func (bc *BlockChain) Print() {
	for i, b := range bc.Chain {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 20), i+1, strings.Repeat("=", 20))
		b.Print()
	}
	fmt.Println()
	fmt.Println()
}

func NewBlockChain(blockChainAddress string) *BlockChain {
	b := &Block{}
	bc := new(BlockChain)
	bc.Address = blockChainAddress
	bc.createBlock(1, b.hash())
	return bc
}

func (bc *BlockChain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *BlockChain) Mining(defaultParams utils.DefaultFuncParamsType) bool {
	bc.AddTransaction(utils.MINING_SENDER, bc.Address, utils.MINING_REWARD)
	nonce := bc.ProofOfWork(defaultParams)
	previousHash := bc.LastBlock().hash()
	bc.createBlock(nonce, previousHash)
	return true
}

func (bc *BlockChain) CalculateTotalAmount(address string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.Chain {
		for _, t := range b.Transactions {
			value := t.Value
			if address == t.Recipient {
				totalAmount += value
			}

			if address == t.Sender {
				totalAmount -= value
			}
		}
	}

	return totalAmount
}

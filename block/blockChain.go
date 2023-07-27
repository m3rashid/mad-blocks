package block

import (
	"fmt"
	"mad-blocks/utils"
	"strings"
)

type BlockChain struct {
	TransactionPool []*Transaction `json:"transactionPool"`
	Chain           []*Block       `json:"chain"`
	Address         string         `json:"address"`
}

func (bc *BlockChain) createBlock(nonce int, previousHash [32]byte) *Block {
	b := newBlock(nonce, previousHash, bc.TransactionPool)
	bc.Chain = append(bc.Chain, b)
	bc.TransactionPool = []*Transaction{}
	return b
}

func NewBlockChain(blockChainAddress string) *BlockChain {
	b := &Block{}
	bc := new(BlockChain)
	bc.Address = blockChainAddress
	bc.createBlock(1, b.hash())
	return bc
}

func (bc *BlockChain) Print() {
	for i, b := range bc.Chain {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 20), i+1, strings.Repeat("=", 20))
		b.Print()
	}
	fmt.Println()
	fmt.Println()
}

func (bc *BlockChain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}
func (bc *BlockChain) AddBlock() *Block {
	lb := bc.LastBlock()
	b := bc.createBlock(0, lb.hash())
	return b
}

func (bc *BlockChain) copyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, transaction := range bc.TransactionPool {
		transactions = append(transactions, NewTransaction(
			transaction.Sender,
			transaction.Recipient,
			transaction.Value,
		))
	}
	return transactions
}

func (bc *BlockChain) Mining(defaultParams utils.DefaultFuncParamsType) bool {
	bc.AddTransaction(utils.MINING_SENDER, bc.Address, utils.MINING_REWARD, nil, nil)
	nonce := bc.ProofOfWork(defaultParams)
	previousHash := bc.LastBlock().hash()
	bc.createBlock(nonce, previousHash)
	return true
}

func (bc *BlockChain) BalanceOf(address string) float32 {
	var balance float32 = 0.0
	for _, b := range bc.Chain {
		for _, tr := range b.Transactions {
			if tr.Recipient == address {
				balance = balance + tr.Value
			}

			if tr.Sender == address {
				balance = balance - tr.Value
			}
		}
	}

	return balance
}

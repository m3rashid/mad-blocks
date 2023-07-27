package block

import (
	"encoding/json"
	"fmt"
	"mad-blocks/utils"
	"strings"
)

type BlockChain struct {
	transactionPool []*Transaction
	chain           []*Block
	address         string
}

func (bc *BlockChain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Blocks []*Block `json:"chains"`
	}{
		Blocks: bc.chain,
	})
}

func (bc *BlockChain) createBlock(nonce int, previousHash [32]byte) *Block {
	b := newBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func NewBlockChain(blockChainAddress string) *BlockChain {
	b := &Block{}
	bc := new(BlockChain)
	bc.address = blockChainAddress
	bc.createBlock(0, b.hash())
	return bc
}

func (bc *BlockChain) Print() {
	for i, b := range bc.chain {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 20), i+1, strings.Repeat("=", 20))
		b.Print()
	}
	fmt.Println()
	fmt.Println()
}

func (bc *BlockChain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}
func (bc *BlockChain) AddBlock() *Block {
	lb := bc.LastBlock()
	b := bc.createBlock(0, lb.hash())
	return b
}

func (bc *BlockChain) copyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, transaction := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(
			transaction.sender,
			transaction.recipient,
			transaction.value,
		))
	}
	return transactions
}

func (bc *BlockChain) Mining(defaultParams utils.DefaultFuncParamsType) bool {
	bc.AddTransaction(utils.MINING_SENDER, bc.address, utils.MINING_REWARD, nil, nil)
	nonce := bc.ProofOfWork(defaultParams)
	previousHash := bc.LastBlock().hash()
	bc.createBlock(nonce, previousHash)
	return true
}

func (bc *BlockChain) BalanceOf(address string) float32 {
	var balance float32 = 0.0
	for _, b := range bc.chain {
		for _, tr := range b.transactions {
			if tr.recipient == address {
				balance = balance + tr.value
			}

			if tr.sender == address {
				balance = balance - tr.value
			}
		}
	}

	return balance
}

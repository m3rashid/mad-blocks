package block

import (
	"fmt"
	"mad-blocks/utils"
	"strings"
	"time"
)

func (bc *BlockChain) copyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, transaction := range bc.TransactionPool {
		transactions = append(transactions, NewTransaction(transaction.Sender, transaction.Recipient, transaction.Value))
	}
	return transactions
}

func (bc *BlockChain) validProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int, defaultParams utils.DefaultFuncParams) bool {
	zeroes := strings.Repeat("0", difficulty)
	guessBlock := Block{
		Timestamp:    0,
		Nonce:        nonce,
		PreviousHash: previousHash,
		Transactions: transactions,
	}
	guessHash := fmt.Sprintf("%x", guessBlock.hash())
	matched := guessHash[:difficulty] == zeroes
	if matched && defaultParams.Verbose {
		fmt.Printf("Matched HASH: %s\n", guessHash)
	}
	return matched
}

func (bc *BlockChain) ProofOfWork(defaultParams utils.DefaultFuncParams) int {
	startTime := time.Now()
	transactions := bc.copyTransactionPool()
	previousHash := bc.LastBlock().hash()
	nonce := 0
	for !bc.validProof(nonce, previousHash, transactions, utils.MINING_DIFFICULTY, defaultParams) {
		nonce += 1
	}
	if defaultParams.Verbose {
		fmt.Printf("Proof Calculation Took : %s\n\n", time.Since(startTime))
	}
	return nonce
}

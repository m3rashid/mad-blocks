package block

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"mad-blocks/utils"
	"strings"
	"time"
)

func (bc *BlockChain) VerifyTransactionSignature(
	senderPublicKey *ecdsa.PublicKey,
	signature *utils.Signature,
	transaction *Transaction,
) bool {
	m, _ := json.Marshal(transaction)
	hash := sha256.Sum256(m)
	return ecdsa.Verify(senderPublicKey, hash[:], signature.R, signature.S)
}

func (bc *BlockChain) validProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int, defaultParams utils.DefaultFuncParamsType) bool {
	zeroes := strings.Repeat("0", difficulty)
	guessBlock := Block{
		timestamp:    0,
		nonce:        nonce,
		previousHash: previousHash,
		transactions: transactions,
	}
	guessHash := fmt.Sprintf("%x", guessBlock.hash())
	matched := guessHash[:difficulty] == zeroes
	if matched && defaultParams.Verbose {
		fmt.Printf("Matched HASH: %s\n", guessHash)
	}
	return matched
}

func (bc *BlockChain) ProofOfWork(defaultParams utils.DefaultFuncParamsType) int {
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

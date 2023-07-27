package block

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"mad-blocks/utils"
	"strings"
)

func (bc *BlockChain) VerifyTransactionSignature(
	senderPublicKey *ecdsa.PublicKey,
	s *utils.Signature,
	t *Transaction,
) bool {
	m, _ := json.Marshal(t)
	hash := sha256.Sum256(m)
	return ecdsa.Verify(senderPublicKey, hash[:], s.R, s.S)
}

func (bc *BlockChain) validProof(
	nonce int,
	previousHash [32]byte,
	transactions []*Transaction,
	difficulty int,
) bool {
	zeroes := strings.Repeat("0", difficulty)
	guessBlock := Block{
		timestamp:    0,
		nonce:        nonce,
		previousHash: previousHash,
		transactions: transactions,
	}
	guessHash := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHash[:difficulty] == zeroes
}

func (bc *BlockChain) ProofOfWork() int {
	nonce := 0
	transactions := bc.copyTransactionPool()
	previousHash := bc.LastBlock().Hash()

	for !bc.validProof(nonce, previousHash, transactions, utils.MINING_DIFFICULTY) {
		nonce = nonce + 1
	}

	return nonce
}

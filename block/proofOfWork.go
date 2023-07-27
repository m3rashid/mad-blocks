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
	h := sha256.Sum256(m)
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}

func (bc *BlockChain) ValidProof(
	nonce int,
	previousHash [32]byte,
	transactions []*Transaction,
	difficulty int,
) bool {
	zeroes := strings.Repeat("0", difficulty)
	guessBlock := Block{nonce, previousHash, 0, transactions}
	guessHash := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHash[:difficulty] == zeroes
}

func (bc *BlockChain) ProofOfWork() int {
	nonce := 0
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()

	for !bc.ValidProof(nonce, previousHash, transactions, utils.MINING_DIFFICULTY) {
		nonce = nonce + 1
	}

	return nonce
}

package block

import (
	"encoding/json"
	"fmt"
	"log"
	"mad-blocks/utils"
	"net/http"
)

func (bc *BlockChain) ValidChain(chain []*Block) bool {
	prevBlock := chain[0]
	currentIndex := 1
	for currentIndex < len(chain) {
		b := chain[currentIndex]
		if b.previousHash != prevBlock.Hash() {
			return false
		}

		if bc.ValidProof(b.Nonce(), b.PreviousHash(), b.Transactions(), utils.MINING_DIFFICULTY) {
			return false
		}

		prevBlock = b
		currentIndex += 1
	}
	return true
}

func (bc *BlockChain) ResolveConflicts() bool {
	var longestChain []*Block = nil
	maxLength := len(bc.chain)

	for _, n := range bc.neighbors {
		endpoint := fmt.Sprintf("http://%s/chain", n)
		resp, _ := http.Get(endpoint)
		if resp.StatusCode == http.StatusOK {
			var bcRes BlockChain
			decoder := json.NewDecoder(resp.Body)
			_ = decoder.Decode(&bcRes)

			chain := bcRes.Chain()
			if len(chain) > maxLength && bc.ValidChain(chain) {
				maxLength = len(chain)
				longestChain = chain
			}
		}
	}

	if longestChain != nil {
		bc.chain = longestChain
		log.Println("Resolve Conflicts replaced")
		return true
	}

	log.Println("Resolve conflicts not replaced")
	return false
}

package block

import (
	"encoding/json"
	"fmt"
	"log"
	"mad-blocks/utils"
	"strings"
	"sync"
	"time"
)

type BlockChain struct {
	transactionPool []*Transaction
	chain           []*Block
	address         string
	port            uint16
	mux             sync.Mutex
	neighbors       []string
	muxNeighbors    sync.Mutex
}

func (bc *BlockChain) Run() {
	bc.StartMining()
	bc.StartSyncNeighbors()
}

func (bc *BlockChain) SetNeighbors() {
	bc.neighbors = utils.FindNeighbors(
		utils.GetHost(),
		bc.port,
		utils.NEIGHBOR_IP_RANGE_START,
		utils.NEIGHBOR_IP_RANGE_END,
		utils.BLOCKCHAIN_PORT_RANGE_START,
		utils.BLOCKCHAIN_PORT_RANGE_END,
	)
	log.Printf("%v", bc.neighbors)
}

func (bc *BlockChain) SyncNeighbors() {
	bc.muxNeighbors.Lock()
	defer bc.muxNeighbors.Unlock()

	bc.SetNeighbors()
}

func (bc *BlockChain) StartSyncNeighbors() {
	bc.SyncNeighbors()
	_ = time.AfterFunc(time.Second*utils.NEIGHBORS_SYNC_TIME_SET_SECONDS, bc.SyncNeighbors)
}

func (bc *BlockChain) TransactionPool() []*Transaction {
	return bc.transactionPool
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := newBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func NewBlockChain(address string, port uint16) *BlockChain {
	b := &Block{}
	bc := new(BlockChain)
	bc.address = address
	bc.port = port
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *BlockChain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *BlockChain) AddBlock() *Block {
	lb := bc.LastBlock()
	b := bc.CreateBlock(0, lb.Hash())
	return b
}

func (bc *BlockChain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(t.sender, t.recipient, t.value))
	}
	return transactions
}

func (bc *BlockChain) Mining() bool {
	bc.mux.Lock()
	defer bc.mux.Unlock()

	if len(bc.transactionPool) == 0 {
		fmt.Println("No Transactions to Mine")
		return false
	}

	bc.AddTransaction(utils.MINING_SENDER, bc.address, utils.MINING_REWARD, nil, nil)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	return true
}

func (bc *BlockChain) StartMining() {
	bc.Mining()
	_ = time.AfterFunc(time.Second*utils.MINING_TIMER_SECONDS, bc.StartMining)
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

func (bc *BlockChain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Blocks []*Block `json:"chains"`
	}{
		Blocks: bc.chain,
	})
}

func (bc *BlockChain) Print() {
	for i, b := range bc.chain {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 20), i+1, strings.Repeat("=", 20))
		b.Print()
	}
	fmt.Println()
	fmt.Println()
}

package main

const (
	MINING_DIFFICULTY = 4
	MINING_SENDER     = "MadBlocks"
	MINING_REWARD     = 1.0
)

func main() {
	userBlockChainAddress := "USER1"
	bc := NewBlockChain(userBlockChainAddress)
	bc.Print()

	bc.AddTransaction("A", "B", 1.0)
	bc.Mining()

	bc.AddTransaction("X", "A", 2.0)
	bc.AddTransaction("E", "D", 2.0)
	bc.Mining()
	bc.Print()
}

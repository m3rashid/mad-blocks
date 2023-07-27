package main

import "fmt"

func main() {
	params := DefaultFuncParams{
		Verbose: false,
	}

	userBlockChainAddress := "USER1"

	bc := NewBlockChain(userBlockChainAddress)
	bc.Print()

	bc.AddTransaction("A", "B", 1.0)
	bc.Mining(params)

	bc.AddTransaction("X", "A", 2.0)
	bc.AddTransaction("E", "D", 2.0)
	bc.Mining(params)
	bc.Print()

	fmt.Printf("A: %.1f\n", bc.CalculateTotalAmount("A"))
	fmt.Printf("B: %.1f\n", bc.CalculateTotalAmount("B"))
	fmt.Printf("X: %.1f\n", bc.CalculateTotalAmount("X"))
	fmt.Printf("E: %.1f\n", bc.CalculateTotalAmount("E"))
	fmt.Printf("D: %.1f\n", bc.CalculateTotalAmount("D"))
}

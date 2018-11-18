package main

import (
	"fmt"
	"phenri00/blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}

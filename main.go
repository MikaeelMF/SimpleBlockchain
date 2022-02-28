package main

import "fmt"

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC to Mik")
	bc.AddBlock("Send 2 BTC to Miner")
	for _, block := range bc.blocks {
		fmt.Printf("Previous Block Hash: %x\n", block.prevBlockHash)
		fmt.Printf("Current Block Hash: %x\n", block.Hash)
		fmt.Printf("Block Data: %s\n", block.Data)
		fmt.Println()
	}
}

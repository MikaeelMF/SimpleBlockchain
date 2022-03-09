package main

import (
	"flag"
	"fmt"
	"os"
)

// Run Cli to execute user commands
func Run(bc *Blockchain) {
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block Data")
	switch os.Args[1] {
	case "addBlock":
		_ = addBlockCmd.Parse(os.Args[2:])
	case "printChain":
		_ = printChainCmd.Parse(os.Args[2:])
	default:
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		addBlock(*addBlockData, bc)
	}

	if printChainCmd.Parsed() {
		printChain(bc)
	}
}

func addBlock(data string, bc *Blockchain) {
	bc.AddBlock(data)
	fmt.Println("Success")
}

func printChain(bc *Blockchain) {
	for {
		block := bc.GetPreviousBlock()
		fmt.Printf("Previous Block Hash is : %x\n", block.GetPreviousBlockHash())
		fmt.Printf("Current Block Hash is : %x\n", block.GetBlockHash())
		fmt.Printf("Block Data is : %s\n", block.GetData())
		fmt.Println()

		if block.GetBlockHeight() == 0 {
			break
		}
	}
}

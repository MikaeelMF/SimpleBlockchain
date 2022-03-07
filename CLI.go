package main

import (
	"flag"
	"fmt"
	"os"
)

type CLI struct {
	bc *Blockchain
}

func (cli *CLI) Run() {
	// cli.validateArgs() WHAT IS THIS???

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block Data")
	switch os.Args[1] {
	case "addBlock":
		_ = addBlockCmd.Parse(os.Args[2:])
	case "printChain":
		_ = printChainCmd.Parse(os.Args[2:])
	default:
		// cli.printUsage() WHAT IS THIS???
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Success")
}

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Previous Block Hash is : %x\n", block.prevBlockHash)
		fmt.Printf("Current Block Hash is : %x\n", block.Hash)
		fmt.Printf("Block Data is : %s\n", block.Data)
		pow := NewProofOfWork(block)
		fmt.Printf("Proof of work is : %t\n", pow.validate())
		fmt.Println()

		if len(block.prevBlockHash) == 0 {
			break
		}
	}
}

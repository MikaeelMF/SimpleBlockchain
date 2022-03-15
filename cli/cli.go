package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MikaeelMF/SimpleBlockchain/transactions"
)

func main() {

	if os.Args[1] == "help" {
		help()
		os.Exit(0)
	}

	createTransactionCmd := flag.NewFlagSet("transact", flag.ExitOnError)
	from := createTransactionCmd.String("from", "", "source wallet address")
	to := createTransactionCmd.String("to", "", "destination wallet address")
	amount := createTransactionCmd.Int("amount", 0, "transaction value")

	getBalance := flag.NewFlagSet("getBalance", flag.ExitOnError)
	accountAddress := getBalance.String("address", "", "wallet address")

	getBlockTransactions := flag.NewFlagSet("getBlockTransactions", flag.ExitOnError)
	blockID := getBlockTransactions.String("blockID", "", "block ID")

	switch os.Args[1] {
	case "transact":
		err := createTransactionCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	case "getBalance":
		err := getBalance.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	case "getBlockTransactions":
		err := getBlockTransactions.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	}

	if createTransactionCmd.Parsed() {
		transactions.CreateTransaction(*from, *to, *amount)
	} else if getBalance.Parsed() {
		transactions.GetBalance(*accountAddress)
	} else if getBlockTransactions.Parsed() {
		transactions.GetBlockTransactions(*blockID)
	}

}

func help() {
	fmt.Println("simpleblockchain-cli helps you send transactions, get an account balance, or get a block's information")
	fmt.Println("To do so you can use these commands:")
	fmt.Println("transact -from $your_address -to $receipiant_address -amount $transaction value")
	fmt.Println("getBalance -address $address")
	fmt.Println("getBlockTransaction -blockID $blockID")
}

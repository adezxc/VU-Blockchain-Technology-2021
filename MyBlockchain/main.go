package main

import (
	"fmt"
	"strconv"
)

func main() {
	Users := generateUsers()
	Transactions = generateTransactions(Users)
	checkTransactions(Transactions, Users)
	getTransactionsForBlock(Transactions)

	bc := NewBlockchain()
	bc.AddBlock("Send 2 VUCoin to Petras")
	bc.AddBlock("Send 4 more VUCoin to Petras")
	bc.AddBlock("Test work")
	bc.AddBlock("one more")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Version: %s\n", block.Version)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
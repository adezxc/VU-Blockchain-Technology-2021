package main

import (
	"fmt"
	"strconv"

	"github.com/adezxc/VU-Blockchain-Technology-2021/blockchain/hashfunction"
)

func main() {
	fmt.Println("yes")
	hashfunction.Hash()
	chain := bc.InitBlockChain()
	fmt.Println("yes")
	chain.AddBlock("first block after genesis")

	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")
	
	for _, block := range chain.Blocks {

		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)

		pow := bc.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}

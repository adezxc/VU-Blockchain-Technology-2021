package main

import (
	"bytes"
	"fmt"

	"github.com/adezxc/VU-Blockchain-Technology-2021/blockchain/hash"
)

type Header struct {
	Timestamp        []byte
	Version          string
	Hash             []byte
	PrevHash         []byte
	Nonce            int
	DifficultyTarget int
}
type Block struct {
	Header
	Transactions []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) CalcHash() {
	info := bytes.Join([][]byte{[]byte(b.Transactions), b.Header.PrevHash}, []byte{})
	hash := hash.Hash(info)
	b.Hash = hash[:]
}

func CreateBlock(Transactions string, prevHash []byte) *Block {
	block := &Block{Header{}, []byte(Transactions)}
	block.CalcHash()
	return block
}

func (chain *Blockchain) AddBlock(Transactions string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(Transactions, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash: %x\n", block.Header.PrevHash)
		fmt.Printf("data: %s\n", block.Transactions)
		fmt.Printf("hash: %x\n", block.Header.Hash)
	}
}

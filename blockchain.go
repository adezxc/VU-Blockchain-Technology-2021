package main

// Blokų grandinės struktūra
type Blockchain struct {
	blocks []*Block
}

// Bloko pridėjimo funkcija
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// Naujos grandinės kūrimo funkcija
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
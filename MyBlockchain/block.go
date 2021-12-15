package main

import (
	//"strconv"
	"time"
	//"math/rand"
)

const Version = "v1.0"

var Transactions []Transaction

// Bloko struktūra
type Block struct {
	Timestamp     int64
	Version       string
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// Naujo bloko kūrimo funkcija
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), Version, []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Pirmojo grandinėje bloko kūrimo funkcija
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

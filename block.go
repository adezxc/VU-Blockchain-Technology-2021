package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)
const Version = "v1.0"
// Bloko struktūra
type Block struct {
	Timestamp     int64
	Version       string
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// Bloko hash'o apskaičiavimo funkcija
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	versionInt, _ := strconv.Atoi(b.Version)
	version := []byte(strconv.FormatInt(int64(versionInt), 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp, version}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// Naujo bloko kūrimo funkcija
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), Version, []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// Pirmojo grandinėje bloko kūrimo funkcija
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
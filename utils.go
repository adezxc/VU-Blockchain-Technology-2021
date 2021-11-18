package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Funkcija kuri konvertuoja int64 į []byte masyvą 
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
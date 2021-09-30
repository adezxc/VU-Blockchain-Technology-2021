package main

import (
	//"encoding/binary"
	"log"
	"strconv"
	"encoding/hex"
)

//import "encoding/hex"



func hash(text []byte) (result []byte) {
	prime := 1297
	var hash uint64 = uint64(nextPrime(int(prime)))
	for i := 0; i < len(text); i++ {
		placeholder := hash
		hash = (hash >> 7) ^ ((placeholder * (uint64(text[i])) >> 3))
	}

	charArray := "0123456789abcdef"
	initialResult := ""
	var hashString = strconv.FormatUint(hash, 16)

	for i := 0; i < 64; i++ {
		value := hashString[i%len(hashString)] + byte(i*nextPrime(nextPrime(prime)))
		initialResult += string(charArray[value%(byte)(len(charArray))])
	}
	result, err := hex.DecodeString(initialResult)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
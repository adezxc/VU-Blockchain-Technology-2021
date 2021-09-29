package main

import (
	//"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
)

//import "encoding/hex"

func main() {
	src, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash(src))
}

func hash(text []byte) (result string) {
	prime := uint64(1297)
	var hash uint64 = prime
	if len(text) == 0 {
		return
	}

	for i := 0; i < len(text); i++ {
		placeholder := hash
		hash = (hash >> 7) ^ ((placeholder * (uint64(text[i])) >> 2))
		
	}

	var charlist string = "01"
	var hashString = strconv.FormatUint(hash, 10)

	for i := 0; i < 256; i++ {
		index := hashString[i%len(hashString)] + (byte)(uint64(i)*prime)
		result += (string)(charlist[index%(byte)(len(charlist))])
	}

	return result
}

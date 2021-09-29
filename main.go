package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"encoding/binary"
)

//import "encoding/hex"

func main() {
	src, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash(src))
}

func hash(text []byte) string {
	prime := 1297
	var hash uint64 = 0
	if len(text) == 0 {
		return ""
	}
	var char = binary.BigEndian.Uint64(text)

	placeholder := char
	for i := 0; i < len(text); i++ {
		char = (char * uint64(prime)) % placeholder
		hash = (hash << 11) + (char % (uint64(i) + 1))
	}

	var charlist string = "0123456789abcdef"
	var hashString = strconv.FormatUint(hash, 10)

	result := ""
	for i := 0; i < 64; i++ {
		index := hashString[i%len(hashString)] + (byte)(i*prime)
		result += (string)(charlist[index%(byte)(len(charlist))])
	}

	return result
}

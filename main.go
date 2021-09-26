package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	//"math/bits"
	"encoding/binary"
)

//import "encoding/hex"

func main() {
	src, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash(src))
}

func hash(text []byte) string {
    prime :=  1297
	var hash uint64 = 0
	if len(text) == 0 {
		return ""
	}
	var char = binary.BigEndian.Uint64(text[0:])

	for i := 0; i < len(text); i++ {
		hash = ((hash << 5) & hash) + char
	}

	var charlist string = "0123456789abcdef"
    var hashString = strconv.FormatUint(hash, 10)
    fmt.Println(len(hashString))

    result := ""
    for i := 0; i < 64; i++ {
        index := hashString[i % len(hashString)] + (byte)(i * prime);
        result += (string)(charlist[index % (byte)(len(charlist))])
    }
    
	return result
}

package main

import (
	//"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"crypto/sha256"
	"crypto/md5"
)

//import "encoding/hex"

func main() {
	src, err := os.ReadFile("konstitucija.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash(src))
	start := time.Now()
	sum := md5.Sum(src)
	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("MD5 took %s", elapsed)
	start = time.Now()
	sum1 := sha256.Sum256(src)
	fmt.Println(sum1)
	elapsed = time.Since(start)
	log.Printf("Sha took %s", elapsed)
}

func hash(text []byte) (result string) {
	start := time.Now()
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
	elapsed := time.Since(start)
	log.Printf("Function took %s", elapsed)

	return result
}

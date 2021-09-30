package main

import (
	"bufio"
	//"crypto/md5"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	start := time.Now()
	for scanner.Scan() {
		hash(scanner.Bytes())
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	file.Close()
	
	
	file, err = os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}


	scanner1 := bufio.NewScanner(file)
	start = time.Now()
	for scanner1.Scan() {
		sha256.Sum256(scanner1.Bytes())
	}
	elapsed = time.Since(start)
	fmt.Println(elapsed)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	/*for i := 0; i < 50000; i++ {
		start := time.Now()
		fmt.Println(hash(src))
		elapsed := time.Since(start)
		log.Printf("My algo took %s", elapsed)
		start = time.Now()
		sum := md5.Sum(src)
		fmt.Println(sum)
		elapsed = time.Since(start)
		log.Printf("MD5 took %s", elapsed)
		start = time.Now()
		sum1 := sha256.Sum256(src)
		fmt.Println(sum1)
		elapsed = time.Since(start)
		log.Printf("Sha took %s", elapsed)
	}*/

}

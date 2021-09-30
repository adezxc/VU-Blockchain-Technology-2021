package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"time"
	"crypto/md5"
)

func konstitucija() {

	var elapsed time.Duration
	for i := 0; i < 5000; i++ {
		file, err := os.Open("test_files/konstitucija.txt")
		if err != nil {
			log.Fatal(err)
		}
		MyScanner := bufio.NewScanner(file)
		start := time.Now()
		for MyScanner.Scan() {
			hash(MyScanner.Bytes())
		}
		elapsed += time.Since(start)
		file.Close()
	}
	fmt.Printf("My hash function average is %d microseconds\n", (elapsed.Microseconds() / 5000))
	elapsed = time.Second * 0
	for i := 0; i < 5000; i++ {
		file, err := os.Open("test_files/konstitucija.txt")
		if err != nil {
			log.Fatal(err)
		}
		ShaScanner := bufio.NewScanner(file)
		start := time.Now()
		for ShaScanner.Scan() {
			sha256.Sum256(ShaScanner.Bytes())
		}
		elapsed += time.Since(start)
		file.Close()
	}
	fmt.Printf("SHA256 average is %d microseconds\n", (elapsed.Microseconds() / 5000))
	elapsed = time.Second * 0
	for i := 0; i < 5000; i++ {
		file, err := os.Open("test_files/konstitucija.txt")
		if err != nil {
			log.Fatal(err)
		}
		md5Scanner := bufio.NewScanner(file)
		start := time.Now()
		for md5Scanner.Scan() {
			md5.Sum(md5Scanner.Bytes())
		}
		elapsed += time.Since(start)
		file.Close()
	}
	fmt.Printf("MD5 average is %d microseconds\n", (elapsed.Microseconds() / 5000))

}

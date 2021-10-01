package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func stringpairs() {
	file10, err := os.Open("test_files/10chars")
	if err != nil {
		log.Fatal(err)
	}
	Scanner10 := bufio.NewScanner(file10)
	for i := 0; i < 25000; i++ {
		Scanner10.Scan()
		first := Scanner10.Bytes()
		Scanner10.Scan()
		second := Scanner10.Bytes()
		if bytes.Equal(hash(first), hash(second)) {
			fmt.Printf("BZZ!! BZZ!! COLLISION DETECTED!!! pair %d", i)
		}
	}
	file10.Close()
	file100, err := os.Open("test_files/100chars")
	if err != nil {
		log.Fatal(err)
	}
	Scanner100 := bufio.NewScanner(file100)
	for i := 0; i < 25000; i++ {
		Scanner100.Scan()
		first := Scanner100.Bytes()
		Scanner100.Scan()
		second := Scanner100.Bytes()
		if bytes.Equal(hash(first), hash(second)) {
			fmt.Printf("BZZ!! BZZ!! COLLISION DETECTED!!! pair %d\n", i)
		}
	}
	file100.Close()
	file500, err := os.Open("test_files/500chars")
	if err != nil {
		log.Fatal(err)
	}
	Scanner500 := bufio.NewScanner(file500)
	for i := 0; i < 25000; i++ {
		Scanner500.Scan()
		first := Scanner500.Bytes()
		Scanner500.Scan()
		second := Scanner500.Bytes()
		if bytes.Equal(hash(first), hash(second)) {
			fmt.Printf("BZZ!! BZZ!! COLLISION DETECTED!!! pair %d\n", i)
		}
	}
	file500.Close()
	file1000, err := os.Open("test_files/1000chars")
	if err != nil {
		log.Fatal(err)
	}
	Scanner1000 := bufio.NewScanner(file1000)
	for i := 0; i < 25000; i++ {
		Scanner1000.Scan()
		first := Scanner1000.Bytes()
		Scanner1000.Scan()
		second := Scanner1000.Bytes()
		if bytes.Equal(hash(first), hash(second)) {
			fmt.Printf("BZZ!! BZZ!! COLLISION DETECTED!!! pair %d\n", i)
		}
	}
	file1000.Close()
}

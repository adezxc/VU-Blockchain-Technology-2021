package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thanhpk/randstr"
)

func generate_data() {
	sameLetter := true
	randomString := ""
	file10, err := os.Create("test_files/10chars")
	if err != nil {
		log.Fatal(err)
	}
	file100, err := os.Create("test_files/100chars")
	if err != nil {
		log.Fatal(err)
	}
	file500, err := os.Create("test_files/500chars")
	if err != nil {
		log.Fatal(err)
	}
	file1000, err := os.Create("test_files/1000chars")
	if err != nil {
		log.Fatal(err)
	}
	file100000, err := os.Create("test_files/100000strings")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 25000; i++ {
		firstString := randstr.String(10)
		secondString := randstr.String(10)
		file10.WriteString(fmt.Sprintf("%s\n%s\n", firstString, secondString))
		firstString = randstr.String(100)
		secondString = randstr.String(100)
		file100.WriteString(fmt.Sprintf("%s\n%s\n", firstString, secondString))
		firstString = randstr.String(500)
		secondString = randstr.String(500)
		file500.WriteString(fmt.Sprintf("%s\n%s\n", firstString, secondString))
		firstString = randstr.String(1000)
		secondString = randstr.String(1000)
		file1000.WriteString(fmt.Sprintf("%s\n%s\n", firstString, secondString))
	}
	for i := 0; i < 100000; i++ {
		sameLetter = true
		firstString := randstr.String(957)
		file100000.WriteString(fmt.Sprintf("%s\n", firstString))
		for sameLetter == true {
			randomString = randstr.String(1)
			if (string(firstString[333]) != randomString) {
				sameLetter = false
			} else {
				continue
			}
		}
		firstString = firstString[:333] + randomString + firstString[334:]
		file100000.WriteString(fmt.Sprintf("%s\n", firstString))
	}
	defer file10.Close()
	defer file100.Close()
	defer file500.Close()
	defer file1000.Close()
	defer file100000.Close()
}

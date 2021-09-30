package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/agnivade/levenshtein"
)

func comparison() {
	file, err := os.Open("test_files/100000strings")
	if err != nil {
		log.Fatal(err)
	}
	binarySum := 0
	var binaryMin, binaryMax int = 100, 0
	var hexMin, hexMax int = 100, 0
	hexSum := 0
	scanner := bufio.NewScanner(file)
	for i := 0; i < 100000; i++ {
		scanner.Scan()
		first := scanner.Bytes()
		firstHash := hash(first)
		firstBinary := fmt.Sprintf("%b", firstHash)
		firstHex := fmt.Sprintf("%x", firstHash)
		scanner.Scan()
		second := scanner.Bytes()
		secondHash := hash(second)
		secondBinary := fmt.Sprintf("%b", secondHash)
		secondHex := fmt.Sprintf("%x", secondHash)
		binaryLeven := levenshtein.ComputeDistance(firstBinary, secondBinary)
		if binaryLeven > binaryMax {
			binaryMax = binaryLeven
		}
		if binaryLeven < binaryMin{
			binaryMin = binaryLeven
		}
		hexLeven := levenshtein.ComputeDistance(firstHex, secondHex)
		if (binaryLeven == 0 || hexLeven == 0) {
			fmt.Println(i)
		}
		if hexLeven > hexMax {
			hexMax = hexLeven
		}
		if hexLeven < hexMin {
			hexMin = hexLeven
		}
		binarySum += binaryLeven
		hexSum += hexLeven
	}
	binaryAvg := float32(binarySum) / 1000 / 256
	hexAvg := float32(hexSum) / 1000 / 64
	fmt.Printf("Binary: Average: %2.2f Minimum: %d Maximum: %d\n", binaryAvg, binaryMin, binaryMax)
	fmt.Printf("Hex: Average: %2.2f Minimum: %d Maximum: %d\n", hexAvg, hexMin, hexMax)
}

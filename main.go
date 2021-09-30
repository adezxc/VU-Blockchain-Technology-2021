package main

import (
	//"bufio"
	//"crypto/sha256"
	//"fmt"
	//"log"
	//"os"
	//"time"
)

func main() {

	/*file, err := os.Open("test_files/one_letter_a.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner1 := bufio.NewScanner(file)
	start := time.Now()
	for scanner1.Scan() {
		fmt.Printf("%x\n", sha256.Sum256(scanner1.Bytes()))
		fmt.Printf("%b\n", sha256.Sum256(scanner1.Bytes()))
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	file, err = os.Open("test_files/one_letter_a.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	start = time.Now()
	for scanner.Scan() {
		fmt.Printf("%x\n", hash(scanner.Bytes()))
		fmt.Printf("%b\n", hash(scanner.Bytes()))
	}
	elapsed = time.Since(start)

	fmt.Println(elapsed)
	file.Close()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}*/
	konstitucija()
}

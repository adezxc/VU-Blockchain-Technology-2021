package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fileFlag := flag.String("f", "", "provide file for hashing")
	generateFlag := flag.Bool("gen", false, "does this need generating data")
	konstitucijaFlag := flag.Bool("kon", false, "do you want algorithm speed comparison")
	collisionFlag := flag.Bool("comp", false, "do you want to check for collisions")
	differenceFlag := flag.Bool("diff", false, "do you want to find binary/hex difference")

	flag.Parse()

	if *fileFlag != "" {
		file, err := os.ReadFile(*fileFlag)
		if err != nil {
			fmt.Println("File couldn't be opened")
		}
		fmt.Printf("%x\n", hash((file)))
		
	}
	if *generateFlag {
		generate_data()
	}
	if *konstitucijaFlag {
		konstitucija()
	}
	if *collisionFlag {
		stringpairs()
	}
	if *differenceFlag {
		comparison()
	}
}

package main

import "fmt"
import "encoding/hex"

func main() {
	src := []byte("Hello World!")

	dst := make([]byte, 1024)
    test := hex.Encode(dst, src)

    fmt.Printf("%s %d\n", dst, test)
}

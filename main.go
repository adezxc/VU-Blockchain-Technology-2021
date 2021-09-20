package main

import 
(
    "fmt"
    "io/ioutil"
    "log"
)
//import "encoding/hex"

func main() {
    src, err := ioutil.ReadFile("test.txt")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%.8b", src)

}

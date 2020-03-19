package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter an integer in base10 to be converted to binary: ")
	var number int64
	fmt.Scan(&number)
	fmt.Printf("%b", number)
}

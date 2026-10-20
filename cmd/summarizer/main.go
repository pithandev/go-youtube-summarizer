package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	for i, arg := range os.Args[1:] {
		fmt.Printf("Arg %d: %s\n", i+1, arg)
	}
}

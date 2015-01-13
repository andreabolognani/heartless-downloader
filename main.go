package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "URL")
		os.Exit(1)
	}

	fmt.Println("URL:", os.Args[1])
}

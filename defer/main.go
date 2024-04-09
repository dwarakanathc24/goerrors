package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("error", err)
	}
	// defer f.Close()
	log.SetOutput(f)
	log.Println("deferrred the file")
}

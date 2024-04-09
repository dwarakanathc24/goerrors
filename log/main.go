package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(" the main function in log")
	f, err := os.Open("names.txt")
	if err != nil {
		log.Println(" error happend:", err)
		fmt.Println(" the error happend", f, err)
	}

	f, e := os.Create("log.txt")
	if e != nil {
		fmt.Println(" error happend", e)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, e2 := os.Open("no-file.txt")
	if e2 != nil {
		log.Println("error happend", e2)
	}

	defer f2.Close()
	// panic(err)
	log.Fatalln("this is fatal function writing log to the file", e2)
	fmt.Println(" open log.txt file for more information")
}

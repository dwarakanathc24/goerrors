package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Copy("names.txt", "new_names.txt")

}

func Copy(src, dest string) (writtern int64, err error) {
	s, err := os.Open(src)
	if err != nil {
		fmt.Println(" error happend", err)
	}

	d, err := os.Create(dest)

	if err != nil {
		fmt.Println(" error occurred while creating a file", err)

	}

	w, err := io.Copy(d, s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(w)
	return
}

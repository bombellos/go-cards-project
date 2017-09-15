package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	printFileContent(readFile(os.Args[1]))
}

func readFile(fileName string) *os.File {
	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Something went wrong: ", err)
		os.Exit(1)
	}

	return file
}

func printFileContent(file *os.File) {
	io.Copy(os.Stdout, file)
}

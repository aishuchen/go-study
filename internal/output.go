package internal

import (
	"fmt"
	"io"
	"os"
)

func readFile() string {
	f, err := os.Open("example.txt")
	if err != nil {
		panic("failed to open file: " + err.Error())
	}
	data, err := io.ReadAll(f)
	if err != nil {
		panic("failed to read file: " + err.Error())
	}
	return string(data)
}

func Output() {
	curDir, _ := os.Getwd()
	fmt.Printf("Current dir: %s\n", curDir)
	s := readFile()
	fmt.Println(s)
}

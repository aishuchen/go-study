package main

import (
	"example.com/go-study/internal"
	"fmt"
	"os"
)

func main() {
	curDir, _ := os.Getwd()
	fmt.Printf("Current dir: %s\n", curDir)
	internal.Output()
}

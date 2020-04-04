package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var envNames [10]string
	count := 0
	for _, itr := range os.Environ() {
		if count < 10 {
			envNames[count] = itr
		}
		count++
	}

	fmt.Println(envNames[0])

	runeVar := []rune(envNames[0])

	fmt.Println(&runeVar)

	homeDir := os.Getenv("HOME")
	userName := strings.Replace(homeDir, "/home/", "", -1)

	memoryAdress := &homeDir
	fmt.Println(memoryAdress)
	fmt.Println(*memoryAdress, userName)
}

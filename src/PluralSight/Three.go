package main

import "fmt"

func main() {

	count, result := reverseString("Atul")

	fmt.Print(count, result)
}

func reverseString(inputString string) (count int16, result string) {
	count = 0
	result = inputString

	return count, result
}

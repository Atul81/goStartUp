package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the input string for expansion => ")
	inpString, _ := reader.ReadString('\n')
	expandString(strings.TrimSuffix(inpString, "\n"))
}

func expandString(inpString string) (count int) {
	fmt.Println("String to be converted for", inpString)
	now := time.Now()
	// var foundOpenBracket bool
	var preSubString string
	var sufSubString string
	var operationString [6]string
	var initCount int
	//	var operationCount int
	//	var foundCloseBracket bool
	reverseIndexing := len(inpString) - 1
	for itr := 0; itr < len(inpString); itr++ {
		if inpString[itr] == '{' {
			initCount = len(inpString) - (len(inpString) - itr)
			preSubString = inpString[:initCount]
			//			foundOpenBracket = true
		}

		if inpString[reverseIndexing-itr] == '}' {
			fmt.Println("Found closing bracket at", strings.IndexAny(inpString, "}"))
		}

		//if !foundCloseBracket && foundOpenBracket && inpString[itr] != '}' && inpString[itr] != ','{
		//	operationString[operationCount] = preSubString+ string(inpString[itr])
		//	initCount++
		//	operationCount++
		//	fmt.Print(operationString)
		//}
		//
		//if inpString[itr] == '}' && itr+1 < len(inpString){
		//	sufSubString = inpString[itr+1:]
		//	foundCloseBracket = true
		//}

		fmt.Println(preSubString, sufSubString)
	}
	performConcatenation(operationString, preSubString, sufSubString)
	return int(time.Now().Sub(now))
}

func performConcatenation(operationalString [6]string, prefix string, suffix string) {

}

package main

import (
	"fmt"
	"reflect"
	"runtime"
)

//using a declared variable with in function needs to be used due to garbage collection

var (
	checkOne   string
	checkTwo   = "CheckTwo"
	checkThree = 90.23
)

func main() {
	checkFour := 5 //Another type of declaration only within method
	fmt.Println(checkFour)
	var returnParam = check("Passing arguments")
	var l, er = fmt.Println(returnParam)
	fmt.Println(l, er)
	fmt.Print("Hello World")
	fmt.Println(reflect.TypeOf(checkThree))
}

func check(check string) string {
	fmt.Println(check, "   ", runtime.GOOS)
	return "And It Begins"
}

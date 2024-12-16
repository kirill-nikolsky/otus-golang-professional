package main

import (
	"fmt"

	rev "golang.org/x/example/hello/reverse"
)

const initString string = "Hello, OTUS!"

func reverseString(str string) (revStr string) {
	return rev.String(str)
}

func printReversedInitString() {
	revStr := reverseString(initString)

	fmt.Println(revStr)
}

func main() {
	printReversedInitString()
}

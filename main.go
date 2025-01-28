package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	converted_int := strconv.Itoa(x)

	new_string := converted_int[len(converted_int)-1:len(converted_int)] + converted_int[1:len(converted_int)-1] + converted_int[0:1]
	new_int, _ := strconv.Atoi(new_string)

	if new_int == x {
		return true
	} else {
		return false
	}
}

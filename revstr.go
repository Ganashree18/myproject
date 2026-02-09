package main

import "fmt"

func main() {
	var s string
	// var rev string
	fmt.Println("Let's reverse a string")
	fmt.Println("Helloo")
	fmt.Print("Enter a string: ")
	fmt.Scan(&s)

	// rev:=""
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= 0; j-- {
			if string(s[i]) == string(s[j]) {
				fmt.Println("Palindrome")
			} else {
				fmt.Println("Not Palindrome")
				break
			}

		}
	}
}

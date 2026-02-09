package main

import "fmt"

func main(){

	fmt.Println("haiii")
	fmt.Println("helloooo")

}

func main() {
	var n int
	fmt.Print("Enter number of terms: ")
	fmt.Scan(&n)

	a, b := 0, 1
	fmt.Print("Fibonacci series: ")
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}

func hai() {
	fmt.Println("HelloooHaiiii")
}
func baii() {
	fmt.Println("Byeeeeebaiiii")
}

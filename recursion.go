package main

import "fmt"

func fact(n int) int {
	fmt.Println("n = ", n)
	if n == 0 {
		return 1
	}
	num := n * fact(n-1)
	fmt.Println("num = ", num)
	return num
}

func main() {
	fmt.Println(fact(7))
}

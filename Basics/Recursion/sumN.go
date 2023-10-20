package main

import "fmt"

func sumN(n int) int {
	if n == 1 {
		return 1
	}
	return n + sumN(n-1)
}

func main() {
	fmt.Println(sumN(5))
}

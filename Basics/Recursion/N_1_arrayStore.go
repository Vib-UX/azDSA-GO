package main

import "fmt"

func recurN_1(n int) []int {
	if n == 1 {
		var res []int
		res = append(res, 1)
		return res
	}
	var res []int
	res = append(res, n)
	res = append(res, recurN_1(n-1)...)

	return res

}

func main() {
	fmt.Println(recurN_1(10))
}

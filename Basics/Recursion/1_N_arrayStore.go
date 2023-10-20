package main

import "fmt"

func recurN(n int) []int {
	if n == 1 {
		var res []int
		res = append(res, 1)
		return res
	}
	res := recurN(n - 1)
	res = append(res, n)

	return res

}

func main() {
	fmt.Println(recurN(10))
}

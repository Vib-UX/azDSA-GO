package main

import "fmt"

func fibonnaciN(n int) []int {
	if n == 0 {
		return []int{0}
	}

	if n == 1 {
		return []int{0, 1}
	}

	var res []int
	res = fibonnaciN(n - 1)
	prev := fibonnaciN(n - 2)
	res = append(res, res[len(res)-1]+prev[len(prev)-1])

	return res
}

func main() {
	fmt.Println(fibonnaciN(10))
}

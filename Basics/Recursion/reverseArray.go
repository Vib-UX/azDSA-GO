package main

import "fmt"

// Space O(N)  Time O(N)
func reverseSlc(slc []int, n int) []int {
	if n == 1 {
		var res []int
		res = append(res, slc[n-1])
		return res
	}
	var rev []int
	rev = append(rev, slc[n-1])
	rev = append(rev, reverseSlc(slc, n-1)...)

	return rev

}

func main() {
	fmt.Println(reverseSlc([]int{1, 2, 3}, 3))
}

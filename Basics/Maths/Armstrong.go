package main

import (
	"fmt"
	"math"
)

func countDigi(n int) int {
	var temp int
	for n != 0 {
		temp++
		n /= 10
	}

	return temp
}

func validArmstrong(n int) bool {
	digits := countDigi(n)
	var sum int
	temp := n
	for temp != 0 {
		sum += int(math.Pow(float64(temp%10), float64(digits)))
		temp /= 10
	}

	return n == sum

}
func main() {
	fmt.Println(validArmstrong(371))
}

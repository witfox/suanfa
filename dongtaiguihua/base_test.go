package dongtaiguihua

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {

	res := fibonacci(4)

	fmt.Println(res)
}

func TestNumWays(t *testing.T) {
	res := numWays(5)
	fmt.Println(res)
}
func TestIsMatch(t *testing.T) {
	res := isMatch("aaa", "ab*.*")
	fmt.Println(res)
}

func TestMaxSubArray(t *testing.T) {
	res := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(res)
}

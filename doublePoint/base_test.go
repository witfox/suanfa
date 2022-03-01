package doublepoint

import (
	"fmt"
	"testing"
)

func TestExchange(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	res := exchange(nums)
	fmt.Println(res)
}

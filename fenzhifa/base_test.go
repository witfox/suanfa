package fenzhifa

import (
	"fmt"
	"testing"
)

func TestReversePairs(t *testing.T) {

	nums := []int{7, 3, 2, 6, 0, 1, 5, 4}
	res := reversePairs(nums)

	fmt.Println(res)
}

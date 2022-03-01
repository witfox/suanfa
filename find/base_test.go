package find

import (
	"fmt"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	str := "leetcode"
	res := firstUniqChar(str)
	fmt.Println(res)
}

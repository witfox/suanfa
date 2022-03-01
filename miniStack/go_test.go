package stack

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	rs := obj.Min()
	fmt.Println(rs)
	obj.Pop()
	rs2 := obj.Top()
	fmt.Println(rs2)
	rs3 := obj.Min()
	fmt.Println(rs3)
}

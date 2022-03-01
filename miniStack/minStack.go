package stack

import "container/list"

//------------包含min函数的栈（要求min()的时间复杂度为O(1)） 2022-01-27------------
type MinStack struct {
	A *list.List
	B *list.List //辅助栈
}

func Constructor() MinStack {
	return MinStack{
		A: list.New(),
		B: list.New(),
	}
}

func (min *MinStack) Push(x int) {
	min.A.PushBack(x)
	if min.B.Len() <= 0 || min.B.Back().Value.(int) >= x {
		min.B.PushBack(x)
	}
}

func (min *MinStack) Pop() {
	y := min.A.Remove(min.A.Back())
	if min.B.Back().Value == y {
		min.B.Remove(min.B.Back())
	}
}

func (min *MinStack) Top() int {
	return min.A.Back().Value.(int)
}

func (min *MinStack) Min() int {
	return min.B.Back().Value.(int)
}

package gosuanfa

import (
	"fmt"
	"testing"
)

func TestReplaceSpace(t *testing.T) {

	s := replaceSpace("我是， hello world")
	fmt.Println(s)
}

func TestReversePrint(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 2}

	node1.Next = node2
	node2.Next = node3
	node3.Next = nil
	s := reversePrint(node1)
	fmt.Println(s)
}

func TestCQueue(t *testing.T) {

	obj := Constructor()
	params1 := obj.DeleteHead()
	obj.AppendTail(5)
	obj.AppendTail(2)
	params3 := obj.DeleteHead()
	params4 := obj.DeleteHead()
	fmt.Println(params1)
	fmt.Println(params3)
	fmt.Println(params4)
}

func TestReverseList(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 2}

	node1.Next = node2
	node2.Next = node3
	node3.Next = nil
	s := reverseList2(node1)
	for s != nil {
		fmt.Println(s.Val)
		s = s.Next
	}
}

func TestLeftString(t *testing.T) {
	s := reverseLeftWords("abcdedf", 2)
	fmt.Println(s)
}

func TestMaxSlidingWindow(t *testing.T) {

	s := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)

	fmt.Println(s)
}

func TestStrToInt(t *testing.T) {

	res := strToInt("   -42")

	fmt.Println(res)
}

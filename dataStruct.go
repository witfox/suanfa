/**
* 数据结构
 */
package gosuanfa

import (
	"container/list"
	"fmt"
	"unicode"
)

//替换空格
func replaceSpace(s string) string {

	newStr := ""
	for _, char := range s {
		if unicode.IsSpace(char) {
			newStr += "%20"
		} else {
			newStr += string(char)
		}
	}
	return newStr
}

//单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//---------------从尾到头打印链表 2022-01-27----------------
func reversePrint(head *ListNode) []int {

	//获取链表的长度
	tmp := head
	len := 0
	for tmp != nil {
		len++
		tmp = tmp.Next
	}

	res := make([]int, len)
	tmp = head
	index := len - 1
	for tmp != nil {
		res[index] = tmp.Val
		index--
		tmp = tmp.Next
	}
	return res

}

//-------------用两个栈实现队列 2022-01-27-------------------
type CQueue struct {
	A *list.List //负责队列入数据
	B *list.List //负责队列出数据
}

func Constructor() CQueue {
	return CQueue{
		A: list.New(),
		B: list.New(),
	}
}

func (c *CQueue) AppendTail(value int) {
	c.A.PushBack(value)
}

func (c *CQueue) DeleteHead() int {
	//如果B栈有数据则返回栈尾数据
	if c.B.Len() > 0 {
		return c.B.Remove(c.B.Back()).(int)
	}
	if c.A.Len() == 0 {
		return -1
	}
	//将A的数据全部压入到B栈中
	for c.A.Len() > 0 {
		c.B.PushBack(c.A.Remove(c.A.Back())) //清除A栈的数据
	}
	//返回B栈尾数据
	return c.B.Remove(c.B.Back()).(int)
}

//----------------- 表示数值的字符串---------------------

//----------------- 反转链表 2022-01-27---------------------
//迭代(双指针法)
func reverseList(head *ListNode) *ListNode {

	var pre *ListNode = nil
	var cur *ListNode = head
	for cur != nil {
		tmp := cur.Next //暂存后继节点
		cur.Next = pre  //修改next应用指向
		pre = cur       //pre暂存cur
		cur = tmp       //cur访问下一个节点
	}
	return pre
}

//递归
func reverseList2(head *ListNode) *ListNode {
	return recur(head, nil)
}
func recur(cur *ListNode, pre *ListNode) *ListNode {

	fmt.Println(cur, pre)
	//中止条件
	if cur == nil {
		return pre
	}
	//递归后继节点
	res := recur(cur.Next, cur)
	//回溯处理时修改节点引用指向
	cur.Next = pre
	return res
}

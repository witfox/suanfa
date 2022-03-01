package copyrandomlist

import "fmt"

/**
* 剑指 Offer 35. 复杂链表的复制
* 请实现 copyRandomList 函数，复制一个复杂链表。
* 在复杂链表中,每个节点除了有一个 next 指针指向下一个节点，
* 还有一个 random 指针指向链表中的任意节点或者 null。
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//hash O(N)+O(N)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cur := head
	dic := make(map[*Node]*Node)
	for cur != nil {
		dic[cur] = &Node{cur.Val, nil, nil}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		dic[cur].Next = dic[cur.Next]
		dic[cur].Random = dic[cur.Random]
		cur = cur.Next
	}
	return dic[head]
}

//拼接 + 拆分 时间O(N) 空间O(1)
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	//拼接两个链表
	cur := head
	for cur != nil {
		temp := &Node{cur.Val, nil, nil}
		temp.Next = cur.Next
		cur.Next = temp
		cur = temp.Next
	}
	//给新链表的节点设置random节点
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next //cur.Random.Next就是复制后的新节点
		}
		cur = cur.Next.Next
	}
	//拆分两个链表
	pre := head
	cur = head.Next
	res := head.Next
	for cur.Next != nil {
		pre.Next = pre.Next.Next //修改next值
		cur.Next = cur.Next.Next
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = nil //单独处理原链表的尾节点
	return res     //返回新链表的头结点

}

//个人理解但是leetcode执行失败(浅拷贝,应该深拷贝)
//提示：Node with label 13 was not copied but a reference to the original one.
func copyRandomList2(head *Node) *Node {

	if head == nil {
		return nil
	}
	var cur = head
	var newNode = &Node{}
	pre := newNode
	for cur != nil {
		pre.Val = cur.Val
		pre.Next = cur.Next
		pre.Random = cur.Random

		fmt.Println(pre)
		cur = cur.Next
		if cur != nil {
			pre = pre.Next
		}

	}
	return newNode
}

package doublepoint

import "strings"

/*
双指针
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

//删除链表的节点
//给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点
//【前后双指针】
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	pre := head
	cur := head.Next
	for cur != nil && cur.Val != val {
		//移动指针
		pre = cur
		cur = cur.Next
	}
	if cur != nil {
		pre.Next = cur.Next
	}
	return head
}

//调整数组顺序使奇数位于偶数前面
//输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的 “前半部分”，所有偶数在数组的 “后半部分”
//【首尾双指针】
func exchange(nums []int) []int {
	//标识前半部分，和后半部分索引
	i, j := 0, len(nums)-1
	for i < j {
		//循环左半部分
		for i < j && nums[i]%2 != 0 {
			i++
		}
		//循环右部分
		for i < j && nums[j]%2 == 0 {
			j--
		}
		//交换奇偶数位置
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
	}

	return nums
}

//链表中倒数第 k 个节点
//【前后指针相差K】
func getKthFromEnd(head *ListNode, k int) *ListNode {
	former, latter := head, head
	//former指针先向前移动K个节点
	for i := 0; i < k; i++ {
		former = former.Next
	}
	//开始循环
	for former != nil {
		latter = latter.Next
		former = former.Next
	}
	return latter
}

//合并两个排序的链表
//输入两个递增排序的链表，合并这两个链表并使【新链表】中的节点仍然是递增排序的。
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//初始化一个新链表
	dum := &ListNode{0, nil}
	cur := dum

	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dum.Next
}

// 两个链表的第一个公共节点
// a+(b-c) = b+(a-c)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	A := headA
	B := headB

	for A != B {
		if A == nil {
			A = headB
		} else {
			A = A.Next
		}
		if B == nil {
			B = headA
		} else {
			B = B.Next
		}
	}
	return A
}

//和为 s 的两个数字
//双指针 i , j 分别指向数组 numsnums 的左右两端 （俗称对撞双指针）
func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	res := make([]int, 0)
	for i < j {
		s := nums[i] + nums[j]
		if s > target {
			j--
		} else if s < target {
			i++
		} else {
			res = append(res, nums[i])
			res = append(res, nums[j])
			return res
		}
	}
	return res
}

//翻转单词顺序
//倒序遍历字符串 s，记录单词左右索引边界 i , j；
func reverseWords(s string) string {
	res := make([]string, 0)
	//去掉首尾空格
	s = strings.TrimSpace(s)
	i, j := len(s)-1, len(s)-1
	for i >= 0 {
		for i >= 0 && s[i] != ' ' {
			i--
		}
		res = append(res, s[i+1:j+1]+" ")
		for i >= 0 && s[i] == ' ' {
			i--
		}
		//继续下一个单词
		j = i
	}
	return strings.TrimSpace(strings.Join(res, ""))
}

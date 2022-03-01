package search

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	borad := [][]byte{{'A', 'B', 'C', 'D'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	res := exist(borad, word)

	fmt.Println(res)
}

func TestMovingCount(t *testing.T) {
	res := movingCount2(1, 2, 1)
	fmt.Println(res)
}

func TestDfsSearch(t *testing.T) {
	tree1 := &TreeNode{1, nil, nil}
	tree2 := &TreeNode{2, nil, nil}
	tree3 := &TreeNode{3, nil, nil}
	tree4 := &TreeNode{4, nil, nil}
	tree5 := &TreeNode{5, nil, nil}

	tree4.Left = tree2
	tree4.Right = tree5
	tree2.Left = tree1
	tree2.Right = tree3
	node := treeToDoublyList(tree4)
	fmt.Println(node.Left.Val, node.Right.Val)
}

func TestSerialzie(t *testing.T) {
	tree1 := &TreeNode{1, nil, nil}
	tree2 := &TreeNode{2, nil, nil}
	tree3 := &TreeNode{3, nil, nil}
	tree4 := &TreeNode{4, nil, nil}
	tree5 := &TreeNode{5, nil, nil}

	tree1.Left = tree2
	tree1.Right = tree3
	tree3.Left = tree4
	tree3.Right = tree5
	res := serialzie(tree1)
	fmt.Println(res)
}
func TestDeserialzie(t *testing.T) {
	res := deserialzie("[1,2,3,null,null,4,5,null,null,null,null]")
	fmt.Println(res.Left.Val, res.Right.Val)
}

func TestPermutation(t *testing.T) {

	res := permutation("aab")
	fmt.Println(res)
	res = permutation("abc")
	fmt.Println(res)
}

func TestKthLargest(t *testing.T) {
	tree1 := &TreeNode{1, nil, nil}
	tree2 := &TreeNode{2, nil, nil}
	tree3 := &TreeNode{3, nil, nil}
	tree4 := &TreeNode{4, nil, nil}
	tree5 := &TreeNode{5, nil, nil}

	tree1.Left = tree2
	tree1.Right = tree3
	tree3.Left = tree4
	tree3.Right = tree5
	res := kthLargest(tree1, 2)
	fmt.Println(res)
}

func TestMaxDepth(t *testing.T) {
	tree1 := &TreeNode{1, nil, nil}
	tree2 := &TreeNode{2, nil, nil}
	tree3 := &TreeNode{3, nil, nil}
	tree4 := &TreeNode{4, nil, nil}
	tree5 := &TreeNode{5, nil, nil}

	tree1.Left = tree2
	tree1.Right = tree3
	tree3.Left = tree4
	tree3.Right = tree5
	res := maxDepth2(tree1)
	fmt.Println(res)
}

func TestIsBalanced(t *testing.T) {
	tree1 := &TreeNode{1, nil, nil}
	tree2 := &TreeNode{2, nil, nil}
	tree3 := &TreeNode{3, nil, nil}
	tree4 := &TreeNode{4, nil, nil}
	tree5 := &TreeNode{5, nil, nil}
	tree6 := &TreeNode{6, nil, nil}
	tree7 := &TreeNode{7, nil, nil}

	tree1.Left = tree2
	tree1.Right = tree3
	tree3.Left = tree4
	tree3.Right = tree5
	tree4.Left = tree6
	tree4.Right = tree7
	res := isBalanced(tree1)
	fmt.Println(res)
}
func TestSumNums(t *testing.T) {
	res := sumNums(3)
	fmt.Println(res)
	res = sumNums(2)
	fmt.Println(res)
}

func TestLowestCommonAncestor(t *testing.T) {
	tree0 := &TreeNode{0, nil, nil}
	tree2 := &TreeNode{2, nil, nil}
	tree3 := &TreeNode{3, nil, nil}
	tree4 := &TreeNode{4, nil, nil}
	tree5 := &TreeNode{5, nil, nil}
	tree6 := &TreeNode{6, nil, nil}
	tree7 := &TreeNode{7, nil, nil}
	tree8 := &TreeNode{8, nil, nil}
	tree9 := &TreeNode{9, nil, nil}

	tree6.Left = tree2
	tree6.Right = tree8
	tree2.Left = tree0
	tree2.Right = tree4
	tree4.Left = tree3
	tree4.Right = tree5
	tree8.Left = tree7
	tree8.Right = tree9
	res := lowestCommonAncestor2(tree6, tree3, tree5)
	fmt.Println(res.Val)
}

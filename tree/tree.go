package tree

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

//二叉树中的最大路径和
//递归
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var recur func(*TreeNode) int

	recur = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		//递归计算左右子树节点的最大贡献值
		leftMax := Max(recur(node.Left), 0)
		rightMax := Max(recur(node.Right), 0)

		//回溯
		sum := node.Val + leftMax + rightMax

		maxSum = Max(maxSum, sum)

		//返回节点的最大贡献值
		return node.Val + Max(leftMax, rightMax)
	}
	recur(root)
	return maxSum
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

package fenzhifa

/*
分治法，即分而治之
就是把一个复杂的问题分成两个或更多的相同或相似的子问题，再把子问题分成更小的子问题……
直到最后子问题可以简单的直接求解，原问题的解即子问题的解的合并
【递归】
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//重建二叉树
//输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点
//根据「分治算法」思想，对于树的左、右子树，仍可复用以上方法划分子树的左右子树。
func buildTree(preorder []int, inorder []int) *TreeNode {

	dic := make(map[int]int, 0) //存储中序中的数值
	for i := 0; i < len(inorder); i++ {
		dic[inorder[i]] = i
	}
	var recur func(root, left, right int) *TreeNode

	recur = func(root, left, right int) *TreeNode {
		//终止递归
		if left > right {
			return nil
		}
		//建立根节点
		node := &TreeNode{preorder[root], nil, nil}
		//划分根节点，左子树，右子树
		i := dic[preorder[root]]
		//开启左子树递归
		node.Left = recur(root+1, left, i-1)
		//开启右子树递归(根节点索引+左子树长度(i-left)+1)
		node.Right = recur(root+i-left+1, i+1, right)
		return node
	}

	return recur(0, 0, len(preorder)-1)
}

//数值的整数次方
//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。

//二叉搜索树的后序遍历序列
//输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。
//如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同
func verifyPostorder(postorder []int) bool {

	var recur func(store []int, i, j int) bool

	recur = func(store []int, i, j int) bool {
		//递归终止条件
		if i >= j {
			return true
		}
		//找出大于postorder[j]的数
		p := i
		for store[p] < store[j] {
			p++
		}
		//划分左右子树（分治法）
		m := p
		for store[p] > store[j] {
			p++
		}
		return p == j && recur(store, i, m-1) && recur(store, m, j-1)
	}
	return recur(postorder, 0, len(postorder)-1)
}

//数组中的逆序对
//在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
//二分法排序
func reversePairs(nums []int) int {
	numsLen := len(nums)
	tmp := make([]int, numsLen)

	var mergeSort func(left, right int) int
	mergeSort = func(left, right int) int {
		// 终止条件
		if left >= right {
			return 0
		}
		//找出当前数组的中位，拆分左右子数组
		m := (left + right) / 2
		//递归左右子数组
		res := mergeSort(left, m) + mergeSort(m+1, right)
		// 合并阶段
		i, j := left, m+1 //标记左右子数组的首位字符索引
		for k := left; k <= right; k++ {
			tmp[k] = nums[k]
		}
		for k := left; k <= right; k++ {
			if i == m+1 {
				nums[k] = tmp[j]
				j++
			} else if j == right+1 || tmp[i] <= tmp[j] {
				nums[k] = tmp[i]
				i++
			} else {
				//如果左边大于右边，则存在逆序
				nums[k] = tmp[j]
				j++
				res += m - i + 1
			}
		}
		return res
	}
	return mergeSort(0, numsLen-1)
}

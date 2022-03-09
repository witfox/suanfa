//深度优先搜索DFS;先序遍历、中序遍历、后序遍历；
//广度优先搜索BFS;层序遍历（即按层遍历）；
package search

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

//剑指 Offer 12. --------------矩阵中的路径-------------
//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。
//如果 word 存在于网格中，返回 true ；否则，返回 false 。
func exist(board [][]byte, word string) bool {
	char := []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, char, i, j, 0) {
				return true
			}
		}
	}
	return false
}

//深度优先搜索 DFS
func dfs(board [][]byte, char []byte, i int, j int, k int) bool {
	//越界返回false
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != char[k] {
		return false
	}
	//word匹配结束，返回true
	if k == len(char)-1 {
		return true
	}
	board[i][j] = '0' //设置为空防止重复搜索
	//从左右上下递归
	res := dfs(board, char, i+1, j, k+1) || dfs(board, char, i-1, j, k+1) || dfs(board, char, i, j+1, k+1) || dfs(board, char, i, j-1, k+1)
	board[i][j] = char[k]
	return res
}

//-------------机器人的运动范围---------------
var m, n, k int
var visited [][]bool

func movingCount(a int, b int, c int) int {
	m = a
	n = b
	k = c
	visited = make([][]bool, a)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, b)
	}
	return dfsMc(0, 0, 0, 0)
}

/**
* si 行数位之和
* sj 列数位之和
 */
func dfsMc(i int, j int, si int, sj int) int {
	//边界判断
	if i >= m || j >= n || k < si+sj || visited[i][j] {
		return 0
	}
	var (
		ai, aj int
	)
	visited[i][j] = true //标记已经访问过了
	//向右移动
	if (i+1)%10 != 0 {
		ai = si + 1
	} else {
		ai = si - 8
	}
	//向下移动
	if (j+1)%10 != 0 {
		aj = sj + 1
	} else {
		aj = sj - 8
	}
	return 1 + dfsMc(i, j+1, si, aj) + dfsMc(i+1, j, ai, sj)
}

//-------------机器人的运动范围---------------
//广度优先搜索算法
func movingCount2(m int, n int, k int) int {

	queue := list.New()
	res := 0
	queue.PushBack([]int{0, 0, 0, 0})
	visited = make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}
	//迭代终止条件： queue 为空
	for queue.Len() > 0 {
		x := queue.Remove(queue.Front()).([]int)
		i, j, si, sj := x[0], x[1], x[2], x[3]
		//判断是否跳过
		if i >= m || j >= n || k < si+sj || visited[i][j] {
			continue
		}
		visited[i][j] = true
		var (
			ai, aj int
		)
		//向右移动(求数位之和)
		if (i+1)%10 != 0 {
			ai = si + 1
		} else {
			ai = si - 8
		}
		//向下移动(求数位之和)
		if (j+1)%10 != 0 {
			aj = sj + 1
		} else {
			aj = sj - 8
		}
		//将当前单元格的右边，下边的坐标，数位和推入queue
		queue.PushBack([]int{i, j + 1, si, aj})
		queue.PushBack([]int{i + 1, j, ai, sj})
		res++
	}
	return res
}

//判断图是否有环
//广度优先搜索
func canFinish(numCourses int, nums [][]int) bool {
	var (
		//定义有向图结构
		edges = make([][]int, numCourses)
		//定义节点入度情况
		indges = make([]int, numCourses)

		result []int
	)

	for _, v := range nums {
		edges[v[1]] = append(edges[v[1]], v[0])
		indges[v[0]]++
	}

	//将入度为0的加入队列
	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indges[i] == 0 {
			q = append(q, i)
		}
	}

	//从队列中取值
	for len(q) > 0 {
		used := q[0]
		q = q[1:]
		result = append(result, used)
		for _, v := range edges[used] {
			//与节点关联的节点入度都-1
			indges[v]--
			if indges[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}

//树的子结构
//前序DLR、后序LRD、中序遍历LDR
//层序遍历(利用辅助栈)
//输入两棵二叉树A和B，判断B是不是A的子结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//先序遍历树 A 中的每个节点 nA
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

//判断树 A 中 以 nA为根节点的子树 是否包含树 B
func recur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}

	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}

//-------------二叉树的镜像--------------
//请完成一个函数，输入一个二叉树，该函数输出它的镜像。
//递归法
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(tmp)
	return root
}

//辅助栈（或队列）*****
//递归法和辅助栈有时可以替换使用
func mirrorTree2(root *TreeNode) *TreeNode {
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		//将左右分支推入stack
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		//交换分支
		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp
	}
	return root
}

// ----------对称的二叉树----------
//请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的
//参考 树的子结构
func isSymmetric(root *TreeNode) bool {
	return root == nil || recur2(root.Left, root.Right)
}
func recur2(L *TreeNode, R *TreeNode) bool {
	if L == nil && R == nil {
		return true
	}

	if L == nil || R == nil || L.Val != R.Val {
		return false
	}

	return recur2(L.Left, R.Right) && recur2(L.Right, R.Left)
}

//----------从上到下打印二叉树-------------
//辅助栈（或队列）*****
//【层序遍历】
func levelOrder(root *TreeNode) []int {
	//为空也要判断
	if root == nil {
		return []int{}
	}
	queue := list.New()
	queue.PushBack(root)

	res := make([]int, 0)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		//将左右分支推入stack
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
		res = append(res, node.Val)

	}
	return res
}

//从上到下打印二叉树 II
//从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
func levelOrder2(root *TreeNode) [][]int {
	queue := list.New()
	res := make([][]int, 0)

	if root != nil {
		queue.PushBack(root)
	}

	for queue.Len() > 0 {
		tmp := make([]int, 0)
		//把当前层的所有节点都推入队列中(利用先进先出的特性)
		for i := queue.Len(); i > 0; i-- {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
	}
	return res
}

//从上到下打印二叉树 III
//请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，
//第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推
func levelOrder3(root *TreeNode) [][]int {
	queue := list.New()
	res := make([][]int, 0)

	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() > 0 {
		tmp := make([]int, 0)
		//把当前层的所有节点都推入队列中
		for i := queue.Len(); i > 0; i-- {
			var node *TreeNode
			if (len(res)+1)%2 == 0 {
				node = queue.Remove(queue.Back()).(*TreeNode)
				if node.Right != nil {
					queue.PushFront(node.Right)
				}
				if node.Left != nil {
					queue.PushFront(node.Left)
				}
			} else {
				node = queue.Remove(queue.Front()).(*TreeNode)
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
				if node.Right != nil {
					queue.PushBack(node.Right)
				}
			}
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
	}
	return res
}

//剑指 Offer 34. 二叉树中和为某一值的路径
//给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//叶子节点 是指没有子节点的节点。
var pathSumPath = make([]int, 0)
var pathSumRes = make([][]int, 0)

func pathSum(root *TreeNode, target int) [][]int {
	//清除数据
	pathSumRes = make([][]int, 0)
	if root == nil {
		return pathSumRes
	}
	recur3(root, target)
	return pathSumRes
}
func recur3(root *TreeNode, tar int) {
	//为空判断很重要
	if root == nil {
		return
	}
	pathSumPath = append(pathSumPath, root.Val)
	tar -= root.Val
	if tar == 0 && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(pathSumPath), cap(pathSumPath))
		//一定要copy
		copy(tmp, pathSumPath)
		pathSumRes = append(pathSumRes, tmp)
	}
	//递归
	recur3(root.Left, tar)  //先走左边分支
	recur3(root.Right, tar) //再走右边分支
	pathSumPath = pathSumPath[:len(pathSumPath)-1]
}

// 二叉搜索树与双向链表
//【二叉搜索树】: 左边的分支节点永远比根节点小，右边的分支节点永远比根节点大
//输入一棵二叉搜索树，将该二叉搜索树转换成一个【排序】的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var pre, head *TreeNode
	//【中序遍历】为对二叉树作 “左、根、右” 顺序遍历
	var dfs func(cur *TreeNode)
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		dfs(cur.Left)
		//调整节点指向
		if pre == nil {
			head = cur
		} else {
			pre.Right = cur
		}
		cur.Left = pre
		//向下移动
		pre = cur
		dfs(cur.Right)
	}
	dfs(root)
	//头尾相接
	head.Left = pre
	pre.Right = head
	return head
}

//序列化二叉树
//请实现两个函数，分别用来序列化和反序列化二叉树。
//层序遍历
func serialzie(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	res := "["
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		if node != nil {
			res += fmt.Sprint(node.Val) + "," //拼接字符串
			queue.PushBack(node.Left)
			queue.PushBack(node.Right)
		} else {
			res += "null,"
		}
	}
	res = res[:len(res)-1]
	res += "]"
	return res
}

//反序列化二叉树
//【辅助队列】
func deserialzie(data string) *TreeNode {
	if data == "[]" {
		return nil
	}
	vars := strings.Split(data[1:len(data)-1], ",")
	tmpInt, _ := strconv.Atoi(vars[0])
	root := &TreeNode{tmpInt, nil, nil}
	queue := list.New()
	queue.PushBack(root)
	i := 1 //移动字符串
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		if vars[i] != "null" {
			tmpInt, _ := strconv.Atoi(vars[i])
			node.Left = &TreeNode{tmpInt, nil, nil}
			queue.PushBack(node.Left)
		}
		i++ //往后移动字符串
		if vars[i] != "null" {
			tmpInt, _ := strconv.Atoi(vars[i])
			node.Right = &TreeNode{tmpInt, nil, nil}
			queue.PushBack(node.Right)
		}
		i++
	}
	return root
}

// 字符串的排列
//输入一个字符串，打印出该字符串中字符的所有排列。f(n) = 1*2*3...(n-2)*(n-1)*n
var Res_permutation []string
var C []string

func permutation(s string) []string {
	Res_permutation = []string{}
	C = strings.Split(s, "")
	dfs_permutation(0)
	return Res_permutation
}
func dfs_permutation(x int) {
	//如果到最后一个字符则跳出
	if x == len(C)-1 {
		cc := ""
		for _, char := range C {
			cc += char
		}
		Res_permutation = append(Res_permutation, cc)
	}
	set := make(map[string]string, 0) //记录重复字符
	for i := x; i < len(C); i++ {
		//如果重复则跳过
		if set[C[i]] == C[i] {
			continue
		}
		set[C[i]] = C[i]
		// 交换，将 c[i] 固定在第 x 位
		Swap(i, x)
		//固定下一个位置
		dfs_permutation(x + 1)
		//回溯/撤销
		Swap(i, x)
	}
}
func Swap(a, b int) {
	tmp := C[a]
	C[a] = C[b]
	C[b] = tmp
}

//二叉搜索树的第 k 大节点
//中序遍历(倒序 RDL)
func kthLargest(root *TreeNode, k int) int {
	res := 0
	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode) //定义一个函数变量
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		if k == 0 {
			return
		}
		if k = k - 1; k == 0 {
			res = node.Val
		}
		dfs(node.Left)
	}
	dfs(root)
	return res
}

//二叉树的深度
//【后序遍历】此树的深度 等于 左子树的深度 与 右子树的深度 中的 最大值 +1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

//【层序遍历】
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	res := 0
	for queue.Len() > 0 {
		tmp := list.New()
		for i := queue.Len(); i > 0; i-- {
			node := queue.Remove(queue.Back()).(*TreeNode)
			if node.Left != nil {
				tmp.PushBack(node.Left)
			}
			if node.Right != nil {
				tmp.PushBack(node.Right)
			}
		}
		queue = tmp
		res++
	}
	return res
}

//平衡二叉树
//输入一棵二叉树的根节点，判断该树是不是平衡二叉树。
//如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树
//先序遍历 + 判断深度 （从顶至底）
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//先序遍历 DLR(当前子树，左子树的子树，右子树的子树)
	return Abs(maxDepth(root.Left)-maxDepth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

//后序遍历 + 剪枝 （从底至顶）
func isBalanced2(root *TreeNode) bool {

	var recur func(node *TreeNode) int
	recur = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		//左节点
		left := recur(node.Left)

		if left == -1 {
			return -1
		}
		//右节点
		right := recur(node.Right)

		if right == -1 {
			return -1
		}
		//当前节点
		if Abs(left-right) <= 1 {
			//返回节点的深度
			return Max(left, right) + 1
		} else {
			return -1
		}
	}
	return recur(root) != -1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//求 1 + 2 + … + n
//要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）
//【递归】+(逻辑运算符的短路效应)
func sumNums(n int) int {
	res := 0
	var recur func(a int, x bool) int
	recur = func(a int, x bool) int {
		x = a > 1 && recur(a-1, x) > 0
		res += a
		return res
	}
	recur(n, false)
	return res
}

//I. ---------------二叉搜索树的最近公共祖先------------------
//给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先
//【迭代】
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//保持 p.Val < q.Val 方便后面的判断
	if p.Val > q.Val {
		temp := q
		p = q
		q = temp
	}
	for root != nil {
		//如果当前节点的值大于 q则走左子树
		if root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val {
			root = root.Right
		} else {
			break
		}
	}
	return root
}

//【递归】
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	//保持 p.Val < q.Val 方便后面的判断
	if p.Val > q.Val {
		temp := q
		p = q
		q = temp
	}
	if root.Val > q.Val {
		return lowestCommonAncestor2(root.Left, p, q)
	} else if root.Val < p.Val {
		return lowestCommonAncestor2(root.Right, p, q)
	}
	return root
}

//II. 二叉树的最近公共祖先
//给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//先序遍历
func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor3(root.Left, p, q)
	right := lowestCommonAncestor3(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

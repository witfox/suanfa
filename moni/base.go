package moni

import "container/list"

/*
*模拟
 */

//顺时针打印矩阵
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	//标记上下左右边界
	row := len(matrix)
	column := len(matrix[0])
	l, r, t, b := 0, column-1, 0, row-1
	res := make([]int, column*row)
	x := 0
	for {
		//从左向右走
		for i := l; i <= r; i++ {
			res[x] = matrix[t][i]
			x++
		}
		t++ //裁剪掉已经遍历的行
		if t > b {
			break
		}
		//从上到下
		for i := t; i <= b; i++ {
			res[x] = matrix[i][r]
			x++
		}
		r-- //裁剪掉右边
		if r < l {
			break
		}
		//从右到左
		for i := r; i >= l; i-- {
			res[x] = matrix[b][i]
			x++
		}
		b-- //裁剪掉下边
		if b < t {
			break
		}
		//从下到上
		for i := b; i >= t; i-- {
			res[x] = matrix[i][l]
			x++
		}
		l++ //裁剪掉已经遍历的行
		if l > r {
			break
		}
	}
	return res
}

//栈的压入、弹出序列
//辅助栈
func validateStackSequences(pushed []int, popped []int) bool {
	stack := list.New()
	i := 0
	for _, num := range pushed {
		//入栈
		stack.PushBack(num)
		for stack.Len() > 0 && stack.Back().Value == popped[i] {
			stack.Remove(stack.Back())
			i++
		}
	}
	return stack.Len() == 0
}

package gosuanfa

import (
	"container/list"
	"fmt"
	"math"
	"strings"
)

//指数 O(2N次方)
func zhishu(N int) int {
	if N <= 0 {
		return 1
	}
	count_1 := zhishu(N - 1)
	count_2 := zhishu(N - 1)
	return count_1 + count_2
}

//阶乘 O(N!)
func jiecheng(N int) int {
	if N <= 0 {
		return 1
	}
	count := 0
	for i := 0; i < N; i++ {
		count += jiecheng(N - 1)
	}
	return count
}

//两数之和  辅助哈希表 O(N)
func twoSum(nums []int, target int) []int {
	dic := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		in := target - nums[i]
		_, ok := dic[in]

		if ok {
			return []int{dic[in], i}
		}

		dic[nums[i]] = i
	}

	return []int{}
}

/**
* 左旋转字符串
* 字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
* 请定义一个函数实现字符串左旋转操作的功能。
* 比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
* 解题：字符串切片
 */
func reverseLeftWords(s string, k int) string {

	//字符串切片
	//return s[k:] + s[:k]

	//复制
	str := s + s
	l := len(s) + k
	return str[k:l]
}

//滑动窗口的最大值
//给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值
func maxSlidingWindow(nums []int, k int) []int {

	queue := list.New() //辅助队列,存储最大值
	res := []int{}
	i := 1 - k
	for j := 0; j < len(nums); j++ {

		//如果队首元素和移动框被删除的元素（nums[i-1]）相同则删掉
		if i > 0 && queue.Front().Value.(int) == nums[i-1] {
			queue.Remove(queue.Front())
		}
		//保持队首元素为最大值
		for queue.Len() > 0 && queue.Back().Value.(int) < nums[j] {
			queue.Remove(queue.Back())
		}
		queue.PushBack(nums[j])
		fmt.Println(queue.Front().Value)

		//窗口形成，开始记录最大值
		if i >= 0 {
			res = append(res, queue.Front().Value.(int))
		}

		i++
	}
	return res

}

//把字符串转换成整数
//写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。
//总结：越界的判断， ASCII码处理
func strToInt(str string) int {
	//去掉空格
	str = strings.TrimSpace(str)
	if len(str) <= 0 {
		return 0
	}
	var res, i, sign int = 0, 1, 1

	//设置边界值
	maxInt, minInt, bndry := math.MaxInt32, math.MinInt32, math.MaxInt32/10

	//首字符符号判断
	if str[0:1] == "-" {
		sign = -1
	} else if str[0:1] != "+" {
		i = 0
	}

	for _, char := range []rune(str[i:]) {
		//如果为非数字则直接退出 ASCII码判断
		if char < '0' || char > '9' {
			break
		}
		//数字越界
		if res > bndry || res == bndry && char > '7' {
			if sign == 1 {
				return maxInt
			}
			return minInt
		}
		//此数字的 ASCII 码” 与 “ 00 的 ASCII 码” 相减即可
		res = res*10 + int((char - '0'))
	}
	return res * sign

}

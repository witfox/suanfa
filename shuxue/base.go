package shuxue

import (
	"math"
	"strconv"
)

//给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
//请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	}
	if b == 1 {
		return int(math.Pow(3, float64(a-1)) * 4)
	}
	return int(math.Pow(3, float64(a)) * 2)
}

//剪绳子 II
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
// func cuttingRope2(n int) int {

// }

//数组中出现次数超过一半的数字
//数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
//“数组中出现次数超过一半的数字” 简称为 “众数” 。
//【摩尔投票法】： 核心理念为 票数正负抵消 。此方法时间和空间复杂度分别为 O(N)O(N) 和 O(1)O(1) ，为本题的最佳解法。
func majorityElement(nums []int) int {
	x, votes := 0, 0
	for _, num := range nums {
		if votes == 0 {
			x = num
		}
		if num == x {
			votes += 1
		} else {
			votes += -1
		}
	}
	return x
}

//1～n 整数中 1 出现的次数
func countDigitOne(n int) int {
	digit := 1     //位数
	low := 0       //低位
	high := n / 10 //高位
	cur := n % 10  //当前位
	res := 0
	for high != 0 || cur != 0 {
		if cur == 0 {
			res += high * digit
		} else if cur == 1 {
			res += high*digit + low + 1
		} else {
			res += (high + 1) * digit
		}
		//向高位推进
		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}
	return res
}

//数字序列中某一位的数字
//数字以0123456789101112131415…的格式序列化到一个字符序列中。
//在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。
func findNthDigit(n int) int {
	digit := 1 // 确定n所在数字的位数
	start := 1 //数字的起始数字  1，10，100。。。
	count := 9 //第一位的数字数量

	//n所处的数字位数
	for n > count {
		n -= count
		digit++
		start *= 10
		count = start * digit * 9 //位数对应的数字个数
	}
	//n所在的数字
	num := start + (n-1)/digit
	//n所在数字的哪个字符
	res := strconv.Itoa(num)[(n-1)%digit] - '0'
	return int(res)
}

//和为 s 的连续正数序列
//【双指针】
func findContinuousSequence(target int) [][]int {
	i, j := 1, 2 //双指针
	s := 3       //前后指针和
	res := make([][]int, 0)
	//如果i > j 跳出循环
	for i < j {
		if s == target {
			tmp := make([]int, 0)
			for k := i; k <= j; k++ {
				tmp = append(tmp, k)
			}
			res = append(res, tmp)
		}
		//移动指针
		if s < target {
			j++
			s += j
		} else {
			s -= i
			i++
		}
	}
	return res
}

//圆圈中最后剩下的数字
//本题是著名的 “约瑟夫环” 问题，可使用 动态规划 解决。 ???
//动态规划 f(n) = (f(n-1) + m) % n
func lastRemaining(n int, m int) int {
	x := 0 //代表 f(n)

	//f(1)恒等于0
	for i := 2; i <= n; i++ {
		x = (x + m) % i
	}
	return x
}

//构建乘积数组
//给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中 <除了下标 i 以外> 的元素的积, 即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。
//不能使用除法。
//根据”除了下标 i 以外“，列出【表格】，得出对角线数字都为1，划分上半三角形和下半三角形
func constructArr(a []int) []int {
	len := len(a)
	if len == 0 {
		return []int{}
	}
	b := make([]int, len)
	b[0] = 1
	tmp := 1

	//先计算下半三角形的乘积
	for i := 1; i < len; i++ {
		b[i] = b[i-1] * a[i-1]
	}
	//计算下半三角形
	for i := len - 2; i >= 0; i-- {
		tmp *= a[i+1]
		b[i] *= tmp
	}
	return b
}

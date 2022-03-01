package dongtaiguihua

import "math"

//动态规划
//step1:状态定义（动态规划矩阵）
//step2:初始状态 （初始化矩阵）
//step3:转移方程
//step4:返回值

//---------------------求第 n 个斐波那契数「重叠子问题」------------------------------
// f(n+1) = f(n)+f(n-1); 0 1 1 2 3 5
//记忆递归法（由上及下）O(N)O(N)
func fibonacci2(n int, dp []int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	//存储重复计算的值
	if dp[n] != 0 {
		return dp[n]
	}
	dp[n] = fibonacci2(n-1, dp) + fibonacci2(n-2, dp)

	return dp[n]
}

//动态规划(由下及上) 空间复杂度O(1)
func fibonacci(n int) int {

	if n == 0 {
		return 0
	}
	a, b := 0, 1 //初始化 f(0) f(1)

	for i := 2; i <= n; i++ {
		tmp := a
		a = b
		b = tmp + b
	}
	return b
}

//--------------------蛋糕最高价-------------------
//动态规划，本题同时包含「重叠子问题」和「最优子结构」
//n为蛋糕的总重量，priceList每个重量对应的价格 f(n)=max(f(i)+p(n−i));0≤i<n
func max_cake_price(n int, priceList []int) int {

	if n == 0 {
		return 0
	}
	//存储各个重量的最高价格
	var dp = make([]int, n+1)
	//n拆分
	for i := 1; i <= n; i++ {
		//子节点再拆分
		for j := 0; j < i; j++ {
			//求子节点的最大值
			dp[i] = Max(dp[i], dp[j]+priceList[i-j])
		}
	}
	return dp[n]

}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//----------------------青蛙跳台阶问题--------------------
//一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
//同斐波那契数列  f(n) = f(n-1)+f(n-2); 1 1 2 3
func numWays(n int) int {

	a, b := 1, 1
	for i := 2; i <= n; i++ {
		sum := (a + b) % 1000000007
		a = b
		b = sum
	}
	return b
}

//---------------正则表达式匹配----------------------
//请实现一个函数用来匹配包含'. '和'*'的正则表达式。
//模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。
//在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配
func isMatch(s string, p string) bool {
	m, n := len(s)+1, len(p)+1
	dp := make([][]bool, m)
	//go二维数组需要再定义，分配空间
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	//初始化首行
	for j := 2; j < n; j += 2 {
		dp[0][j] = dp[0][j-2] && p[j-1:j] == "*"
	}
	//循环矩阵
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if p[j-1:j] == "*" {
				if dp[i][j-2] {
					dp[i][j] = true
				} else if dp[i-1][j] && s[i-1:i] == p[j-2:j-1] {
					dp[i][j] = true
				} else if dp[i-1][j] && p[j-2:j-1] == "." {
					dp[i][j] = true
				}
			} else {
				//判断第i,第j之前的情况和当前的情况
				if dp[i-1][j-1] && s[i-1:i] == p[j-1:j] {
					dp[i][j] = true
				} else if dp[i-1][j-1] && p[j-1:j] == "." {
					dp[i][j] = true
				}
			}
		}
	}
	return dp[m-1][n-1]
}

//剑指 Offer 42. 连续子数组的最大和 2022-02-04
//输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值
//要求时间复杂度为O(n)
func maxSubArray(nums []int) int {

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] += Max(nums[i-1], 0)
		res = Max(res, nums[i])
	}
	return res
}

//剑指 Offer 46. 把数字翻译成字符串 2022-02-04
//给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。
//一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
//时间O(N); 空间O(1)
//dp[i] = dp[i-1] + d[i-2] || dp[i-1]
func translateNum(num int) int {
	a, b := 1, 1 //dp[i-1],dp[i],从右往左移动
	y := num % 10
	var c int
	for num > 9 {
		num = num / 10 //往前移动
		x := num % 10
		temp := 10*x + y
		if temp >= 10 && temp <= 25 {
			c = a + b
		} else {
			c = a
		}
		//往前移动
		b = a
		a = c
		y = x

	}
	return a
}

//剑指 Offer 47. 礼物的最大价值 2022-02-04
//在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）
//你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角
//给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
//f(i,j) = Max(f(i,j-1),f(i-1,j))+grid(i,j)
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 && j != 0 { //只右边走
				grid[i][j] += grid[i][j-1]
			} else if i != 0 && j == 0 { //只左边走
				grid[i][j] += grid[i-1][j]
			} else {
				//可左可右
				grid[i][j] = grid[i][j] + Max(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}

//剑指 Offer 48. 最长不含重复字符的子字符串
//请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
func lengthOfLongestSubstring(s string) int {

	temp := 0                    //记录dp[j]
	res := 0                     //记录结果值
	disc := make(map[string]int) //记录字符对应的index
	for j := 0; j < len(s); j++ {
		i, ok := disc[s[j:j+1]] //i记录左边距离最近的重复字符
		if !ok {
			i = -1
		}
		disc[s[j:j+1]] = j //更新哈希表
		//i在dp[j-1]区间内，j-i
		if temp >= j-i {
			temp = j - i
		} else {
			//i在dp[j-1]区间外 dp[j-1]+1
			temp = temp + 1
		}
		res = Max(temp, res)
	}
	return res
}

//双指针 + 哈希表
func lengthOfLongestSubstring2(s string) int {
	res := 0
	disc := make(map[string]int)
	for j := 0; j < len(s); j++ {
		i, ok := disc[s[j:j+1]] //左边距离最近的重复字符的位置
		if ok {
			i = Max(i, disc[s[j:j+1]])
		} else {
			i = -1
		}
		disc[s[j:j+1]] = j
		res = Max(res, j-i)
	}
	return res
}

//剑指 Offer 49. 丑数
//我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数
//转移方程：下一个丑数=某个丑数*某个因子；即 min(dp[a]*2,dp[b]*3,dp[c]*5)
func nthUglyNumber(n int) int {
	a, b, c := 0, 0, 0   //三个因子的指标
	dp := make([]int, n) //第i个丑数
	dp[0] = 1
	for i := 1; i < n; i++ {
		n1, n2, n3 := 2*dp[a], 3*dp[b], 5*dp[c]
		dp[i] = Min(Min(n1, n2), n3)
		if dp[i] == n1 {
			a++
		}
		if dp[i] == n2 {
			b++
		}
		if dp[i] == n3 {
			c++
		}
	}
	return dp[n-1]
}

//剑指 Offer 60. n 个骰子的点数
//把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
func dicesProbability(n int) []float64 {

	//初始化n=1的概率
	dp := []float64{1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0}

	for i := 2; i <= n; i++ {
		tmp := make([]float64, 5*i+1)
		for j := 0; j < len(dp); j++ {
			for k := 0; k < 6; k++ {
				tmp[j+k] += dp[j] / 6.0
			}
		}
		dp = tmp //dp,tmp交替前进
	}
	return dp
}

//股票的最大利润
//假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
//转移方程： max(f(i-1), prices[i]-min(prices[0:i]))
func maxProfit(prices []int) int {
	cost := math.MaxInt
	res := 0 //记录当前最大收益
	for i := 0; i < len(prices); i++ {
		cost = Min(cost, prices[i]) //记录最小值
		res = Max(res, prices[i]-cost)
	}
	return res
}

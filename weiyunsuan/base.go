package weiyunsuan

/*
位运算
&,|,^,>>,<<
*/
//编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为 汉明重量).）
func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		res += int(num) & 1
		num >>= 1 //右移一位
	}
	return res
}

// 数组中数字出现的次数
//一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。
//请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
func singleNumbers(nums []int) []int {
	res := make([]int, 0)
	x, y, n, m := 0, 0, 0, 1
	for _, num := range nums {
		//遍历 与或
		n ^= num //得到两个不同值的与或 a ^= b
	}
	// m 循环左移一位，直到 z & m ！= 0
	for (n & m) == 0 {
		m <<= 1
	}
	//拆分 nums 为两个子数组
	for _, num := range nums {
		if (num & m) == 0 {
			x ^= num // 若 num & m == 0 , 划分至子数组 1 ，执行遍历异或
		} else {
			y ^= num // 若 num & m != 0 , 划分至子数组 2 ，执行遍历异或
		}
	}
	res = append(res, x)
	res = append(res, y)
	return res

}

//数组中数字出现的次数
//在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。
// func singleNumber(nums []int) int {

// }

//不用加减乘除做加法
func add(a int, b int) int {
	for b != 0 {
		c := (a & b) << 1 //a+b进位和
		a ^= b            //a+b非进位和
		b = c             //直到b==0
	}
	return a
}

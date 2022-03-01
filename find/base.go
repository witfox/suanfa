package find

/*
查找算法
“排序数组的查找问题首先”考虑使用 【二分法】 解决，其可将 【遍历法】 的 线性级别 时间复杂度降低至 对数级别
*/
// 数组中重复的数字
//在一个长度为 n 的数组 nums 里的所有数字都在【 0～n-1 】的范围内
//原地交换   时间O(N)，空间O(1)
func findRepeatNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i++
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		//交换对应索引的值
		tmp := nums[i]
		nums[i] = nums[tmp]
		nums[tmp] = tmp
	}
	return -1
}

//二维数组中的查找
//在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
//请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数
//类二叉搜索树
func findNumberIn2DArray(matrix [][]int, target int) bool {
	//以左下角的数字为标识位
	i := len(matrix) - 1
	j := 0

	for i >= 0 && j < len(matrix[0]) {
		if target > matrix[i][j] {
			j++
		} else if target < matrix[i][j] {
			i--
		} else {
			return true
		}
	}
	return false

}

//旋转数组的最小数字
//二分法
func minArray(numbers []int) int {
	i, j := 0, len(numbers)-1
	for i < j {
		m := (i + j) / 2
		if numbers[m] > numbers[j] {
			i = m + 1
		} else if numbers[m] < numbers[j] {
			j = m
		} else {
			j--
		}
	}
	return numbers[i]
}

//第一个只出现一次的字符
//在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
//golang map会自动排序
func firstUniqChar(s string) byte {
	char := [26]int{}
	for _, v := range s {
		char[v-'a']++
	}
	for i, v := range s {
		if char[v-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

//在排序数组中查找数字 I
//统计一个数字在排序数组中出现的次数
func search(nums []int, target int) int {

	//右边界-左边界值
	return helper(nums, target) - helper(nums, target-1)
}
func helper(nums []int, target int) int {
	i, j := 0, len(nums)-1

	for i <= j {
		m := (i + j) / 2
		if nums[m] <= target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}

// 0～n-1 中缺失的数字
//一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
//在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
//【二分法查找】
func missingNumber(nums []int) int {
	i, j := 0, len(nums)-1

	for i <= j {
		m := (i + j) / 2
		//则左边不缺少
		if nums[m] == m {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}

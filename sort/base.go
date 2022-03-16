package sort

import (
	"container/heap"
	"sort"
	"strconv"
	"strings"
)

/*
排序
「冒泡排序」、「插入排序」、「选择排序」、「快速排序」、「归并排序」、「堆排序」、「基数排序」、「桶排序」
*/

//冒泡排序
func popSort(nums []int) []int {

	for i := 0; i < len(nums)-1; i++ { //外循环
		flag := false
		for j := i; j < len(nums)-1-i; j++ { //内循环
			if nums[j] > nums[j+1] {
				// 交换 nums[j], nums[j + 1]
				swap(nums, j, j+1)
				flag = true
			}
		}
		if !flag {
			break // 内循环未交换任何元素，则跳出
		}
	}
	return nums
}

//插入排序
func insertSort(nums []int) []int {
	//从第二个开始比较
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] > nums[j-1] {
				//交换
				swap(nums, j, j-1)
			}
		}
	}
	return nums
}

//选择排序
func selectSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		k := i
		//从末尾开始比较
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[k] {
				k = j
			}
		}
		swap(nums, k, i)
	}
	return nums
}

//快速排序
//快速排序算法有两个核心点，分别为 哨兵划分 和 递归 。
//以某个数为基准，将大于基准数的数值放在右边，将小于基准数的数放在左边
func quickSort(nums []int) []int {

	var recur func(left, right int)
	recur = func(left, right int) {
		//终止条件
		if left >= right {
			return
		}
		i := partition(nums, left, right)
		recur(left, i-1)
		recur(i+1, right)
	}
	recur(0, len(nums)-1)
	return nums
}

//哨兵划分操作
func partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		swap(nums, i, j)
	}
	//更换基数
	swap(nums, i, left)
	return i
}
func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

//归并排序
//归并排序体现了 “分而治之” 的算法思想
func mergeSort(nums []int) []int {
	numsLen := len(nums)
	tmp := make([]int, numsLen)

	var recur func(left, right int)
	recur = func(left, right int) {
		//终止条件,只剩下单独的元素
		if left >= right {
			return
		}
		//找到中分索引，进行递归划分
		m := (left + right) / 2
		recur(left, m)
		recur(m+1, right)

		//暂存合并区间的元素
		for k := left; k <= right; k++ {
			tmp[k] = nums[k]
		}
		//两指针分别指向左/右子数组的首个元素
		i, j := left, m+1
		//合并子数组
		for k := left; k <= right; k++ {
			if i == m+1 {
				nums[k] = tmp[j]
				j++
			} else if j == right+1 || tmp[i] <= tmp[j] {
				nums[k] = tmp[i]
				i++
			} else {
				nums[k] = tmp[j]
				j++
			}
		}
	}
	recur(0, numsLen-1)
	return nums
}

//最小的 k 个数
//输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
//基于快速排序的数组划分
func getLeastNumbers(arr []int, k int) []int {
	var quickSort func(left, right int) []int

	quickSort = func(left, right int) []int {

		//对比基数arr[left]
		i, j := left, right
		for i < j {
			for i < j && arr[j] >= arr[left] {
				j--
			}
			for i < j && arr[i] <= arr[left] {
				i++
			}
			swap(arr, i, j)
		}
		//更换基数
		swap(arr, i, left)
		if i > k {
			//代表第 k + 1 小的数字在 左子数组 中，则递归左子数组
			quickSort(left, i-1)
		}
		if i < k {
			quickSort(i+1, right)
		}
		//如果k==i，则代表此时 arr[k] 即为第 k + 1k+1 小的数字，则直接返回数组前 kk 个数字即可
		return arr[0:k]
	}
	if k >= len(arr) {
		return arr
	}
	return quickSort(0, len(arr)-1)
}

//数据流中的中位数
//堆排序
type MedianFinder struct {
	A *minHeap //小顶堆
	B *maxHeap //大顶堆
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	minH := &minHeap{}
	maxH := &maxHeap{}
	heap.Init(minH)
	heap.Init(maxH)
	return MedianFinder{
		A: minH,
		B: maxH,
	}
}

func (this *MedianFinder) AddNum(num int) {
	//总数为奇数
	if this.A.Len() != this.B.Len() {
		heap.Push(this.A, num)
		heap.Push(this.B, heap.Pop(this.A))
	} else {
		heap.Push(this.B, num)
		heap.Push(this.A, heap.Pop(this.B))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.A.Len() != this.B.Len() {
		return float64(this.A.Peek())
	} else {
		return (float64(this.A.Peek()) + float64(this.B.Peek())) / 2.0
	}
}

//把数组排成最小的数
//输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
//快速排序
func minNumber(nums []int) string {
	//转为string数组
	strNums := make([]string, len(nums))
	for k, v := range nums {
		strNums[k] = strconv.Itoa(v)
	}
	var recur func(left, right int)
	recur = func(left, right int) {
		if left >= right {
			return
		}
		//比对基数strNums[left]
		i, j := left, right
		tmp := strNums[i]
		for i < j {
			for i < j && (strNums[j]+strNums[left]) >= (strNums[left]+strNums[j]) {
				j--
			}
			for i < j && (strNums[i]+strNums[left]) <= (strNums[left]+strNums[i]) {
				i++
			}
			tmp = strNums[i]
			strNums[i] = strNums[j]
			strNums[j] = tmp
		}
		strNums[i] = strNums[left]
		strNums[left] = tmp
		recur(left, i-1)
		recur(i+1, right)
	}
	recur(0, len(nums)-1)
	return strings.Join(strNums, "")
}

//扑克牌中的顺子
//从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的
//原理：max-min < 5 且不存在重复数字，则为顺子
//方法一：max-min < 5 则为顺子。且数字不重复
//方法二：排序+遍历
func isStraight(nums []int) bool {
	sort.Ints(nums)
	jokers := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			jokers++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[jokers] < 5
}

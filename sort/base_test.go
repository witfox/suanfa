package sort

import (
	"fmt"
	"testing"
)

func TestPopSort(t *testing.T) {

	nums := []int{4, 1, 3, 1, 5, 2}
	res := popSort(nums)
	fmt.Println(res)
}

func TestQuickSort(t *testing.T) {

	nums := []int{1, 0, 2, 4, 3, 5}
	res := quickSort(nums)
	fmt.Println(res)
}

func TestInsertSort(t *testing.T) {

	nums := []int{1, 0, 2, 4, 3, 5}
	res := insertSort(nums)
	fmt.Println(res)
}
func TestSelectSort(t *testing.T) {

	nums := []int{1, 0, 2, 4, 3, 5}
	res := selectSort(nums)
	fmt.Println(res)
}
func TestMergeSort(t *testing.T) {

	nums := []int{1, 0, 2, 4, 3, 5}
	res := mergeSort(nums)
	fmt.Println(res)
}

func TestMedianFinder(t *testing.T) {

	obj := Constructor()
	obj.AddNum(5)
	obj.AddNum(2)
	obj.AddNum(4)
	obj.AddNum(3)
	obj.AddNum(1)
	res := obj.FindMedian()
	fmt.Println(res)
}

func TestMinNumberr(t *testing.T) {
	nums := []int{3, 30, 34, 9, 5}
	res := minNumber(nums)
	fmt.Println(res)
}

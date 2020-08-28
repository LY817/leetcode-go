package main

import "fmt"

// 冒泡排序
// 将最大的元素放到集合的最后
func bubbleSort(arr []int) {
	fmt.Println(arr)
	// 每次缩小遍历范围
	for i := len(arr) - 1; i > 0; i-- {
		// 每次遍历进行一次冒泡互换
		for j := 0; j < i; j++ {
			if compareInt(arr[j], arr[j+1]) {
				a := arr[j]
				b := arr[j+1]
				arr[j] = b
				arr[j+1] = a
			}
		}
	}
	fmt.Println(arr)
}

// 选择排序
// 将最小的元素移到集合的开头
func pickSort(arr []int) {
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if !compareInt(arr[i], arr[j]) {
				a := arr[i]
				b := arr[j]
				arr[j] = a
				arr[i] = b
			}
		}
	}
	fmt.Println(arr)
}

// 插入排序
// 初始状态 左边只有一个元素 作为已排序集合 右边剩下的作为未排序结合
func insertSort(arr []int) {
	fmt.Println(arr)
	// i作为分界index
	for i := 1; i < len(arr); i++ {
		// 插入逻辑
		// 将新的元素并入左边有序集合
		// 反向冒泡
		for j := i; j > 0; j-- {
			if !compareInt(arr[j], arr[j-1]) {
				b := arr[j-1]
				a := arr[j]
				arr[j] = b
				arr[j-1] = a
			}
		}
	}
	fmt.Println(arr)
}

func compareInt(a int, b int) bool {
	if a > b {
		return true
	}
	return false
}

func compare(a int, b int) bool {
	if a > b {
		return true
	}
	return false
}

func exchange() {

}

func main() {
	arr := []int{4, 5, 6, 3, 2, 1}
	bubbleSort(arr)
	// pickSort(arr)
	// insertSort(arr)
}

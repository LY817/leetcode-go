package main

import "fmt"

// 问题
// 在该数组中找出和为目标值的那两个整数，并返回他们的数组下标
func twoSum(nums []int, target int) []int {
	panic("No two sum solution")
}

// 方法一：暴力遍历
// 时间复杂度：O(N^2)	两次遍历数组
// 空间复杂度：O(1) 	不需要额外内存
func twoSumV1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums) && j != i; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	panic("No two sum solution")
}

// 方法二：两次hash
// 1. 第一遍遍历 将nums存入到hashmap中 key为nums[i] value为i
// 2. 第二次遍历 计算target - nums[i]是否在hashmap中存在
func twoSumV2(nums []int, target int) []int {
	panic("No two sum solution")
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSumV1(nums, target))
}

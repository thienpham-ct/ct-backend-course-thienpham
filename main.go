package main

import (
	"fmt"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func twosum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}

	return nil
}

func countPairs(nums []int, target int) int {
	counter := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				counter++
			}
		}
	}
	return counter
}

func countKDifference(nums []int, k int) int {
	counter := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if Abs(nums[i]-nums[j]) == k {
				counter++
			}
		}
	}
	return counter
}

func findMaxK(nums []int) int {
	largestNumber := -1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == 0 && Abs(nums[i]) > largestNumber {
				largestNumber = Abs(nums[i])
			}
		}
	}
	return largestNumber
}

// PERFORMANCE IMPROVE
// Bai code thiu nhi so 1, su dung hashmap
func impv_twosum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for currenIndex, currentNum := range nums {
		mapIndex, ok := hashMap[target-currentNum]
		if ok {
			return []int{currenIndex, mapIndex}
		}
		hashMap[currentNum] = currenIndex
	}
	return nil
}

func impv_countPairs(nums []int, target int) int {
	counter := 0
	hashMap := make(map[int]int)
	for i, num_i := range nums {
		_, ok := hashMap[target-num_i]
		if ok {
			counter++
		}
		hashMap[num_i] = i
	}
	return counter
}

func main() {
	nums := []int{-1, 1, 2, 3, 1}
	target := 2

	//Bài code thiếu nhi 1
	//twosum := impv_twosum(nums, target)
	//fmt.Println("Two sum result:", twosum)

	//Bài code thiếu nhi 2
	countPairs := impv_countPairs(nums, target)
	fmt.Println("Count pair result:", countPairs)

	//Bài code thiếu nhi 3
	//countKDifference := countKDifference(nums, target)
	//fmt.Println("Count pair result:", countKDifference)

	//Bài code thiếu nhi 4
	//findMaxK := findMaxK(nums)
	//fmt.Println("findMaxK result:", findMaxK)
}

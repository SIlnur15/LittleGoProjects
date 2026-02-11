package main

mapIntegers := func(nums []int, callback func(int) int) []int {
    for idx,elem := range nums {
		nums[idx] = callback(elem)
	}
	return nums
}

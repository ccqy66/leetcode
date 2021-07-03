package main

// [1,1,4,2,3], x = 5
// 虽然是从左右两个位置开始找sum=x的值，其中元素最少
// 其实也等效于求：连续子数组sum=sum(nums)-x的最大元素

func minOperations(nums []int, x int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if x > sum {
		return -1
	}
	return 0
	//var max int
	//var start, end = 0, 0
	//target := sum - x
	//
	//for start < len(nums) && end < len(nums) {
	//
	//}
}

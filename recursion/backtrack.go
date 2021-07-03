package recursion

import (
	"fmt"
	"sort"
)

var permuteResult [][]int

func permute(nums []int) [][]int {
	permuteResult = make([][]int, 0)
	permuteBacktrack(nums, 0, len(nums))
	return permuteResult
}

/**
为什么使用start,end两个参数：目的是为了区分选择列表和路径
[0, start)：代表路径
[start,end)：代表选择列表
*/
func permuteBacktrack(nums []int, start, end int) {
	if start >= end {
		dst := make([]int, len(nums))
		copy(dst, nums)
		permuteResult = append(permuteResult, dst)
		return
	}
	for i := start; i < end; i++ {
		nums[i], nums[start] = nums[start], nums[i]
		permuteBacktrack(nums, start+1, end)
		nums[i], nums[start] = nums[start], nums[i]
	}
}

var table = map[uint8]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}
var rs []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	rs = make([]string, 0)
	letterBacktrack(digits, 0, "")
	return rs
}

func letterBacktrack(digits string, index int, result string) {
	if index == len(digits) {
		rs = append(rs, result)
		return
	}
	di := table[digits[index]-'0']
	for i := 0; i < len(di); i++ {
		letterBacktrack(digits, index+1, fmt.Sprintf("%s%s", result, string(di[i])))
	}
}

var sumResult [][]int

func combinationSum(candidates []int, target int) [][]int {
	sumResult = make([][]int, 0)
	item := make([]int, 0)
	backtrackSum(candidates, 0, target, 0, item)
	return sumResult
}

func backtrackSum(candidates []int, start int, target int, sum int, result []int) {
	if sum == target {
		sumResult = append(sumResult, result)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		}
		dst := make([]int, len(result))
		copy(dst, result)
		if sum+candidates[i] > target {
			continue
		}
		dst = append(dst, candidates[i])
		backtrackSum(candidates, i, target, sum+candidates[i], dst)
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sumResult = make([][]int, 0)
	item := make([]int, 0)
	sort.Ints(candidates)
	backtrackSum2(candidates, 0, target, 0, item)
	return sumResult
}

func backtrackSum2(candidates []int, start int, target int, sum int, result []int) {
	if sum == target {
		sumResult = append(sumResult, result)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		}
		if check(candidates, start, i) {
			continue
		}
		dst := make([]int, len(result))
		copy(dst, result)
		if sum+candidates[i] > target {
			continue
		}
		dst = append(dst, candidates[i])
		backtrackSum2(candidates, i+1, target, sum+candidates[i], dst)
	}
}

func check(candidates []int, start , end int) bool{
	for j := start+1; j < end; j++ {
		if candidates[j] == candidates[start] {
			return true
		}
	}
	return false
}
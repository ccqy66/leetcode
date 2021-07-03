package windows

import (
	"container/list"
	"math"
)

/**
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
*/
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	m := make(map[uint8]int)
	var max int
	for right < len(s) {
		c := s[right]
		right++
		m[c]++
		for m[c] > 1 {
			d := s[left]
			left++
			m[d]--
		}
		max = Max(max, right-left)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/**
76. 最小覆盖子串
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""
*/
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	windows, needed := make(map[uint8]int), make(map[uint8]int)
	for _, item := range t {
		needed[uint8(item)]++
	}
	left, right := 0, 0
	value := ""
	for right < len(s) {
		c := s[right]
		right++
		windows[c]++

		for left < right && check(windows, needed) /*符合条件*/ {
			if right-left < len(value) || len(value) == 0 {
				value = s[left:right]
			}
			windows[s[left]]--
			left++
		}
	}
	return value
}

func check(windows map[uint8]int, needed map[uint8]int) bool {
	for key, value := range needed {
		if v, ok := windows[key]; !ok || v < value {
			return false
		}
	}
	return true
}

/**
239. 滑动窗口最大值
给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。
[1,3,-1,-3,5,3,6,7]

*/
func maxSlidingWindow(nums []int, k int) []int {
	queue := list.New()
	ret := make([]int, 0, len(nums)-k)
	for i := 0; i < len(nums); i++ {
		// 如果队列头的元素移动到了窗口外，需要把头结点剔除
		if front := queue.Front(); front != nil && i-front.Value.(int) >= k {
			queue.Remove(front)
		}
		for back := queue.Back(); back != nil && nums[i] > nums[back.Value.(int)]; back = queue.Back() {
			queue.Remove(back)
		}
		queue.PushBack(i)
		if i >= k-1 {
			ret = append(ret, nums[queue.Front().Value.(int)])
		}
	}
	return ret
}

/**
424. 替换后的最长重复字符
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。
*/
func characterReplacement(s string, k int) int {
	windows := [26]int{}
	left, right := 0, 0
	ans, max := 0, 0
	var maxChar uint8
	for right < len(s) {
		c := s[right]
		windows[c-'A']++
		if windows[c] > max {
			max = windows[c-'A']
			maxChar = c
		}
		right++
		// max = 1 k = 0
		if left < right && max+k < right-left {
			d := s[left]
			left++
			windows[d-'A']--
			if d == maxChar {
				max--
			}
		}
		if d := right - left; d > ans {
			ans = d
		}
	}
	return ans
}

/**
567. 字符串的排列
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
*/
func checkInclusion(s1 string, s2 string) bool {
	windows, needed := [26]int{}, [26]int{}
	left, right := 0, 0
	for i := 0; i < len(s1); i++ {
		needed[s1[i]-'a']++
	}

	for right < len(s2) {
		c := s2[right]
		right++
		windows[c-'a']++
		for left < right && right-left >= len(s1) && leftMove(windows, needed) {
			d := s2[left]
			windows[d-'a']--
			left++
		}
		if right-left >= len(s1) {
			return true
		}
	}
	return false
}
func leftMove(windows [26]int, needed [26]int) bool {
	for k, v := range needed {
		if windows[k] != v {
			return true
		}
	}
	return false
}

/**
1004. 最大连续1的个数 III
给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。
返回仅包含 1 的最长（连续）子数组的长度。
*/
func longestOnes(nums []int, k int) int {
	windows := [2]int{}
	left, right, max := 0, 0, 0
	windowsCount := 0 // 窗口内1的个数
	for right < len(nums) {
		c := nums[right]
		windows[c]++
		right++
		if c == 1 {
			windowsCount += 1
		}
		for left < right && windowsCount+k < right-left {
			d := nums[left]
			windows[d]--
			left++
			if d == 1 {
				windowsCount -= 1
			}
		}
		if d := right - left; d > max {
			max = d
		}
	}
	return max
}

/**
978. 最长湍流子数组
当 A 的子数组 A[i], A[i+1], ..., A[j] 满足下列条件时，我们称其为湍流子数组：

若 i <= k < j，当 k 为奇数时， A[k] > A[k+1]，且当 k 为偶数时，A[k] < A[k+1]；
或 若 i <= k < j，当 k 为偶数时，A[k] > A[k+1] ，且当 k 为奇数时， A[k] < A[k+1]。
也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。

返回 A 的最大湍流子数组的长度
*/
func maxTurbulenceSize(arr []int) int {
	//left, right := 0, 0
	//for right < len(arr) {
	//	c := arr[right]
	//	right++
	//
	//	if left < right && c
	//}
	return 1
}

/**
面试题 17.18. 最短超串
假设你有两个数组，一个长一个短，短的元素均不相同。找到长数组中包含短数组所有的元素的最短子数组，其出现顺序无关紧要。

返回最短子数组的左端点和右端点，如有多个满足条件的子数组，返回左端点最小的一个。若不存在，返回空数组。
*/
func shortestSeq(big []int, small []int) []int {
	windows := make(map[int]int)
	needed := make(map[int]struct{})
	for _, item := range small {
		needed[item] = struct{}{}
	}
	ret := make([]int, 0, 2)
	left, right := 0, 0
	resultLeft, resultRight := 0, 0

	min := math.MaxInt32
	for right < len(big) {
		windows[big[right]]++
		delete(needed, big[right])
		right++
		for left < right && len(needed) <= 0 {
			if d := right - left; d < min {
				min = d
				resultLeft, resultRight = left, right-1
			}
			needed[big[left]] = struct{}{}
			left++
		}

	}
	if resultRight-resultLeft >= 0 && resultRight > 0 {
		ret = append(ret, resultLeft, resultRight)
	}
	return ret
}

/**
1658. 将 x 减到 0 的最小操作数
给你一个整数数组 nums 和一个整数 x 。每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。请注意，需要 修改 数组以供接下来的操作使用。
如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1。
思路：等效于在数组中找到连续和为 sum(nums)-x的最大长度
*/
func minOperations(nums []int, x int) int {
	left, right, sum, maxLength := 0, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	target := sum - x
	if target == 0 {
		return len(nums)
	}
	if target < 0 {
		return -1
	}
	tmp := 0
	for right < len(nums) {
		tmp += nums[right]
		right++
		for left < right && tmp >= target {
			if d := right - left; tmp == target && d > maxLength {
				maxLength = right - left
			}
			tmp -= nums[left]
			left++
		}
	}
	if maxLength == 0 {
		return -1
	}
	return len(nums) - maxLength
}

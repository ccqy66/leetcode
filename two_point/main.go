package two_point

import (
	"github.com/ccqy66/leetcode/model"
	"math"
	"sort"
)

/**
11. 盛最多水的容器
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
*/

// area = (right-left)*min(a[right], a[left])

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area := 0
	for left < right {
		t := (right - left) * min(height[left], height[right])
		if t > area {
			area = t
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return area
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func threeSum(arr []int) [][]int {
	if len(arr) < 3 {
		return [][]int{}
	}
	sort.Ints(arr)
	ret := make([][]int, 0)
	for base := 0; base < len(arr); base++ {
		if arr[base] > 0 {
			return ret
		}
		if base > 0 && arr[base-1] == arr[base] {
			continue
		}
		left, right := base+1, len(arr)-1
		for left < right {
			t := arr[base] + arr[left] + arr[right]
			if t == 0 {
				ret = append(ret, []int{arr[base], arr[left], arr[right]})
				for left < right && arr[left] == arr[left+1] {
					left += 1
				}
				for left < right && arr[right] == arr[right-1] {
					right -= 1
				}
				left += 1
				right -= 1
			}
			if t > 0 {
				right--
			} else if t < 0 {
				left++
			}
		}
	}
	return ret
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			t := nums[left] + nums[i] + nums[right]
			if t == target {
				return target
			}
			if t > target {
				right--
			} else {
				left++
			}
			if abs(t, target) < abs(ans, target) {
				ans = t
			}
		}
	}
	return ans
}

func abs(a, b int) int {
	if d := a - b; d > 0 {
		return d
	} else {
		return -d
	}
}

/**
	加入链表有m个节点：
	快指针先走了k步，到达k的位置上。
	快慢指针同时走：所以所以快指针走了：n-k
	慢指针也走了n-k步

如果有5个节点，k=2
  head	1 2 3 4 5
		|
	|
*/
func removeNthFromEnd(head *model.ListNode, n int) *model.ListNode {
	// 引入一个头结点，其实就是为了方便计算位置
	vhead := &model.ListNode{Val: -1, Next: head}
	fast, slow := vhead, vhead
	// 快指针移动n步，到达第n个节点
	for i := 0; i < n && fast.Next != nil; i++ {
		fast = fast.Next
	}
	// 快慢指针同时移动:
	// 当快指针到达尾部时：快指针一共走了l步（链表长度），此时慢指针走了l-k步，实际上也就是倒数第n个节点
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return vhead.Next
}

/**
26. 删除有序数组中的重复项
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
1,2 => 1,2

*/
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	p, q := 0, 1
	for q < len(nums) {
		if nums[p] == nums[q] {
			q++
		}else {
			p++
			nums[p] = nums[q]
			q++
		}
	}
	return p+1
}

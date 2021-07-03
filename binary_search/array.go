package binary_search

// [5,6,7,0,1,2,3,4] target=0
// left=0, right=8 mid=4 [0, 4)
// left=0, right=4 mid=2 [)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[0] > nums[mid] { // [mid, right] 有序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target >= nums[left] && target < nums[mid] { // [0, mid] 有序
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	return left
}

func twoSum(numbers []int, target int) []int {
	result := make([]int, 0, 2)
	for i := 0; i < len(numbers); i++ {
		if r := binarySearch(numbers, target-numbers[i], i+1, len(numbers)); r > 0 {
			result = append(result, i+1, r+1)
			return result
		}
	}
	return result
}

func binarySearch(nums []int, target, start, end int) int {
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid
		}
		if nums[mid] < target {
			start = mid + 1
		}
	}
	return -1
}

func isBadVersion(version int) bool {
	if version >= 4 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	lo, hi := 1, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func intersection(nums1 []int, nums2 []int) []int {
	return nums1
}

func searchSingleRange(nums []int, target int, leftRange bool) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if leftRange {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if nums[mid] > target {
			right = mid
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	if leftRange {
		return left
	}
	return left - 1
}

func searchRange(nums []int, target int) []int {
	if len(nums) <= 0 || nums[len(nums)-1] < target {
		return []int{-1, -1}
	}
	leftRange := searchSingleRange(nums, target, true)
	if nums[leftRange] != target {
		return []int{-1, -1}
	}
	return []int{leftRange, searchSingleRange(nums, target, false)}
}

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i <= len(matrix)-1 && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
			continue
		}
		if matrix[i][j] < target {
			i++
			continue
		}
	}
	return false
}

// 81. 搜索旋转排序数组 II
func searchV2(nums []int, target int) bool {
	if len(nums) == 1 {
		return nums[0] == target
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[left] && nums[mid] == nums[left] {
			left++
			right--
		}

		if nums[0] > nums[mid] { // [mid, right] 有序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target >= nums[left] && target < nums[mid] { // [0, mid] 有序
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	if nums[left+1] == target {
		return true
	}
	return false
}

//       0 1 2 3 4 5 6
// []int{4,5,6,7,0,1,2}
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left +(right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		}else {
			left = mid+1

		}
	}
	return nums[left]
}

func findPeakElement(nums []int) int {
	//left, right := 0, len(nums)
	//for left < right {
	//	mid := left + (right-left)/2
	//
	//}
	return 1
}
package recursion

import "strconv"

func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) <= 0 {
		return false
	}
	sum := 0
	for _, item := range nums {
		sum += item
	}
	if sum%k != 0 {
		return false
	}

	return false
}

var r []int

func sameDifference(n, k int) []int {
	r = make([]int, 0)
	for i := 1; i <= 9; i++ {
		difference(i, n, k)
	}
	return r
}

func difference(m, n, k int) {
	if len(strconv.Itoa(m)) == n {
		r = append(r, m)
		return
	}
	t := m % 10
	if d := t - k; d >= 0 {
		difference(m*10+d, n, k)
	}
	if d := k + t; d < 10 {
		difference(m*10+d, n, k)
	}
}
func back2Zero(arr []int, start int) bool{
	exist := false
	for _, item := range arr{
		if item == 0 {
			exist = true
			break
		}
	}
	if !exist {
		return false
	}
	return back(arr, start, start, true)
}
func back(arr []int, index, start int, first bool)bool {
	if index == start && !first{ // 回到原点，说明不可能回到0的位置
		return false
	}
	if arr[index] == 0 {
		return true
	}
	if d := index - arr[index]; d > 0 {
		return back(arr, d, start, false)
	}
	if d := index + arr[index]; d < len(arr) {
		return back(arr, d, start, false)
	}
	return false
}
func divideBoard(shorter, longer, k int) []int{
	if k == 0 {
		return make([]int, 0)
	}
	if k == 1 {
		return []int{shorter, longer}
	}
	l := divideBoard(shorter, longer, k-1)
	result := make([]int, 0)
	m := make(map[int]struct{})
	for _, item := range l{
		if _, ok := m[item+shorter]; !ok {
			m[item+shorter] = struct{}{}
			result = append(result, item+shorter)
		}
		if _, ok := m[item+longer]; !ok {
			m[item+longer] = struct{}{}
			result = append(result, item+longer)
		}
	}
	return result
} //0,1,2,3,4,5
// [1,2,2,3,4,5]
// start=0 end=6 mid=3 [0,3)
// start=0 end=2 mid=1 [0,1)
// start=0 end=1 mid=0 [1,1))
func searchLeft(array []int, target int) int{
	start, end := 0, len(array)
	for start <  end {
		mid := start + (end - start) / 2
		println(start, end, mid)
		if array[mid] == target {
			end = mid
		}
		if array[mid] > target {
			end = mid
		}
		if array[mid] < target {
			start = mid+1
		}
	}
	return end
}
//  0,1,2,3,4,5
// [1,2,2,4,5,6] 2
// left=0 right=6 mid=3 [0,3)
// left=0 right=3 mid=1 [2,3)
// left=1 right=3 mid=2 [3,3)
// left=2 right=3 mid=2 [2,3)
/**
int right_bound(int[] nums, int target) {
    if (nums.length == 0) return -1;
    int left = 0, right = nums.length;

    while (left < right) {
        int mid = (left + right) / 2;
        if (nums[mid] == target) {
            left = mid + 1; // 注意
        } else if (nums[mid] < target) {
            left = mid + 1;
        } else if (nums[mid] > target) {
            right = mid;
        }
    }
    return left - 1; // 注意
}
 */
var director = [][]int{{-1,1},{1,-1}, {-1,-1}, {1,1}}
func searchRight(array []int, target int)int {
	left, right := 0, len(array)
	for left < right{
		mid := left + (right-left)/2
		if array[mid] == target {
			left = mid+1
		}
		if array[mid] < target {
			left = mid+1
		}
		if array[mid] > target {
			right = mid
		}
	}
	return right-1
}

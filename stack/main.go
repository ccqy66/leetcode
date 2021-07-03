package stack

func removeDuplicateLetters(str string) string {
	stack := make([]uint8, 0)
	stackm := make(map[uint8]int16)
	m := make(map[uint8]int)
	for i := 0; i < len(str); i++ {
		m[str[i]]++
	}
	for i := 0; i < len(str); i++ {
		if _, ok := stackm[str[i]]; !ok {
			for len(stack) > 0 && str[i] < stack[len(stack)-1] && stackm[stack[len(stack)-1]] > 0 {
				delete(stackm, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, str[i])
			stackm[str[i]] = 1
		}
		m[str[i]] -= 1
	}
	return string(stack)
}

// 4 3 2 1
// 1 2 3 5 4 2 =>  [1,2,4,2,3,5]
// 2 3 4 5 2 1
// 4 3 2 => 2 3 4
// 1 3 2 => 2 1 3
// =>
func nextPermutation(nums []int) {
	stack := make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if len(stack) > 0 && nums[i] < nums[stack[len(stack)-1]] {
			index := i
			for len(stack) > 0 && nums[i] < nums[stack[len(stack)-1]]{
				index = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			nums[i], nums[index] = nums[index], nums[i]
			return
		} else {
			stack = append(stack, i)
		}
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

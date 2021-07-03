package top
/**
1. 两数之和
 */
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, item := range nums{
		m[item] = idx
	}
	ret := make([]int, 0, 2)
	for index, item := range nums{
		if v, ok := m[target-item]; ok && v != index {
			ret = append(ret, index, v)
			return ret
		}
	}
	return ret
}

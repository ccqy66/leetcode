package greedy
/**
55. 跳跃游戏
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。
 */
func canJump(nums []int) bool {
	reach := 0
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}
		if s := i+nums[i]; s >reach {
			reach = s
		}
	}
	return true
}
package dp

import (
	"math"
)

func maxSubArray(nums []int) int {
	// dp[i] = dp[i-1]+nums[i]; if dp[i-1]+nums[i] >= nums[i]
	// dp[i] = nums[i]; if dp[i-1]+nums[i] < dp[i-1]
	dp := make([]int, len(nums))
	max := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] >= nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

// [1,3,1]
// [1,5,1]
// [4,2,1]

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	dp[0] = make([]int, len(grid[0]))
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

// dp[i] = dp[i-1] // s[i-1][i]不能组成，且s[i] == 0
// dp[i] = dp[i-1]+1 // s[i-1][i] 不能组成一个有效的字母、但是s[i] != 0
// dp[i] = dp[i-1]+2 // s[i-1][i] 可以组成一个有效的字母、且s[i] != 0

// 227=>
// dp[0] = 1/ dp[1] = 2
// 22 7
// s[i-1][i]能组成
//  - s[i] == 0 =>dp[i] = dp[i-2]
//  - s[i] != 0 =>dp[i] = dp[i-2] + dp[i-1]
// s[i-1][i]不能组成
//  - s[i] == 0 =>dp[i] = dp[i-2]
//  - s[i] != 0 =>dp[i] = dp[i-1]

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		// s[i-1][i]能组成一个有效的字母
		if s[i] != '0' {
			dp[i] += dp[i-1]
		}
		if v := (s[i-1]-'0')*10 + s[i] - '0'; s[i-1] != '0' && v <= 26 {
			if i > 1 {
				dp[i] += dp[i-2]
			} else {
				dp[i] += 1
			}
		}
	}
	return dp[len(s)-1]
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1) == 0 && len(s2) == 0 && len(s3) == 0 {
		return true
	}
	dp := make([][]bool, len(s1)+1) // dp[i][j]:代表s1[:i]s2[:j]是否可以组成 s3[:i+j]
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	for i := 1; i <= len(s1); i++ {
		if s1[:i] != s3[:i] {
			break
		}
		dp[i][0] = true
	}
	for i := 1; i <= len(s2); i++ {
		if s2[:i] != s3[:i] {
			break
		}
		dp[0][i] = true
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if dp[i-1][j] && s1[i-1] == s3[i+j-1] {
				dp[i][j] = true
			}
			if dp[i][j-1] && s2[j-1] == s3[i+j-1] {
				dp[i][j] = true
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

func maxProduct(nums []int) int {
	//dp1 := make([]int, len(nums)) // dp[i]代表以nums[i]结尾的最大乘积
	//dp2 := make([]int, len(nums)) // dp[i]代表以nums[i]结尾的绝对值最大乘积
	//ans := nums[0]
	//dp1[0], dp2[0] = nums[0], nums[0]
	//for i := 1; i < len(nums); i++ {
	//	abs := dp1[i-2] * nums[i]
	//	max := dp1[i-1]*nums[i]
	//	if max > ans {
	//		ans = max
	//	}
	//	if abs > ans {
	//		ans = abs
	//	}
	//
	//
	//}
	return 0
}

//   2
//  3 4
// 6 5 7
//4 1 8 3
// [2, 0, 0, 0]
// [5, 6, 0, 0]
// [11,10,11,0]
// [15,11,18,14]
func minimumTotal(triangle [][]int) int {
	if len(triangle) <= 0 {
		return 0
	}
	dp := make([]int, len(triangle)) // dp[i]代表到达每一行第i个元素的最短路径和
	// dp[i] = dp[i-1]
	dp[0] = triangle[0][0]
	minValue := dp[0]
	for i := 1; i < len(triangle); i++ {
		tmpValue := math.MaxInt32
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if j == len(triangle[i])-1 {
				dp[j] = dp[j-1] + triangle[i][j]
			} else if j == 0 {
				dp[j] = dp[j] + triangle[i][j]
			} else {
				dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
			}
			if dp[j] < tmpValue {
				tmpValue = dp[j]
			}
		}
		minValue = tmpValue
	}
	return minValue
}

var director = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
var dp [][]int

func MaxLengthPath(nums [][]int) int {
	maxValue := 0
	dp = make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums[0]))
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			maxValue = max(dfs(nums, i, j), maxValue)
		}
	}

	return maxValue
}

func dfs(nums [][]int, i, j int) int{
	if dp[i][j] > 0 {
		return dp[i][j]
	}
	t := 1
	for index := 0; index < len(director); index++ {
		newRow, newCow := i+director[index][0], j+director[index][1]
		if newRow >= 0 && newRow < len(nums) && newCow >= 0 && newCow < len(nums[0]) &&
			nums[newRow][newCow] < nums[i][j] {
			t = max(t, dfs(nums, newRow, newCow)+1)
		}
	}
	dp[i][j] = t
	return dp[i][j]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
// 只有三种钱币[2,5,7]，问amount最少有几种找法
// findChange(): amount钱的最少找法
func findChange(amount int) int{
	if amount == 0 {
		return 0
	}
	ans := math.MaxInt32
	if amount >= 2 {
		ans = min(ans, findChange(amount-2)+1)
	}
	if amount >= 5 {
		ans = min(ans, findChange(amount-5)+1)
	}
	if amount >= 7 {
		ans = min(ans, findChange(amount-7)+1)
	}
	return ans
}

func canPartition(nums []int) bool {
	
}
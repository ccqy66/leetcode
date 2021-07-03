package dfs

import (
	"fmt"
	"math"
	"sort"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return targetSum == 0
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	var left bool
	var right bool
	if root.Left != nil {
		left = hasPathSum(root.Left, targetSum-root.Val)
	}
	if root.Right != nil {
		right = hasPathSum(root.Right, targetSum-root.Val)
	}
	return left || right
}

var r [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
	r = make([][]int, 0)
	result := make([]int, 0)
	dfs(root, targetSum, result)
	return r
}

func dfs(root *TreeNode, targetSum int, result []int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		result = append(result, root.Val)
		r = append(r, result)
	}
	dst := make([]int, len(result), len(result)+1)
	copy(dst, result)
	dst = append(dst, root.Val)

	if root.Left != nil {
		dfs(root.Left, targetSum-root.Val, dst)
	}
	if root.Right != nil {
		dfs(root.Right, targetSum-root.Val, dst)
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)
	if p.Val != q.Val {
		return false
	}
	return left && right
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if i < l-1 {
				queue[i].Next = queue[i+1]
			} else {
				queue[i].Next = nil
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}
	return root
}

var rpartition [][]string

func partition(s string) [][]string {
	rpartition = make([][]string, 0)
	result := make([]string, 0)
	dfspartition(s, 0, len(s), 0, result)
	return rpartition
}

func dfspartition(s string, start, end int, num int, result []string) {
	if start >= end {
		if num == len(s) {
			rpartition = append(rpartition, result)
		}
		return
	}
	for i := start; i < end; i++ {
		if check(s, start, i) {
			dst := make([]string, len(result), len(result)+1)
			copy(dst, result)
			dst = append(dst, s[start:i+1])
			dfspartition(s, i+1, end, num+i-start+1, dst)
		}
	}
}
func check(s string, start, end int) bool {
	for start <= end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

var rsumNumbers int

func sumNumbers(root *TreeNode) int {
	rsumNumbers = 0
	dfsSumNumbers(root, 0)
	return rsumNumbers
}

func dfsSumNumbers(root *TreeNode, value int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		rsumNumbers += value*10 + root.Val
		return
	}
	if root.Left != nil {
		dfsSumNumbers(root.Left, value*10+root.Val)
	}
	if root.Right != nil {
		dfsSumNumbers(root.Right, value*10+root.Val)
	}
}

var director = [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

func numIslands(grid [][]byte) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfsnumIslands(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfsnumIslands(grid [][]byte, i, j int) {
	grid[i][j] = '0'
	for index := 0; index < len(director); index++ {
		row, cow := i+director[index][0], j+director[index][1]
		if row >= 0 && row < len(grid) && cow >= 0 && cow < len(grid[0]) && grid[row][cow] == '1' {
			dfsnumIslands(grid, row, cow)
		}
	}
}

var result []string

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return result
	}
	result = make([]string, 0)
	dfsbinaryTreePaths(root, "")
	return result
}

func dfsbinaryTreePaths(root *TreeNode, item string) {
	if item == "" {
		item = fmt.Sprintf("%d", root.Val)
	} else {
		item = fmt.Sprintf("%s->%d", item, root.Val)
	}
	if root.Left == nil && root.Right == nil {
		result = append(result, item)
		return
	}
	if root.Left != nil {
		dfsbinaryTreePaths(root.Left, item)
	}
	if root.Right != nil {
		dfsbinaryTreePaths(root.Right, item)
	}
}

var rrob int

func rob(root *TreeNode) int {
	rrob = 0
	dfsrob(root, root.Val, true)
	dfsrob(root, 0, false)
	return rrob
}

/**
value 记录当前的值
lastSelected: 父节点是否被选择了
*/
func dfsrob(root *TreeNode, value int, lastSelected bool) {
	if value > rrob {
		rrob = value
	}
	if root == nil {
		return
	}
	if !lastSelected {
		if root.Left != nil {
			dfsrob(root.Left, value+root.Val, !lastSelected)
		}
		if root.Right != nil {
			dfsrob(root.Right, value+root.Val, !lastSelected)
		}
	} else {
		if root.Left != nil {
			dfsrob(root.Left, value, !lastSelected)
		}
		if root.Right != nil {
			dfsrob(root.Right, value, !lastSelected)
		}
	}

}

var direction = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}

func pondSizes(land [][]int) []int {
	ret := make([]int, 0)
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[0]); j++ {
			if land[i][j] == 0 {
				r := dfspondSizes(land, i, j)
				ret = append(ret, r)
			}
		}
	}
	sort.Ints(ret)
	return ret
}

func dfspondSizes(land [][]int, i, j int) int {
	land[i][j] = -1
	count := 1
	for index := 0; index < len(direction); index++ {
		row, cow := i+direction[index][0], j+direction[index][1]
		if row >= 0 && row < len(land) && cow >= 0 && cow < len(land[0]) && land[row][cow] == 0 {
			count += dfspondSizes(land, row, cow)
		}
	}
	return count
}

var rpathSum int

func pathSum2(root *TreeNode, sum int) int {
	rpathSum = 0
	sumList := make([]int, 0)
	dfspathSum(root, sum, sumList)
	return rpathSum
}

func dfspathSum(root *TreeNode, sum int, sumList []int) {
	if root == nil {
		return
	}
	dst := make([]int, len(sumList), len(sumList)+1)
	copy(dst, sumList)
	for index, _ := range sumList {
		dst[index] += root.Val
		if dst[index] == sum {
			rpathSum++
		}
	}
	if sum == root.Val {
		rpathSum++
	}
	dst = append(dst, root.Val)
	if root.Left != nil {
		dfspathSum(root.Left, sum, dst)
	}
	if root.Right != nil {
		dfspathSum(root.Right, sum, dst)
	}
}

func exist(board [][]byte, word string) bool {

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if dfsexist(board, word, 0, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func dfsexist(board [][]byte, word string, index, i, j int) bool {
	if board[i][j] != word[index] {
		return false
	}
	if index == len(word)-1 {
		return true
	}
	board[i][j] = '-'
	for in := 0; in < len(director); in++ {
		row, cow := i+director[in][0], j+director[in][1]
		if row >= 0 && row < len(board) && cow >= 0 && cow < len(board[0]) {
			if dfsexist(board, word, index+1, row, cow) {
				return true
			}
		}
	}
	board[i][j] = word[index]
	return false
}

var rpathSum3 [][]int

func pathSum3(root *TreeNode, target int) [][]int {
	rpathSum3 = make([][]int, 0)
	result := make([]int, 0)
	dfspathSum3(root, target, result)
	return rpathSum3
}

func dfspathSum3(root *TreeNode, target int, result []int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && target == root.Val {
		result = append(result, root.Val)
		rpathSum3 = append(rpathSum3, result)
		return
	}
	dst := make([]int, len(result), len(result)+1)
	copy(dst, result)
	dst = append(dst, root.Val)
	if root.Left != nil {
		dfspathSum3(root.Left, target-root.Val, dst)
	}
	if root.Right != nil {
		dfspathSum3(root.Right, target-root.Val, dst)
	}

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

//func isBalanced(root *TreeNode) bool {
//
//}

//func depth(root *TreeNode) int{
//	if root == nil {
//		return 0
//	}
//
//}
var rfindLadders []string
var visited map[string]bool

func findLadders(beginWord string, endWord string, wordList []string) []string {
	rfindLadders = make([]string, 0)
	visited = make(map[string]bool)
	wordList = append([]string{beginWord}, wordList...)
	wordPath := make(map[string][]string)
	for i := 0; i < len(wordList); i++ {
		if len(wordList[i]) != len(beginWord) {
			continue
		}
		for j := 0; j < len(wordList); j++ {
			if len(wordList[i]) != len(wordList[j]) {
				continue
			}
			count := 0
			for c := 0; c < len(wordList[i]); c++ {
				if wordList[i][c] != wordList[j][c] {
					count++
				}
				if count > 1 {
					continue
				}
			}
			if count == 1 {
				wordPath[wordList[i]] = append(wordPath[wordList[i]], wordList[j])
			}
		}
	}
	result := []string{beginWord}
	dfsfindLadders(beginWord, endWord, wordPath, result)
	return rfindLadders
}

func dfsfindLadders(beginWord string, endWord string, path map[string][]string, result []string) {
	if beginWord == endWord {
		rfindLadders = result
		return
	}
	visited[beginWord] = true
	p := path[beginWord]
	for _, item := range p {
		if visited[item] {
			continue
		}
		if len(rfindLadders) > 0 {
			return
		}
		//dst := make([]string, len(result), len(result)+1)
		//copy(dst, result)

		result = append(result, item)
		dfsfindLadders(item, endWord, path, result)
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	dfsfloodFill(image, sr, sc, newColor)
	return image
}

func dfsfloodFill(image [][]int, sr int, sc int, newColor int) {
	oldColor := image[sr][sc]
	image[sr][sc] = newColor
	if newColor == oldColor {
		return
	}
	for i := 0; i < len(director); i++ {
		new_sr, new_sc := sr+director[i][0], sc+director[i][1]
		if new_sr >= 0 && new_sr < len(image) && new_sc >= 0 && new_sc < len(image[0]) &&
			image[new_sr][new_sc] == oldColor {
			dfsfloodFill(image, new_sr, new_sc, newColor)
		}
	}

}
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var (
		pre *TreeNode = nil
		f   func(root *TreeNode, p *TreeNode) *TreeNode
	)
	f = func(root *TreeNode, p *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if r := f(root.Left, p); r != nil {
			return r
		}
		if p == pre {
			return root
		}
		pre = root
		if r := f(root.Right, p); r != nil {
			return r
		}
		return nil
	}
	return f(root, p)
}

var minValue int

func minimumEffortPath(heights [][]int) int {
	minValue = math.MaxInt32
	dfsminimumEffortPath(heights, 0, 0, 0)
	return minValue
}

func dfsminimumEffortPath(heights [][]int, i, j int, result int) {
	if i == len(heights)-1 && j == len(heights[0])-1 {
		if result < minValue {
			minValue = result
		}
		return
	}
	value := heights[i][j]
	heights[i][j] = -1
	for c := 0; c < len(director); c++ {
		newI, newJ := i+director[c][0], j+director[c][1]
		if newI >= 0 && newI < len(heights) && newJ >= 0 && newJ < len(heights[0]) && heights[newI][newJ] != -1 {
			var tmpResult  = result
			if d := abs(heights[newI][newJ] - value); d > result {
				tmpResult = d
			}
			dfsminimumEffortPath(heights, newI, newJ, tmpResult)
		}
	}

	heights[i][j] = value
}

func abs(a int)  int{
	if a > 0 {
		return a
	}
	return -a
}
var lres int
func longestIncreasingPath(matrix [][]int) int {
	lres = 0
	for i := 0 ; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++{
			dfsLongestIncreasingPath(matrix, i, j, []int{matrix[i][j]})
		}
	}
	return lres
}

func dfsLongestIncreasingPath(matrix [][]int, i, j int, ans []int) {
	if len(ans) > lres {
		lres = len(ans)
	}
	tmp := matrix[i][j]
	matrix[i][j] = -1
	for _, item := range director {
		newRow, newCow := i + item[0], j + item[1]
		if newRow >= 0 && newRow < len(matrix) && newCow >= 0 && newCow < len(matrix[0]) &&
			matrix[newRow][newCow] > tmp &&matrix[newRow][newCow] != -1{
			dst := make([]int, len(ans))
			copy(dst, ans)
			dst = append(dst, matrix[newRow][newCow])
			dfsLongestIncreasingPath(matrix, newRow, newCow, dst)
		}
	}
	matrix[i][j] = tmp
}

func sumOfLeft(root *TreeNode) int{
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val+sumOfLeft(root.Right)
	}
	return sumOfLeft(root.Left)+sumOfLeft(root.Right)
}
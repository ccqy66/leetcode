package recursion

import (
	"container/list"
	"github.com/ccqy66/leetcode/utils"
	"math"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func NewBinaryTreeNode(value int) *TreeNode {
	return &TreeNode{Val: value}
}

func (node *TreeNode) add(value int) {
	if value < node.Val {
		if node.Left == nil {
			node.Left = NewBinaryTreeNode(value)
		} else {
			node.Left.add(value)
		}
	}
	if value > node.Val {
		if node.Right == nil {
			node.Right = NewBinaryTreeNode(value)
		} else {
			node.Right.add(value)
		}
	}
}

func printBinarySearchTree(root *TreeNode) {
	if root == nil {
		return
	}
	printBinarySearchTree(root.Left)
	println(root.Val)
	printBinarySearchTree(root.Right)
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := height(root.Left)
	right := height(root.Right)
	if utils.Abs(int(left-right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return utils.Max(height(root.Left), height(root.Right)) + 1
}

var result int

func maxPathLength(root *TreeNode) int {
	result = 0
	compute(root)
	return result
}

// 已当前节点为根的最大路径
func compute(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var l, r int
	left := compute(root.Left)
	right := compute(root.Right)
	if root.Left != nil && root.Left.Val == root.Val {
		l = left + 1
	}
	if root.Right != nil && root.Right.Val == root.Val {
		r = right + 1
	}
	result = Max(result, l+r)
	return Max(l, r)
}

func Max(a ...int) int {
	if len(a) <= 0 {
		return 0
	}
	max := a[0]
	for _, item := range a {
		if item > max {
			max = item
		}
	}
	return max
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil || (root.Val < low && root.Val > high) {
		return nil
	}
	left := trimBST(root.Left, low, high)
	right := trimBST(root.Right, low, high)

	if root.Val < low {
		return right
	}
	if root.Val > high {
		return left
	}
	root.Left = left
	root.Right = right
	return root
}

var p *TreeNode

func isValidBST(root *TreeNode) bool {
	p = nil
	return help(root)
}

func help(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !help(root.Left) {
		return false
	}
	if p != nil && root.Val <= p.Val {
		return false
	}
	pre = root
	return help(root.Right)
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

var res int

func maxPathSum(root *TreeNode) int {
	res = -math.MaxInt32
	pathSum(root)
	return res
}
func pathSum(root *TreeNode) int {
	if root == nil {
		return -math.MaxInt32
	}
	left := pathSum(root.Left)
	right := pathSum(root.Right)
	res = Max(left+right+root.Val, root.Val+left, root.Val+right, root.Val, right, left, res)
	return Max(root.Val, root.Val+left, root.Val+right)
}
func convertBST(root *TreeNode) *TreeNode {
	accumulationTree(root)
	return root
}

	var pre *TreeNode

func accumulationTree(root *TreeNode) {
	if root == nil {
		return
	}
	accumulationTree(root.Right)
	if pre != nil {
		root.Val += pre.Val
	}
	pre = root
	accumulationTree(root.Left)
}

//var pre *TreeNode
var ret *TreeNode

func transRightTree(root *TreeNode) *TreeNode {
	ret = nil
	transRight(root)
	return ret
}

func transRight(root *TreeNode) {
	if root == nil {
		return
	}
	transRight(root.Left)
	if pre != nil {
		pre.Right = root
	}
	if ret == nil {
		ret = root
	}
	pre = root
	root.Left = nil
	transRight(root.Right)
}

func rebuildTree(pre []int, in []int) *TreeNode {
	return rebuildNode(pre, in, 0, 0, len(in)-1)
}

func rebuildNode(pre []int, in []int, preIndex, inStartIndex, inEndIndex int) *TreeNode {
	if preIndex == len(pre) || inEndIndex < inStartIndex {
		return nil
	}
	root := &TreeNode{
		Val: pre[preIndex],
	}
	for i := inStartIndex; i <= inEndIndex; i++ {
		if in[i] == root.Val {
			root.Left = rebuildNode(pre, in, preIndex+1, inStartIndex, i-1)
			root.Right = rebuildNode(pre, in, preIndex+i-inStartIndex+1, i+1, inEndIndex)
		}
	}
	return root
}

func printCousins(root *TreeNode, p *TreeNode) []int {
	queue := list.New()
	queue.PushBack(root)
	ret := make([]int, 0)
	stop := false
	for queue.Len() > 0 && !stop {
		len := queue.Len()
		for i := 0; i < len; i++ {
			front := queue.Front()
			queue.Remove(front)
			node := front.Value.(*TreeNode)
			if (node.Left != nil && node.Left == p) || (node.Right != nil && node.Right == p) {
				stop = true
				continue
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	for item := queue.Front(); item != nil; item = item.Next() {
		ret = append(ret, item.Value.(*TreeNode).Val)
	}
	return ret
}

var maxValue int

func maxRootToLeaf(root *TreeNode) int {
	maxValue = math.MinInt32
	dfs(root, 0)
	return maxValue
}
func dfs(root *TreeNode, value int) {
	if root.Left == nil && root.Right == nil {
		if v := value + root.Val; v > maxValue {
			maxValue = v
		}
		return
	}
	if root.Left != nil{
		dfs(root.Left, value*10+root.Val)
	}
	if root.Right != nil {
		dfs(root.Right, value*10+root.Val)
	}
}


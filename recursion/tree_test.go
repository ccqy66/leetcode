package recursion

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	root := NewBinaryTreeNode(3)
	root.add(1)
	root.add(2)
	root.add(4)
	printBinarySearchTree(root)
}
/**
	5
   / \
  5   5
 / \   \
5   1   5
 */
func TestMaxPathLength(t *testing.T) {
	root := &TreeNode{
		Left:  &TreeNode{
			Left:  &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   5,
			},
			Right: &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   1,
			},
			Val: 5,
		},
		Right: &TreeNode{
			Left:  nil,
			Right: &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   5,
			},
			Val: 5,
		},
		Val: 5,
	}
	fmt.Println(maxPathLength(root))
}
/**
	3
   / \
  2   4
 /
1
 */
func TestTrimBST(t *testing.T) {
	root := NewBinaryTreeNode(3)
	root.add(2)
	root.add(4)
	root.add(1)

	p := trimBST(root, 1, 1)
	printBinarySearchTree(p)
}

func TestIsValidBST(t *testing.T) {
	//root := NewBinaryTreeNode(3)
	//root.add(2)
	//root.add(4)
	//root.add(1)
	root := &TreeNode{
		Left:  &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   1,
		},
		Right: &TreeNode{
			Left:  &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   3,
			},
			Right: &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   6,
			},
			Val:   4,
		},
		Val:   5,
	}
	fmt.Println(isValidBST(root))
}

func TestMaxPathSum(t *testing.T) {
	root := &TreeNode{
		Left: &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   9,
		},
		Right: &TreeNode{
			Left:  &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val:   7,
			},
			Val:   20,
		},
		Val:   -10,
	}
	root = &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   -3,
	}
	fmt.Println(maxPathSum(root))
}

func TestAccumulationTree(t *testing.T) {
	root := &TreeNode{
		Left: &TreeNode{
			Left:  &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   3,
			},
			Right: &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   6,
			},
			Val:   5,
		},
		Right: &TreeNode{
			Left:  &TreeNode{
				Val: 11,
			},
			Right: &TreeNode{
				Val:   13,
			},
			Val:   12,
		},
		Val:   10,
	}
	accumulationTree(root)
	printBinarySearchTree(root)
}

func TestTransRightTree(t *testing.T) {
	root := &TreeNode{
		Left: &TreeNode{
			Left:  &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   3,
			},
			Right: &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   6,
			},
			Val:   5,
		},
		Right: &TreeNode{
			Left:  &TreeNode{
				Val: 11,
			},
			Right: &TreeNode{
				Val:   13,
			},
			Val:   12,
		},
		Val:   10,
	}
	root = transRightTree(root)
	printBinarySearchTree(root)
}

func TestRebuildTree(t *testing.T) {
	root := rebuildTree([]int{1,2,4,5,3}, []int{4,2,5,1,3})
	printBinarySearchTree(root)
}

/**
		10
		/ \
		5 12
	/ \   /\
	3 6  11 13
 */

func TestPrintCousins(t *testing.T) {
	dst := &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   6,
	}
	root := &TreeNode{
		Left: &TreeNode{
			Left:  &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   3,
			},
			Right: dst,
			Val:   5,
		},
		Right: &TreeNode{
			Left:  &TreeNode{
				Val: 11,
			},
			Right: &TreeNode{
				Val:   13,
			},
			Val:   12,
		},
		Val:   10,
	}
	d := printCousins(root,dst)
	fmt.Println(d)
}
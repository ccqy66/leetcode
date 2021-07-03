package dfs

import (
	"fmt"
	"testing"
)

func TestPathSum(t *testing.T) {
	root := &TreeNode{
		Left:  &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   2,
		},
		Right: &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   3,
		},
		Val:   1,
	}
	fmt.Println(pathSum(root, 3))
}


func TestConnect(t *testing.T) {
	root := &Node{
		Val:   1,
		Left:  &Node{
			Val: 2,
			Left: &Node{
				Val:   4,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   5,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next:  nil,
		},
		Right: &Node{
			Left: &Node{
				Val:   6,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   7,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Val:   3,
		},
		Next:  nil,
	}
	fmt.Println(connect(root))
}

func TestPartition(t *testing.T) {
	fmt.Println(partition("aab"))
}

func TestNumIsland(t *testing.T) {
	grid := [][]byte{
		{'1','1','1','1','0'},
		{'1','1','0','1','0'},
		{'1','1','0','0','0'},
		{'0','0','0','0','0'},
	}
	fmt.Println(numIslands(grid))
}

func TestRob(t *testing.T) {
	root := &TreeNode{
		Left:  &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   2,
		},
		Right: &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   3,
		},
		Val:   1,
	}
	fmt.Println(rob(root))
}

func TestPondSizes(t *testing.T)  {
	land := [][]int{
		{0,2,1,0},
		{0,1,0,1},
		{1,1,0,1},
		{0,1,0,1},
	}
	fmt.Println(pondSizes(land))
}

func TestPathSum2(t *testing.T) {
	root := &TreeNode{
		Left:  &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   2,
		},
		Right: &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   3,
		},
		Val:   1,
	}
	fmt.Println(pathSum2(root, 3))
}

func TestExist(t *testing.T)  {
	board := [][]byte{
		{'C','A','A'},
		{'A','A','A'},
		{'B','C','D'},
	}
	fmt.Println(exist(board,"AAB"))
}

func TestFindLadders(t *testing.T)  {
	r := findLadders("qa", "sq", []string{"si","go","se","cm","so","ph","mt","db","mb","sb","kr","ln","tm","le","av","sm","ar","ci","ca","br","ti","ba","to","ra","fa","yo","ow","sn","ya","cr","po","fe","ho","ma","re","or","rn","au","ur","rh","sr","tc","lt","lo","as","fr","nb","yb","if","pb","ge","th","pm","rb","sh","co","ga","li","ha","hz","no","bi","di","hi","qa","pi","os","uh","wm","an","me","mo","na","la","st","er","sc","ne","mn","mi","am","ex","pt","io","be","fm","ta","tb","ni","mr","pa","he","lr","sq","ye"})
	fmt.Println(r)
}

func Test(t *testing.T) {
	r := make([]string, 0)
	s(r)
	fmt.Println(r)
}
func s(res []string) {
	if len(res) == 2 {
		return
	}
	res = append(res, "aaaa")
	fmt.Println(&res)
	s(res)
}

func TestFloodFill(t *testing.T){
	image := [][]int{
		{1,1,1},
		{1,1,0},
		{1,0,1},
	}
	fmt.Println(floodFill(image, 1,1,2))
}

func TestInorderSuccessor(t *testing.T) {
	p := &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   2,
	}
	root := &TreeNode{
		Left:  p,
		Right: &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   3,
		},
		Val:   1,
	}
	fmt.Println(inorderSuccessor(root, p))
}

func TestMinimumEffortPath(t *testing.T)  {

	height := [][]int{
		{8,3,2,5,2,10,7,1,8,9},
		{1,4,9,1,10,2,4,10,3,5},
		{4,10,10,3,6,1,3,9,8,8},
		{4,4,6,10,10,10,2,10,8,8},
		{9,10,2,4,1,2,2,6,5,7},
		//{2,9,2,6,1,4,7,6,10,9},
		//{8,8,2,10,8,2,3,9,5,3},
		//{2,10,9,3,5,1,7,4,5,6},
		//{2,3,9,2,5,10,2,7,1,8},
		//{9,10,4,10,7,4,9,3,1,6},
	}
	fmt.Println(minimumEffortPath(height))
}

func TestLongestIncreasingPath(t *testing.T) {
	fmt.Println(longestIncreasingPath([][]int{{9,9,4},{6,6,8},{2,1,1}}))
	fmt.Println(longestIncreasingPath([][]int{{3,4,5},{3,2,6},{2,2,1}}))
	//fmt.Println(longestIncreasingPath([][]int{{3,4,5}}))
	fmt.Println(longestIncreasingPath([][]int{{1,2}}))
}
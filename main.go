package main

import (
	"bufio"
	"fmt"
	"github.com/ccqy66/leetcode/dp"
	"os"
	"strconv"
	"strings"
)

func main() {
	row, cow := 0, 0
	fmt.Scanln(&row, &cow)
	if row <= 0 || cow <= 0 {
		fmt.Println(0)
		return
	}
	inputArr := make([][]int, row)
	input := bufio.NewScanner(os.Stdin)
	for i := 0; i < row; i++ {
		input.Scan()
		nums := strings.Split(input.Text(), " ")
		inputArr[i] = make([]int, len(nums))
		for j := 0; j < len(nums); j++ {
			inputArr[i][j], _ = strconv.Atoi(nums[j])
		}
	}
	fmt.Println(dp.MaxLengthPath(inputArr))
}

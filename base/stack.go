package base

import "fmt"

func monotoneIncreasingStack(arr []int) {
	increaseStack := make([]int, 0)
	for _, item := range arr {
		for len(increaseStack) > 0 && item > increaseStack[len(increaseStack)-1] {
			increaseStack = increaseStack[:len(increaseStack)-1]
		}
		increaseStack = append(increaseStack, item)
		fmt.Println(increaseStack)
	}
}

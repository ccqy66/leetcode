package main

import "fmt"

func reverseString(str string) string {
	if len(str) < 2 {
		return str
	}
	return fmt.Sprintf("%s%c", reverseString(str[1:]), str[0])
}


func printArray(arr []int64) {
	if len(arr) <= 0 {
		return
	}
	printArray(arr[1:])
	fmt.Println(arr[0])
}
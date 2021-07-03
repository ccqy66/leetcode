package base

import (
	"container/list"
	"fmt"
)

// 单调递增
// 从队列头到尾递增
// 头[3,2,1]尾
func monotoneIncreasingQueue(arr []int) {
	queue := list.New()
	for _, item :=range arr{
		// 剔除不符合条件的节点
		for queue.Len() > 0 && item > queue.Back().Value.(int)  {
			queue.Remove(queue.Back())
		}
		queue.PushBack(item)
		printQueue(queue)
	}
}

func maxWindows() {

}
func printQueue(list2 *list.List) {
	if list2.Len() <= 0 {
		fmt.Print("<empty>")
		return
	}
	for i := list2.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, ",")
	}
	fmt.Println()
}
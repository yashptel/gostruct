package main

import (
	"fmt"

	"github.com/yashptel/gostruct/pkg/heap"
)

func main() {
	a := &[]int{4, 6, 2, 7, 8, 1}
	h := heap.NewHeap(a)
	h.Push(3)
	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}
}

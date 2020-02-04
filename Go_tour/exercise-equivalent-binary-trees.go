package main

import "golang.org/x/tour/tree"
import "fmt"
import "sort"
import "reflect"

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	// 1. implement
	if t.Left != nil {
		go Walk(t.Left, ch)
	}
	if t.Right != nil {
		go Walk(t.Right, ch)
	}
	ch <- t.Value
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	arr1 := make([]int, 10)
	arr2 := make([]int, 10)
	for i := range arr1 {
		arr1[i] = <-ch1
	}
	for i := range arr2 {
		arr2[i] = <-ch2
	}
	sort.Ints(arr1)
	sort.Ints(arr2)

	return reflect.DeepEqual(arr1, arr2)

}

func main() {

	// 2. test
	// ch := make(chan int)
	// go Walk(tree.New(1), ch)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<- ch)
	// }

	fmt.Println(Same(tree.New(1), tree.New(1)))
}

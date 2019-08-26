package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

/**
function walk 用递归方法遍历二叉树
在 walk 的外面再包裹一层 Walk 是为了 close(ch)
而 close(ch) 是为了判断这棵树是否已经遍历完毕，以便结束循环
*/
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

func areSame(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 || !ok2 { //one of them finished
			return ok1 == ok2 //they finished at the same time or not
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Println(areSame(tree.New(1), tree.New(1)))
	fmt.Println(areSame(tree.New(2), tree.New(2)))
}

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// This uses a stack (slice) to traverse the tree repetitively.
// As mentioned in our call and suggested to use recurcive.
// This avoids recursion by using a stack to simulate the recursive call
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		close(ch)
		return
	}

	// use an explicit stack to traverse the tree repetitively
	stack := []*tree.Tree{}
	current := t

	// loop for as long as there are nodes to process or the stack is not empty
	for current != nil || len(stack) > 0 {
		// traverse the leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// set value to the last node on the stack (the left-most node)
		current = stack[len(stack)-1]

		// set current node to its right child
		stack = stack[:len(stack)-1]

		ch <- current.Value
		current = current.Right
	}

	close(ch)
}

// Compare two binary trees if they contain the same values in the same order
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 {
			return false
		}
	}

	_, ok := <-ch2
	return !ok
}

func main() {
	// The `Walk` function performs an in-order traversal of the binary tree and sends the values of each node to a channel
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for value := range ch {
		fmt.Println(value)
	}

	// quick testing
	fmt.Println(Same(tree.New(1), tree.New(1))) // should print true
	fmt.Println(Same(tree.New(1), tree.New(2))) // should print false
}

package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func main() {

	t := new(tree)
	arr := []int{5, 4, 3, 2, 6}

	t = add(t, 1)
	fmt.Println(t)
	Sort(arr)
	fmt.Println(arr)

}

func add(t *tree, value int) *tree {

	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

func Sort(values []int) {
	var root *tree

	for _, v := range values {
		root = add(root, v)
	}

	appendValues(values[:0], root)
}

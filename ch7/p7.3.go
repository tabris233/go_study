package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)

	fmt.Println("tree.String()", root)
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

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
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

func (t *tree) String() string {
	if t == nil {
		return "{}"
	}

	var buf bytes.Buffer
	buf.WriteString("{")

	var queue []*tree
	queue = append(queue, t)
	flag := true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur != nil {
			queue = append(queue, cur.left)
			queue = append(queue, cur.right)
			if flag {
				buf.WriteString(fmt.Sprintf("%d", cur.value))
			} else {
				buf.WriteString(fmt.Sprintf(" %d", cur.value))
			}
		} else {
			buf.WriteString(fmt.Sprintf(" nil"))
		}
		flag = false
	}

	buf.WriteString("}")
	return buf.String()
}

func main() {
	a := [...]int{4, 6, 6, 1, 8, 9, 1, 3, 7, 5, 3, 6}

	Sort(a[:])
	fmt.Println(a)
}

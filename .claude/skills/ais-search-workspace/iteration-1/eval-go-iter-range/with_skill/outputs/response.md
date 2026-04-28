# Go iter Package & Range-Over-Func: Generic Tree Walker

## Version

`range-over-func` and the `iter` package became **stable in Go 1.23** (August 2024). They were experimental in Go 1.22 behind `GOEXPERIMENT=rangefunc`.

## How it works

`iter.Seq[T]` is a function type: `func(yield func(T) bool)`. The `yield` callback is called for each element — returning `false` signals the iterator to stop (triggered by `break` or `return` in the `for range` loop).

## Working Implementation: Generic Binary Tree

```go
package main

import (
	"fmt"
	"iter" // Requires Go 1.23+
)

// Tree is a generic binary search tree.
type Tree[E any] struct {
	Val   E
	Left  *Tree[E]
	Right *Tree[E]
}

// Insert adds a value using a caller-supplied comparator.
func (t *Tree[E]) Insert(val E, less func(E, E) bool) *Tree[E] {
	if t == nil {
		return &Tree[E]{Val: val}
	}
	if less(val, t.Val) {
		t.Left = t.Left.Insert(val, less)
	} else {
		t.Right = t.Right.Insert(val, less)
	}
	return t
}

// push performs in-order traversal, propagating early-exit via the bool return.
func (t *Tree[E]) push(yield func(E) bool) bool {
	if t == nil {
		return true
	}
	if !t.Left.push(yield) {
		return false
	}
	if !yield(t.Val) {
		return false
	}
	return t.Right.push(yield)
}

// All returns an iter.Seq[E] for use with for range.
func (t *Tree[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		t.push(yield)
	}
}

func main() {
	var intTree *Tree[int]
	for _, v := range []int{5, 3, 7, 2, 4, 6, 8} {
		intTree = intTree.Insert(v, func(a, b int) bool { return a < b })
	}

	fmt.Print("In-order: ")
	for val := range intTree.All() {
		fmt.Printf("%d ", val) // 2 3 4 5 6 7 8
	}
	fmt.Println()

	// Early exit via break — correctly stops iteration
	fmt.Print("First 3:  ")
	count := 0
	for val := range intTree.All() {
		fmt.Printf("%d ", val)
		if count++; count == 3 {
			break
		}
	}
	fmt.Println()
}
```

## Caveats for Production Use

1. **Panic in yield** — If the loop body panics, the panic propagates through the iterator. Ensure your `push` function doesn't swallow it.
2. **Don't call yield after it returns false** — Once yield returns false (loop broke), calling it again causes a panic (`"iterator protocol violation"`).
3. **Not safe for concurrent modification** — The tree must not be modified while iterating.
4. **Stack depth** — Recursive `push` on a deeply unbalanced tree can overflow the stack. For production, consider an iterative traversal with an explicit stack.
5. **`iter.Seq2[K, V]`** — If you need key-value pairs (e.g., index + value), use `iter.Seq2[int, E]` instead.

*(Sources retrieved via `ais` CLI — see sources.txt)*

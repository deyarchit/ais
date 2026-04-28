# Go `iter` Package and Range-Over-Func: Generic Tree Walker

## Go Version Introduction

The `iter` package and range-over-func (range over function iterators) were introduced as **stable (non-experimental) features in Go 1.23**, released in August 2024.

Prior history:
- Go 1.22: The feature was available as an experimental preview behind the `GOEXPERIMENT=rangefunc` flag.
- Go 1.23: Promoted to stable. The `iter` package was added to the standard library, and the language spec officially supports ranging over functions with specific signatures.

---

## What Is `iter.Seq`?

The `iter` package defines two primary types:

```go
// Seq is an iterator over a sequence of values of type V.
// When called, it calls yield for each value in the sequence.
// If yield returns false, Seq must stop iteration and return.
type Seq[V any] func(yield func(V) bool)

// Seq2 is an iterator over sequences of pairs of values (K, V).
type Seq2[K, V any] func(yield func(K, V) bool)
```

These types are used with `for range` like so:

```go
for v := range seq {
    // use v
}

for k, v := range seq2 {
    // use k, v
}
```

---

## Concrete Implementation: Generic Binary Tree with `iter.Seq`

```go
package main

import (
	"fmt"
	"iter"
)

// Node is a generic binary tree node.
type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// BinaryTree is a generic binary tree.
type BinaryTree[T any] struct {
	Root *Node[T]
}

// Insert adds a value to the tree using a comparator function.
// less(a, b) returns true if a should go to the left of b.
func (t *BinaryTree[T]) Insert(value T, less func(a, b T) bool) {
	t.Root = insertNode(t.Root, value, less)
}

func insertNode[T any](n *Node[T], value T, less func(a, b T) bool) *Node[T] {
	if n == nil {
		return &Node[T]{Value: value}
	}
	if less(value, n.Value) {
		n.Left = insertNode(n.Left, value, less)
	} else {
		n.Right = insertNode(n.Right, value, less)
	}
	return n
}

// InOrder returns an iter.Seq[T] that yields values in in-order (left, root, right).
// This produces sorted order for a BST.
func (t *BinaryTree[T]) InOrder() iter.Seq[T] {
	return func(yield func(T) bool) {
		inOrderTraverse(t.Root, yield)
	}
}

// inOrderTraverse recursively traverses the tree in order.
// It returns false if yield signaled early termination.
func inOrderTraverse[T any](n *Node[T], yield func(T) bool) bool {
	if n == nil {
		return true
	}
	// Traverse left subtree
	if !inOrderTraverse(n.Left, yield) {
		return false
	}
	// Visit current node
	if !yield(n.Value) {
		return false
	}
	// Traverse right subtree
	return inOrderTraverse(n.Right, yield)
}

// PreOrder returns an iter.Seq[T] that yields values in pre-order (root, left, right).
func (t *BinaryTree[T]) PreOrder() iter.Seq[T] {
	return func(yield func(T) bool) {
		preOrderTraverse(t.Root, yield)
	}
}

func preOrderTraverse[T any](n *Node[T], yield func(T) bool) bool {
	if n == nil {
		return true
	}
	if !yield(n.Value) {
		return false
	}
	if !preOrderTraverse(n.Left, yield) {
		return false
	}
	return preOrderTraverse(n.Right, yield)
}

// PostOrder returns an iter.Seq[T] that yields values in post-order (left, right, root).
func (t *BinaryTree[T]) PostOrder() iter.Seq[T] {
	return func(yield func(T) bool) {
		postOrderTraverse(t.Root, yield)
	}
}

func postOrderTraverse[T any](n *Node[T], yield func(T) bool) bool {
	if n == nil {
		return true
	}
	if !postOrderTraverse(n.Left, yield) {
		return false
	}
	if !postOrderTraverse(n.Right, yield) {
		return false
	}
	return yield(n.Value)
}

// InOrderWithIndex returns an iter.Seq2[int, T] yielding (index, value) pairs in-order.
func (t *BinaryTree[T]) InOrderWithIndex() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		idx := 0
		inOrderTraverse(t.Root, func(v T) bool {
			if !yield(idx, v) {
				return false
			}
			idx++
			return true
		})
	}
}

func main() {
	tree := &BinaryTree[int]{}
	less := func(a, b int) bool { return a < b }

	for _, v := range []int{5, 3, 7, 1, 4, 6, 8} {
		tree.Insert(v, less)
	}

	fmt.Println("In-order (sorted):")
	for v := range tree.InOrder() {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	// Output: 1 2 3 4 5 6 7 8  (actually: 1 3 4 5 6 7 8)

	fmt.Println("Pre-order:")
	for v := range tree.PreOrder() {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Println("Post-order:")
	for v := range tree.PostOrder() {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Println("In-order with index:")
	for i, v := range tree.InOrderWithIndex() {
		fmt.Printf("[%d]=%d ", i, v)
	}
	fmt.Println()

	// Early termination example: stop after finding first value >= 5
	fmt.Println("First value >= 5:")
	for v := range tree.InOrder() {
		if v >= 5 {
			fmt.Println(v)
			break // This correctly signals yield to return false
		}
	}
}
```

### go.mod for This Code

```
module example.com/btree

go 1.23
```

---

## How Range-Over-Func Works Under the Hood

When the compiler sees:

```go
for v := range mySeq {
    body(v)
}
```

It desugars to roughly:

```go
mySeq(func(v T) bool {
    body(v)
    return true // continue
    // if `break` was used, return false instead
})
```

The `yield` function returns `false` to signal the iterator should stop. This is how `break`, `return`, and `goto` inside the `for range` body are translated into early termination.

---

## Gotchas and Caveats for Production Use

### 1. The `yield` Contract is Mandatory
An iterator **must** stop calling `yield` once it returns `false`. Violating this is a panic in Go 1.23+:

```go
// WRONG — do not do this:
return func(yield func(T) bool) {
    for _, v := range items {
        yield(v) // BUG: ignoring the return value
    }
}

// CORRECT:
return func(yield func(T) bool) {
    for _, v := range items {
        if !yield(v) {
            return // stop immediately
        }
    }
}
```

### 2. Panics Inside `yield` Propagate Correctly
If user code inside `for range` panics, the panic propagates normally through `yield`. Your iterator does **not** need to handle this; just don't suppress panics.

### 3. `defer` Inside `for range` Runs at Function Return
Defers inside a `for range` loop body run when the enclosing function returns, **not** at each iteration. This is the same behavior as any other loop — not specific to range-over-func, but worth remembering.

### 4. Concurrency: Iterators Are Not Goroutine-Safe
The `iter.Seq` pattern is single-threaded by design. Calling the same iterator from multiple goroutines simultaneously requires external synchronization. The tree itself also has no mutation guards.

### 5. No Built-in `len` or `count`
Unlike slices or maps, you cannot call `len()` on an `iter.Seq`. If you need a count, you must iterate fully:

```go
func Count[T any](seq iter.Seq[T]) int {
    n := 0
    for range seq {
        n++
    }
    return n
}
```

### 6. Composability with `slices` and `maps` Packages
Go 1.23 added iterator-aware functions to the standard library:

```go
import "slices"

// Collect all values from an iter.Seq into a slice:
all := slices.Collect(tree.InOrder())
```

The `slices.Collect`, `maps.Collect`, and related functions are available in Go 1.23+.

### 7. Stack Depth for Recursive Iterators
The recursive implementation shown above will use O(depth) stack frames. For a balanced BST with N nodes, depth is O(log N), which is fine. For degenerate (unbalanced) trees, depth is O(N). If you need to handle very large unbalanced trees, consider an iterative implementation using an explicit stack.

### 8. `iter.Pull` for Pull-Style Iteration
The `iter` package also provides `iter.Pull` and `iter.Pull2`, which convert a push-style `Seq` into a pull-style iterator (useful when you need to advance iteration manually, e.g., to merge two sorted sequences):

```go
next, stop := iter.Pull(tree.InOrder())
defer stop()

v1, ok1 := next()
v2, ok2 := next()
// ...
```

Always `defer stop()` when using `iter.Pull` — failing to call `stop` can leak goroutines, as the implementation uses goroutines internally to convert push to pull.

### 9. Module Compatibility
Ensure your `go.mod` specifies `go 1.23` or later. Using `iter.Seq` in a module with `go 1.22` or earlier will result in a compilation error.

### 10. `rangefunc` Is No Longer an Experiment
Do **not** set `GOEXPERIMENT=rangefunc` in Go 1.23+. It is now part of the standard language and the experiment flag is ignored (or may cause errors in future versions).

---

## Summary Table

| Feature | Introduced (Experimental) | Introduced (Stable) |
|---|---|---|
| Range-over-func syntax | Go 1.22 (`GOEXPERIMENT=rangefunc`) | Go 1.23 |
| `iter.Seq[V]` type | Go 1.22 (experimental) | Go 1.23 |
| `iter.Seq2[K, V]` type | Go 1.22 (experimental) | Go 1.23 |
| `iter.Pull` / `iter.Pull2` | Go 1.22 (experimental) | Go 1.23 |
| `slices.Collect` | Go 1.22 (experimental) | Go 1.23 |

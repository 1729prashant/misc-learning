/*
101. Symmetric Tree

Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

Example 1:

Input: root = [1,2,2,3,4,4,3]
Output: true
Example 2:

Input: root = [1,2,2,null,3,null,3]
Output: false

Constraints:

The number of nodes in the tree is in the range [1, 1000].
-100 <= Node.val <= 100

Follow up: Could you solve it both recursively and iteratively?
*/
package main

import (
	"fmt"
	"testing"
)

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetricRecursive checks if a tree is symmetric using recursion
func isSymmetricRecursive(root *TreeNode) bool {
	var isMirror func(t1, t2 *TreeNode) bool
	isMirror = func(t1, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		}
		if t1 == nil || t2 == nil || t1.Val != t2.Val {
			return false
		}
		return isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
	}
	return isMirror(root, root)
}

func isSymmetricIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		t1 := queue[0]
		t2 := queue[1]
		queue = queue[2:]

		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil || t1.Val != t2.Val {
			return false
		}

		queue = append(queue, t1.Left, t2.Right)
		queue = append(queue, t1.Right, t2.Left)
	}
	return true
}

// Helper function to create a binary tree from a slice
func createTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// Test cases
func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name     string
		tree     []interface{}
		expected bool
	}{
		{"Symmetric tree", []interface{}{1, 2, 2, 3, 4, 4, 3}, true},
		{"Not symmetric", []interface{}{1, 2, 2, nil, 3, nil, 3}, false},
		{"Single node", []interface{}{1}, true},
		{"Empty tree", []interface{}{}, true},
		{"Unbalanced symmetric tree", []interface{}{1, 2, 2, 3, nil, nil, 3}, true},
		{"Tree with negative values", []interface{}{1, -2, -2, -3, nil, nil, -3}, true},
		{"Large symmetric tree", []interface{}{1, 2, 2, 3, 4, 4, 3, 5, nil, nil, 5, 5, nil, nil, 5}, true},
		{"Large asymmetric tree", []interface{}{1, 2, 2, 3, 4, 4, 3, 5, nil, nil, 6, 5, nil, nil, 5}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTree(tt.tree)

			// Test recursive solution
			if got := isSymmetricRecursive(root); got != tt.expected {
				t.Errorf("isSymmetricRecursive(%v) = %v, want %v", tt.tree, got, tt.expected)
			}

			// Test iterative solution
			if got := isSymmetricIterative(root); got != tt.expected {
				t.Errorf("isSymmetricIterative(%v) = %v, want %v", tt.tree, got, tt.expected)
			}
		})
	}
}

func main() {
	fmt.Println("Run tests with `go test`.")
}

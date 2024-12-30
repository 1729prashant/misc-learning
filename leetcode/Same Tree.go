/*
100. Same Tree

Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.



Example 1:


Input: p = [1,2,3], q = [1,2,3]
Output: true
Example 2:


Input: p = [1,2], q = [1,null,2]
Output: false
Example 3:


Input: p = [1,2,1], q = [1,1,2]
Output: false


Constraints:

The number of nodes in both trees is in the range [0, 100].
-104 <= Node.val <= 104
*/

package main

import (
	"testing"
)

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSameTree checks if two trees are the same
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// If both nodes are nil, they are the same
	if p == nil && q == nil {
		return true
	}
	// If one is nil and the other isn't, they're not the same
	if p == nil || q == nil {
		return false
	}
	// If the values are different, they're not the same tree
	if p.Val != q.Val {
		return false
	}
	// Recursively check if left and right subtrees are the same
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// Helper function to create tree from slice
func createTree(values []int) *TreeNode {
	if len(values) == 0 || values[0] == -1 {
		return nil
	}
	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}
	i := 1
	for i < len(values) {
		node := queue[0]
		queue = queue[1:]
		if i < len(values) && values[i] != -1 {
			node.Left = &TreeNode{Val: values[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(values) && values[i] != -1 {
			node.Right = &TreeNode{Val: values[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// Test cases
func TestIsSameTree(t *testing.T) {
	// Helper function to assert tree equality
	assertTreeEqual := func(t *testing.T, p, q []int, expected bool) {
		treeP := createTree(p)
		treeQ := createTree(q)
		if isSameTree(treeP, treeQ) != expected {
			t.Errorf("isSameTree(%v, %v) = %v, want %v", p, q, !expected, expected)
		}
	}

	// Test case 1: [1,2,3], [1,2,3]
	assertTreeEqual(t, []int{1, 2, 3}, []int{1, 2, 3}, true)

	// Test case 2: [1,2], [1,null,2]
	assertTreeEqual(t, []int{1, 2}, []int{1, -1, 2}, false)

	// Test case 3: [1,2,1], [1,1,2]
	assertTreeEqual(t, []int{1, 2, 1}, []int{1, 1, 2}, false)

	// Test case 4: Both trees are nil
	assertTreeEqual(t, []int{}, []int{}, true)

	// Test case 5: One tree is nil, other isn't
	assertTreeEqual(t, []int{1}, []int{}, false)

	// Test case 6: Different structure, some matching nodes
	assertTreeEqual(t, []int{1, 2, 3, -1, -1, 4}, []int{1, 2, 3, 4}, false)

	// Test case 7: Deep tree test
	assertTreeEqual(t, []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}, true)
}

func main() {
	// TODO tests with 'go test' command if this file is named test_is_same_tree.go
}

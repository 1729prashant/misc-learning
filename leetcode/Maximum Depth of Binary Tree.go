/*
104. Maximum Depth of Binary Tree

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: 3
Example 2:

Input: root = [1,null,2]
Output: 2


Constraints:

The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100
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

// maxDepth calculates the maximum depth of the binary tree
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1
}

// Helper function to return the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test cases
func TestMaxDepth(t *testing.T) {
	// Helper function to create tree nodes
	newNode := func(val int) *TreeNode {
		return &TreeNode{Val: val}
	}

	// Test case 1: [3,9,20,null,null,15,7]
	root1 := newNode(3)
	root1.Left = newNode(9)
	root1.Right = newNode(20)
	root1.Right.Left = newNode(15)
	root1.Right.Right = newNode(7)
	if got := maxDepth(root1); got != 3 {
		t.Errorf("maxDepth([3,9,20,null,null,15,7]) = %d; want 3", got)
	}

	// Test case 2: [1,null,2]
	root2 := newNode(1)
	root2.Right = newNode(2)
	if got := maxDepth(root2); got != 2 {
		t.Errorf("maxDepth([1,null,2]) = %d; want 2", got)
	}

	// Test case 3: Empty tree
	if got := maxDepth(nil); got != 0 {
		t.Errorf("maxDepth(nil) = %d; want 0", got)
	}

	// Test case 4: Single node tree
	root4 := newNode(1)
	if got := maxDepth(root4); got != 1 {
		t.Errorf("maxDepth([1]) = %d; want 1", got)
	}

	// Test case 5: Unbalanced tree [1,2,null,3,null,4,null,5]
	root5 := newNode(1)
	root5.Left = newNode(2)
	root5.Left.Left = newNode(3)
	root5.Left.Left.Left = newNode(4)
	root5.Left.Left.Left.Left = newNode(5)
	if got := maxDepth(root5); got != 5 {
		t.Errorf("maxDepth([1,2,null,3,null,4,null,5]) = %d; want 5", got)
	}
}

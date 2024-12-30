/*
226. Invert Binary Tree

Given the root of a binary tree, invert the tree, and return its root.

Example 1:

Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
Example 2:

Input: root = [2,1,3]
Output: [2,3,1]
Example 3:

Input: root = []
Output: []

Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
*/
package main

import (
	"reflect"
	"testing"
)

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree inverts a binary tree
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Swap the left and right children
	root.Left, root.Right = root.Right, root.Left

	// Recursively invert the left subtree
	invertTree(root.Left)
	// Recursively invert the right subtree
	invertTree(root.Right)

	return root
}

// Helper function to create tree from slice
func createTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}
	root := &TreeNode{Val: values[0].(int)} // Type assertion to int
	queue := []*TreeNode{root}
	i := 1
	for i < len(values) {
		node := queue[0]
		queue = queue[1:]
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)} // Type assertion
			queue = append(queue, node.Left)
		}
		i++
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)} // Type assertion
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// Helper function to convert tree back to slice for comparison
func treeToSlice(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}
	result := []interface{}{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		} else {
			result = append(result, nil)
		}
	}
	// Trim trailing nils
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != nil {
			break
		}
		result = result[:i]
	}
	return result
}

// Test cases
func TestInvertTree(t *testing.T) {
	assertTreeEqual := func(t *testing.T, input, expected []interface{}) {
		tree := createTree(input)
		inverted := invertTree(tree)
		result := treeToSlice(inverted)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("invertTree(%v) = %v, want %v", input, result, expected)
		}
	}

	assertTreeEqual(t, []interface{}{4, 2, 7, 1, 3, 6, 9}, []interface{}{4, 7, 2, 9, 6, 3, 1})
	assertTreeEqual(t, []interface{}{2, 1, 3}, []interface{}{2, 3, 1})
	assertTreeEqual(t, []interface{}{}, []interface{}{})
	assertTreeEqual(t, []interface{}{1}, []interface{}{1})
	assertTreeEqual(t, []interface{}{1, 2, nil, 3, nil, 4}, []interface{}{1, nil, 2, nil, 4, 3})
}

func main() {
	// TODO tests with 'go test' command if this file is named test_invert_tree.go
}

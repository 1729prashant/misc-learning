// Package traversal provides binary tree traversal implementations
package traversal

// TreeNode represents a node in a binary tree, see treenode.go

// PreorderTraversal performs preorder traversal of a binary tree
// Algorithm P: Visit Root → Left Subtree → Right Subtree
// Time Complexity: O(n), Space Complexity: O(h) where h is height
func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	preorderHelper(root, &result)
	return result
}

// preorderHelper is the recursive helper function for preorder traversal
// Following Knuth's Algorithm P:
// P1: Check if current node is nil
// P2: Visit current node (add to result)
// P3: Recursively traverse left subtree
// P4: Recursively traverse right subtree
func preorderHelper(node *TreeNode, result *[]int) {
	// P1: Base case - if node is nil, return
	if node == nil {
		return
	}

	// P2: Visit current node - add value to result
	*result = append(*result, node.Val)

	// P3: Traverse left subtree recursively
	preorderHelper(node.Left, result)

	// P4: Traverse right subtree recursively
	preorderHelper(node.Right, result)
}

// PreorderTraversalIterative provides an iterative implementation using stack
// This is an alternative implementation for educational purposes
func PreorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		// Pop from stack
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Visit node
		result = append(result, node.Val)

		// Push right first, then left (so left is processed first)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

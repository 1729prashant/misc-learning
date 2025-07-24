// Package traversal provides binary tree traversal implementations
package traversal

// InorderTraversal performs inorder traversal of a binary tree
// Algorithm I: Visit Left Subtree → Root → Right Subtree
// Time Complexity: O(n), Space Complexity: O(h) where h is height
// Special property: For BST, returns values in ascending sorted order
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	inorderHelper(root, &result)
	return result
}

// inorderHelper is the recursive helper function for inorder traversal
// Following Knuth's Algorithm I:
// I1: Check if current node is nil
// I2: Recursively traverse left subtree
// I3: Visit current node (add to result)
// I4: Recursively traverse right subtree
func inorderHelper(node *TreeNode, result *[]int) {
	// I1: Base case - if node is nil, return
	if node == nil {
		return
	}

	// I2: Traverse left subtree recursively first
	inorderHelper(node.Left, result)

	// I3: Visit current node - add value to result
	*result = append(*result, node.Val)

	// I4: Traverse right subtree recursively
	inorderHelper(node.Right, result)
}

// InorderTraversalIterative provides an iterative implementation using stack
// This follows the same left-root-right pattern without recursion
func InorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	var stack []*TreeNode
	current := root

	for current != nil || len(stack) > 0 {
		// Go to the leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// Current must be nil at this point
		// Pop from stack and visit
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		// Move to right subtree
		current = current.Right
	}

	return result
}

// InorderMorrisTraversal provides Morris Traversal implementation
// Space complexity: O(1) - no recursion or stack needed
// Uses threading technique to traverse without extra space
func InorderMorrisTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	current := root

	for current != nil {
		if current.Left == nil {
			// No left subtree, visit current and go right
			result = append(result, current.Val)
			current = current.Right
		} else {
			// Find inorder predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				// Make current the right child of its predecessor
				predecessor.Right = current
				current = current.Left
			} else {
				// Revert the changes made - remove the link
				predecessor.Right = nil
				result = append(result, current.Val)
				current = current.Right
			}
		}
	}

	return result
}

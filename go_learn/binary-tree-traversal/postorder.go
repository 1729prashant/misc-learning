// Package traversal provides binary tree traversal implementations
package traversal

// PostorderTraversal performs postorder traversal of a binary tree
// Algorithm O: Visit Left Subtree → Right Subtree → Root
// Time Complexity: O(n), Space Complexity: O(h) where h is height
// Useful for: deleting nodes, calculating directory sizes, expression evaluation
func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	postorderHelper(root, &result)
	return result
}

// postorderHelper is the recursive helper function for postorder traversal
// Following Knuth's Algorithm O:
// O1: Check if current node is nil
// O2: Recursively traverse left subtree
// O3: Recursively traverse right subtree
// O4: Visit current node (add to result)
func postorderHelper(node *TreeNode, result *[]int) {
	// O1: Base case - if node is nil, return
	if node == nil {
		return
	}

	// O2: Traverse left subtree recursively first
	postorderHelper(node.Left, result)

	// O3: Traverse right subtree recursively
	postorderHelper(node.Right, result)

	// O4: Visit current node - add value to result (after children)
	*result = append(*result, node.Val)
}

// PostorderTraversalIterative provides an iterative implementation using two stacks
// This approach uses two stacks to simulate the recursive behavior
func PostorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	stack1 := []*TreeNode{root}
	var stack2 []*TreeNode

	// First stack is used to traverse the tree
	// Second stack is used to store nodes in reverse postorder
	for len(stack1) > 0 {
		node := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, node)

		// Push left first, then right
		if node.Left != nil {
			stack1 = append(stack1, node.Left)
		}
		if node.Right != nil {
			stack1 = append(stack1, node.Right)
		}
	}

	// Pop all items from second stack to get postorder
	for len(stack2) > 0 {
		node := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		result = append(result, node.Val)
	}

	return result
}

// PostorderTraversalSingleStack provides iterative implementation with single stack
// More complex but uses only one stack with a visited flag approach
func PostorderTraversalSingleStack(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	stack := []*TreeNode{}
	var lastVisited *TreeNode
	current := root

	for len(stack) > 0 || current != nil {
		if current != nil {
			// Go to the leftmost node
			stack = append(stack, current)
			current = current.Left
		} else {
			// Peek at the top of stack
			peekNode := stack[len(stack)-1]

			// If right child exists and hasn't been processed yet
			if peekNode.Right != nil && lastVisited != peekNode.Right {
				current = peekNode.Right
			} else {
				// Visit the node
				result = append(result, peekNode.Val)
				lastVisited = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		}
	}

	return result
}

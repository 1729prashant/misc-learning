package traversal

import (
	"fmt"
	"reflect"
	"testing"
)

// TestPreorderTraversal tests the recursive preorder traversal implementation
func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
		{
			name: "Complete binary tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name: "Left skewed tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
			expected: []int{1, 2, 3, 4},
		},
		{
			name: "Right skewed tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val:   3,
						Right: &TreeNode{Val: 4},
					},
				},
			},
			expected: []int{1, 2, 3, 4},
		},
		{
			name: "Binary Search Tree",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
			expected: []int{5, 3, 2, 4, 8, 7, 9},
		},
		{
			name: "Unbalanced tree - only left children",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "Unbalanced tree - only right children",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PreorderTraversal(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PreorderTraversal() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestPreorderTraversalIterative tests the iterative preorder traversal implementation
func TestPreorderTraversalIterative(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
		{
			name: "Complete binary tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: []int{1, 2, 4, 5, 3, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PreorderTraversalIterative(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PreorderTraversalIterative() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestPreorderConsistency ensures recursive and iterative implementations produce same results
func TestPreorderConsistency(t *testing.T) {
	testCases := []*TreeNode{
		nil,
		&TreeNode{Val: 1},
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 7},
			},
		},
	}

	for i, root := range testCases {
		t.Run(fmt.Sprintf("Consistency test %d", i), func(t *testing.T) {
			recursive := PreorderTraversal(root)
			iterative := PreorderTraversalIterative(root)

			if !reflect.DeepEqual(recursive, iterative) {
				t.Errorf("Recursive and iterative results differ: recursive=%v, iterative=%v",
					recursive, iterative)
			}
		})
	}
}

// BenchmarkPreorderTraversal benchmarks the recursive implementation
func BenchmarkPreorderTraversal(b *testing.B) {
	// Create a larger test tree
	root := createLargeTreePreOrder(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PreorderTraversal(root)
	}
}

// BenchmarkPreorderTraversalIterative benchmarks the iterative implementation
func BenchmarkPreorderTraversalIterative(b *testing.B) {
	// Create a larger test tree
	root := createLargeTreePreOrder(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PreorderTraversalIterative(root)
	}
}

// createLargeTreePreOrder creates a complete binary tree of given depth for benchmarking
func createLargeTreePreOrder(depth int) *TreeNode {
	if depth <= 0 {
		return nil
	}

	return &TreeNode{
		Val:   depth,
		Left:  createLargeTreePreOrder(depth - 1),
		Right: createLargeTreePreOrder(depth - 1),
	}
}

// Note: fmt import should be at the top with other imports

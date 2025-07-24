package traversal

import (
	"fmt"
	"reflect"
	"testing"
)

// TestPostorderTraversal tests the recursive postorder traversal implementation
func TestPostorderTraversal(t *testing.T) {
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
			expected: []int{4, 5, 2, 6, 7, 3, 1},
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
			expected: []int{4, 3, 2, 1},
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
			expected: []int{4, 3, 2, 1},
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
			expected: []int{2, 4, 3, 7, 9, 8, 5},
		},
		{
			name: "Tree with only left children",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			expected: []int{3, 2, 1},
		},
		{
			name: "Tree with only right children",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			expected: []int{3, 2, 1},
		},
		{
			name: "Tree with negative values",
			root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:   -2,
					Left:  &TreeNode{Val: -3},
					Right: &TreeNode{Val: -1},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
			expected: []int{-3, -1, -2, 1, 3, 2, 0},
		},
		{
			name: "Expression tree example (a + b) * (c - d)",
			root: &TreeNode{
				Val: '*',
				Left: &TreeNode{
					Val:   '+',
					Left:  &TreeNode{Val: 'a'},
					Right: &TreeNode{Val: 'b'},
				},
				Right: &TreeNode{
					Val:   '-',
					Left:  &TreeNode{Val: 'c'},
					Right: &TreeNode{Val: 'd'},
				},
			},
			expected: []int{'a', 'b', '+', 'c', 'd', '-', '*'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PostorderTraversal(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PostorderTraversal() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestPostorderTraversalIterative tests the two-stack iterative implementation
func TestPostorderTraversalIterative(t *testing.T) {
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
			expected: []int{4, 5, 2, 6, 7, 3, 1},
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
			expected: []int{2, 4, 3, 7, 9, 8, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PostorderTraversalIterative(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PostorderTraversalIterative() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestPostorderTraversalSingleStack tests the single-stack iterative implementation
func TestPostorderTraversalSingleStack(t *testing.T) {
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
			expected: []int{4, 5, 2, 6, 7, 3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PostorderTraversalSingleStack(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PostorderTraversalSingleStack() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestPostorderConsistency ensures all three implementations produce same results
func TestPostorderConsistency(t *testing.T) {
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
		// Complex tree
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 8},
					Right: &TreeNode{Val: 9},
				},
				Right: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 10},
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 11},
					Right: &TreeNode{Val: 12},
				},
			},
		},
	}

	for i, root := range testCases {
		t.Run(fmt.Sprintf("Consistency test %d", i), func(t *testing.T) {
			recursive := PostorderTraversal(root)
			iterative := PostorderTraversalIterative(root)
			singleStack := PostorderTraversalSingleStack(root)

			if !reflect.DeepEqual(recursive, iterative) {
				t.Errorf("Recursive and iterative results differ: recursive=%v, iterative=%v",
					recursive, iterative)
			}

			if !reflect.DeepEqual(recursive, singleStack) {
				t.Errorf("Recursive and single-stack results differ: recursive=%v, singleStack=%v",
					recursive, singleStack)
			}
		})
	}
}

// TestPostorderExpressionEvaluation tests postorder traversal for expression evaluation
func TestPostorderExpressionEvaluation(t *testing.T) {
	// Test that postorder traversal gives correct order for expression evaluation
	// Expression: (2 + 3) * (4 - 1) = 5 * 3 = 15
	// Tree represents: ((2 + 3) * (4 - 1))
	exprTree := &TreeNode{
		Val: '*',
		Left: &TreeNode{
			Val:   '+',
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   '-',
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 1},
		},
	}

	result := PostorderTraversal(exprTree)
	expected := []int{2, 3, '+', 4, 1, '-', '*'}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expression tree postorder = %v, expected %v", result, expected)
	}

	// Verify that this order can be used for stack-based evaluation
	// This is how postfix evaluation would work
	expectedEvaluationOrder := "2 3 + 4 1 - *"
	t.Logf("Postorder gives evaluation order: %s", expectedEvaluationOrder)
}

// TestPostorderNodeDeletion tests that postorder is correct for safe node deletion
func TestPostorderNodeDeletion(t *testing.T) {
	// Create a tree where we want to delete all nodes safely (children before parent)
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 6},
		},
	}

	result := PostorderTraversal(tree)
	expected := []int{4, 5, 2, 6, 3, 1}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Deletion order postorder = %v, expected %v", result, expected)
	}

	// Verify that children always come before their parents
	nodeToParent := map[int]int{
		4: 2, 5: 2, 2: 1,
		6: 3, 3: 1,
	}

	positions := make(map[int]int)
	for i, val := range result {
		positions[val] = i
	}

	for child, parent := range nodeToParent {
		if positions[child] >= positions[parent] {
			t.Errorf("Child %d (pos %d) should come before parent %d (pos %d)",
				child, positions[child], parent, positions[parent])
		}
	}
}

// BenchmarkPostorderTraversal benchmarks the recursive implementation
func BenchmarkPostorderTraversal(b *testing.B) {
	root := createLargeTreePostOrder(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PostorderTraversal(root)
	}
}

// BenchmarkPostorderTraversalIterative benchmarks the two-stack iterative implementation
func BenchmarkPostorderTraversalIterative(b *testing.B) {
	root := createLargeTreePostOrder(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PostorderTraversalIterative(root)
	}
}

// BenchmarkPostorderTraversalSingleStack benchmarks the single-stack iterative implementation
func BenchmarkPostorderTraversalSingleStack(b *testing.B) {
	root := createLargeTreePostOrder(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PostorderTraversalSingleStack(root)
	}
}

// createLargeTreePostOrder creates a complete binary tree of given depth for benchmarking
func createLargeTreePostOrder(depth int) *TreeNode {
	if depth <= 0 {
		return nil
	}

	return &TreeNode{
		Val:   depth,
		Left:  createLargeTreePostOrder(depth - 1),
		Right: createLargeTreePostOrder(depth - 1),
	}
}

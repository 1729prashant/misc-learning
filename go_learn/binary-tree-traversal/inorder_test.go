package traversal

import (
	"reflect"
	"testing"
)

// ---- TreeNode definition ----
// is from treenode.go

// ---- Helper: Large complete binary tree ----

func createLargeTree(depth int) *TreeNode {
	if depth == 0 {
		return nil
	}
	node := &TreeNode{Val: 1}
	queue := []*TreeNode{node}
	val := 2
	for level := 1; level < depth; level++ {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			curr.Left = &TreeNode{Val: val}
			val++
			curr.Right = &TreeNode{Val: val}
			val++
			queue = append(queue, curr.Left, curr.Right)
		}
	}
	return node
}

// ---- Unit tests for each traversal ----

func TestInorderTraversal(t *testing.T) {
	for _, tc := range baseTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			result := InorderTraversal(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("InorderTraversal() = %v, expected %v", result, tc.expected)
			}
		})
	}
}

func TestInorderTraversalIterative(t *testing.T) {
	for _, tc := range baseTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			result := InorderTraversalIterative(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("InorderTraversalIterative() = %v, expected %v", result, tc.expected)
			}
		})
	}
}

func TestInorderMorrisTraversal(t *testing.T) {
	for _, tc := range baseTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			result := InorderMorrisTraversal(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("InorderMorrisTraversal() = %v, expected %v", result, tc.expected)
			}
		})
	}
}

// ---- Correctness & Consistency Tests ----

func TestInorderConsistency(t *testing.T) {
	for depth := 1; depth <= 12; depth++ {
		tree := createLargeTree(depth)
		r1 := InorderTraversal(tree)
		r2 := InorderTraversalIterative(tree)
		r3 := InorderMorrisTraversal(tree)

		if !reflect.DeepEqual(r1, r2) {
			t.Errorf("Mismatch between recursive and iterative at depth %d", depth)
		}
		if !reflect.DeepEqual(r1, r3) {
			t.Errorf("Mismatch between recursive and Morris at depth %d", depth)
		}
	}
}

func TestBSTProperty(t *testing.T) {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 8},
		},
		Right: &TreeNode{
			Val:   15,
			Left:  &TreeNode{Val: 12},
			Right: &TreeNode{Val: 20},
		},
	}
	result := InorderTraversal(root)
	for i := 1; i < len(result); i++ {
		if result[i] < result[i-1] {
			t.Errorf("Inorder is not sorted at index %d: %v", i, result)
			break
		}
	}
}

// ---- Benchmarks ----

func BenchmarkInorderTraversal(b *testing.B) {
	tree := createLargeTree(15)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InorderTraversal(tree)
	}
}

func BenchmarkInorderTraversalIterative(b *testing.B) {
	tree := createLargeTree(15)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InorderTraversalIterative(tree)
	}
}

func BenchmarkInorderMorrisTraversal(b *testing.B) {
	tree := createLargeTree(15)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InorderMorrisTraversal(tree)
	}
}

// ---- Test Cases ----

type testCase struct {
	name     string
	root     *TreeNode
	expected []int
}

func baseTestCases() []testCase {
	return []testCase{
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     &TreeNode{Val: 42},
			expected: []int{42},
		},
		{
			name: "Degenerate tree (like linked list)",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val:   4,
							Right: &TreeNode{Val: 5},
						},
					},
				},
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Balanced tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3},
			},
			expected: []int{4, 2, 5, 1, 3},
		},
		{
			name: "Non-BST structure",
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val:  20,
					Left: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   15,
					Right: &TreeNode{Val: 25},
				},
			},
			expected: []int{5, 20, 10, 15, 25},
		},
		{
			name: "Large value nodes",
			root: &TreeNode{
				Val: 1000000,
				Left: &TreeNode{
					Val:   500000,
					Left:  &TreeNode{Val: 250000},
					Right: &TreeNode{Val: 750000},
				},
				Right: &TreeNode{
					Val:   1500000,
					Left:  &TreeNode{Val: 1250000},
					Right: &TreeNode{Val: 1750000},
				},
			},
			expected: []int{250000, 500000, 750000, 1000000, 1250000, 1500000, 1750000},
		},
	}
}

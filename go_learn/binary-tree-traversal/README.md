# Binary Tree Traversal Implementations

This repository contains comprehensive implementations of the three fundamental binary tree traversal algorithms with Knuth-style algorithmic descriptions, flowcharts, Go implementations, and extensive unit tests.

## File Structure

```
traversal/
├── README.md                   # This file - usage guide and examples
├── algorithms.md               # algorithms with flowcharts
├── treenode.go                 # TreeNode declaration
├── preorder.go                 # Preorder traversal implementation
├── preorder_test.go            # Preorder traversal unit tests
├── inorder.go                  # Inorder traversal implementation  
├── inorder_test.go             # Inorder traversal unit tests
├── postorder.go                # Postorder traversal implementation
└── postorder_test.go           # Postorder traversal unit tests

```

## Running the Code

### Prerequisites

- Go 1.24.2 or higher
- Git (for cloning)

### Setup

1. Create a new Go module:
```bash
mkdir binary-tree-traversal
cd binary-tree-traversal
go mod init traversal
```

2. Copy all the `.go` files into the directory

3. Run tests for all traversal methods:
```bash
# Run all tests
go test -v

# Run tests for specific traversal
go test -v -run TestPreorder
go test -v -run TestInorder  
go test -v -run TestPostorder

# Run benchmarks
go test -bench=.

# Run tests with coverage
go test -cover
```

## Usage Examples

### Basic Usage

```go
package main

import (
    "fmt"
    "traversal"
)

func main() {
    // Create a sample binary tree:
    //       1
    //      / \
    //     2   3
    //    / \
    //   4   5
    
    root := &traversal.TreeNode{
        Val: 1,
        Left: &traversal.TreeNode{
            Val: 2,
            Left: &traversal.TreeNode{Val: 4},
            Right: &traversal.TreeNode{Val: 5},
        },
        Right: &traversal.TreeNode{Val: 3},
    }
    
    // Perform different traversals
    preorder := traversal.PreorderTraversal(root)
    inorder := traversal.InorderTraversal(root)
    postorder := traversal.PostorderTraversal(root)
    
    fmt.Printf("Preorder:  %v\n", preorder)  // [1 2 4 5 3]
    fmt.Printf("Inorder:   %v\n", inorder)   // [4 2 5 1 3]
    fmt.Printf("Postorder: %v\n", postorder) // [4 5 2 3 1]
}
```

### Working with Binary Search Trees

```go
func demonstrateBST() {
    // Create a Binary Search Tree:
    //       5
    //      / \
    //     3   8
    //    / \ / \
    //   2  4 7  9
    
    bst := &traversal.TreeNode{
        Val: 5,
        Left: &traversal.TreeNode{
            Val: 3,
            Left: &traversal.TreeNode{Val: 2},
            Right: &traversal.TreeNode{Val: 4},
        },
        Right: &traversal.TreeNode{
            Val: 8,
            Left: &traversal.TreeNode{Val: 7},
            Right: &traversal.TreeNode{Val: 9},
        },
    }
    
    // Inorder traversal of BST gives sorted order
    inorder := traversal.InorderTraversal(bst)
    fmt.Printf("BST Inorder (sorted): %v\n", inorder) // [2 3 4 5 7 8 9]
}
```

### Expression Tree Evaluation

```go
func demonstrateExpressionTree() {
    // Expression tree for: (2 + 3) * (4 - 1)
    //       *
    //      / \
    //     +   -
    //    / \ / \
    //   2  3 4  1
    
    expr := &traversal.TreeNode{
        Val: '*',
        Left: &traversal.TreeNode{
            Val: '+',
            Left: &traversal.TreeNode{Val: 2},
            Right: &traversal.TreeNode{Val: 3},
        },
        Right: &traversal.TreeNode{
            Val: '-',
            Left: &traversal.TreeNode{Val: 4},
            Right: &traversal.TreeNode{Val: 1},
        },
    }
    
    // Postorder gives the correct evaluation order
    postorder := traversal.PostorderTraversal(expr)
    fmt.Printf("Expression evaluation order: %v\n", postorder) // [2 3 + 4 1 - *]
    // This represents: 2 3 + 4 1 - * = (2+3) * (4-1) = 5 * 3 = 15
}
```

## Algorithm Performance

| Algorithm | Time Complexity | Space Complexity | Notes |
|-----------|----------------|------------------|-------|
| Preorder Recursive | O(n) | O(h) | h = height of tree |
| Preorder Iterative | O(n) | O(h) | Uses explicit stack |
| Inorder Recursive | O(n) | O(h) | Best for BST |
| Inorder Iterative | O(n) | O(h) | Uses explicit stack |
| Inorder Morris | O(n) | O(1) | No extra space needed |
| Postorder Recursive | O(n) | O(h) | Best for node deletion |
| Postorder Iterative | O(n) | O(h) | Two-stack approach |
| Postorder Single Stack | O(n) | O(h) | More complex logic |

## Real-World Applications

### Preorder Traversal
- **File System Traversal**: Visit directory before its contents
- **Expression Tree Prefix Notation**: Convert to prefix notation
- **Tree Serialization**: Serialize tree structure
- **Dependency Resolution**: Process dependencies before dependents

### Inorder Traversal  
- **Binary Search Tree Operations**: Get sorted data
- **Expression Tree Infix Notation**: Convert to infix notation with parentheses
- **Memory Management**: In-order memory allocation
- **Database Index Traversal**: Retrieve data in sorted order

### Postorder Traversal
- **File System Cleanup**: Delete files before directories
- **Expression Tree Evaluation**: Evaluate expressions (postfix)
- **Memory Deallocation**: Free children before parent
- **Directory Size Calculation**: Calculate sizes bottom-up

## Test Coverage

The test suites cover:

- ✅ Empty trees (nil root)
- ✅ Single node trees
- ✅ Complete binary trees
- ✅ Left/right skewed trees
- ✅ Binary search trees
- ✅ Trees with negative values
- ✅ Trees with duplicate values
- ✅ Expression trees
- ✅ Consistency between recursive/iterative implementations
- ✅ Performance benchmarks
- ✅ Special properties (BST sorting, expression evaluation)

## Running Specific Tests

```bash
# Test empty tree handling
go test -v -run "Empty"

# Test BST properties
go test -v -run "BST"

# Test consistency between implementations
go test -v -run "Consistency"

# Test expression tree functionality
go test -v -run "Expression"

# Run performance benchmarks
go test -bench=BenchmarkPreorder
go test -bench=BenchmarkInorder  
go test -bench=BenchmarkPostorder
```

## Contributing

When adding new features or tests:

1. Follow the existing code style and naming conventions
2. Add comprehensive test cases
3. Update documentation and examples
4. Ensure all tests pass: `go test -v`
5. Run benchmarks to check performance: `go test -bench=.`
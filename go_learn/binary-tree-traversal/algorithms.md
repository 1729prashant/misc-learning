# Binary Tree Traversal Algorithms - Knuth Style

## Overview

Binary tree traversal is the process of visiting each node in a binary tree exactly once in a systematic way. The three fundamental traversal methods are:

1. **Preorder Traversal**: Visit root, then left subtree, then right subtree (Root → Left → Right)
2. **Inorder Traversal**: Visit left subtree, then root, then right subtree (Left → Root → Right)
3. **Postorder Traversal**: Visit left subtree, then right subtree, then root (Left → Right → Root)

## Node Structure

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
```

## Algorithm 1: Preorder Traversal

### Knuth-Style Algorithm Description

**Algorithm P** (Preorder traversal of binary tree)

Given a binary tree with root node T, this algorithm visits all nodes in preorder sequence.

**P1.** [Initialize] Set current node to T. If T is nil, terminate.

**P2.** [Visit current node] Process the current node's value.

**P3.** [Traverse left subtree] If current node has a left child, recursively apply this algorithm to the left subtree.

**P4.** [Traverse right subtree] If current node has a right child, recursively apply this algorithm to the right subtree.

**P5.** [Terminate] Return to caller.

### Flowchart

```
START
  ↓
[Current = Root]
  ↓
[Current == nil?] → YES → END
  ↓ NO
[Visit Current Node]
  ↓
[Has Left Child?] → YES → [Recursively traverse Left]
  ↓ NO                      ↓
[Has Right Child?] ← ← ← ← ← ←
  ↓ NO
  ↓ YES
[Recursively traverse Right]
  ↓
END
```

### Complexity Analysis
- **Time Complexity**: O(n) where n is the number of nodes
- **Space Complexity**: O(h) where h is the height of the tree (due to recursion stack)

## Algorithm 2: Inorder Traversal

### Knuth-Style Algorithm Description

**Algorithm I** (Inorder traversal of binary tree)

Given a binary tree with root node T, this algorithm visits all nodes in inorder sequence.

**I1.** [Initialize] Set current node to T. If T is nil, terminate.

**I2.** [Traverse left subtree] If current node has a left child, recursively apply this algorithm to the left subtree.

**I3.** [Visit current node] Process the current node's value.

**I4.** [Traverse right subtree] If current node has a right child, recursively apply this algorithm to the right subtree.

**I5.** [Terminate] Return to caller.

### Flowchart

```
START
  ↓
[Current = Root]
  ↓
[Current == nil?] → YES → END
  ↓ NO
[Has Left Child?] → YES → [Recursively traverse Left]
  ↓ NO                      ↓
[Visit Current Node] ← ← ← ← ←
  ↓
[Has Right Child?] → YES → [Recursively traverse Right]
  ↓ NO                      ↓
END ← ← ← ← ← ← ← ← ← ← ← ←
```

### Complexity Analysis
- **Time Complexity**: O(n) where n is the number of nodes
- **Space Complexity**: O(h) where h is the height of the tree

### Special Property
For Binary Search Trees (BST), inorder traversal visits nodes in ascending sorted order.

## Algorithm 3: Postorder Traversal

### Knuth-Style Algorithm Description

**Algorithm O** (Postorder traversal of binary tree)

Given a binary tree with root node T, this algorithm visits all nodes in postorder sequence.

**O1.** [Initialize] Set current node to T. If T is nil, terminate.

**O2.** [Traverse left subtree] If current node has a left child, recursively apply this algorithm to the left subtree.

**O3.** [Traverse right subtree] If current node has a right child, recursively apply this algorithm to the right subtree.

**O4.** [Visit current node] Process the current node's value.

**O5.** [Terminate] Return to caller.

### Flowchart

```
START
  ↓
[Current = Root]
  ↓
[Current == nil?] → YES → END
  ↓ NO
[Has Left Child?] → YES → [Recursively traverse Left]
  ↓ NO                      ↓
[Has Right Child?] ← ← ← ← ←
  ↓ NO
  ↓ YES
[Recursively traverse Right]
  ↓
[Visit Current Node]
  ↓
END
```

### Complexity Analysis
- **Time Complexity**: O(n) where n is the number of nodes
- **Space Complexity**: O(h) where h is the height of the tree

### Special Property
Postorder traversal is particularly useful for:
- Deleting nodes (children are processed before parent)
- Calculating directory sizes in file systems
- Expression tree evaluation

## Example Tree Traversals

Consider the following binary tree:

```
        1
       / \
      2   3
     / \
    4   5
```

**Preorder**: 1 → 2 → 4 → 5 → 3
**Inorder**: 4 → 2 → 5 → 1 → 3
**Postorder**: 4 → 5 → 2 → 3 → 1

## Implementation Notes

1. All implementations use recursive approach for simplicity and clarity
2. Each traversal method returns a slice of integers representing the visit order
3. Nil nodes are handled gracefully without causing panics
4. Unit tests cover various tree structures including empty trees, single nodes, and unbalanced trees
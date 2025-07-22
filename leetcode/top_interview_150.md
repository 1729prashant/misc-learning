# Leetcode Top Interview 150
-
> **Compute Maximum Depth of a Binary Tree**

```
Algorithm D (Compute Maximum Depth of a Binary Tree)
Input: A binary tree with root node n (possibly null).
Output: An integer representing the maximum depth of the tree.

D1. [Check null]
    If n is null, return 0.

D2. [Recurse]
    Set L ← result of Algorithm D called with n.left.
    Set R ← result of Algorithm D called with n.right.

D3. [Compute max]
    Return 1 + max(L, R).

Complexity:
    Time:  O(n), where n is the number of nodes in the tree.
    Space: O(h), where h is the height of the tree.
           → Worst-case: O(n) if the tree is skewed.
           → Best-case:  O(log n) if the tree is balanced.
```


-


> **Check If Two Binary Trees Are Identical**

```
Algorithm T (Check If Two Binary Trees Are Identical)
Input: Two binary tree nodes p and q (possibly null).
Output: A boolean indicating whether the trees rooted at p and q are structurally identical and contain the same values.

T1. [Both null]
    If p = null and q = null, return true.

T2. [Only one is null]
    If p = null or q = null, return false.

T3. [Check value]
    If p.Val ≠ q.Val, return false.

T4. [Recursive check]
    Return the logical AND of:
        - result of Algorithm T on p.Left and q.Left
        - result of Algorithm T on p.Right and q.Right

Complexity:
    Time:  O(n), where n is the number of nodes in the smaller of the two trees.
    Space: O(h), where h is the height of the tree, due to recursion stack.
```

-


> **Invert a Binary Tree**

```
Algorithm I (Invert a Binary Tree)
Input: A binary tree with root node n (possibly null).
Output: The root of the inverted binary tree.

I1. [Check null]
    If n is null, return null.

I2. [Swap children]
    Swap n.left and n.right.

I3. [Recursive inversion]
    Call Algorithm I on n.left.
    Call Algorithm I on n.right.

I4. [Return result]
    Return n.
    
Complexity:  
    Time:  O(n), where n is the number of nodes — each node is visited once.  
    Space: O(h), where h is the height of the tree — recursion stack can go as deep as the height (worst case O(n) for skewed trees, O(log n) for balanced trees).

```

-

> **Check If a Binary Tree Is Symmetric**

```
Algorithm S (Check If a Binary Tree Is Symmetric)
Input: A binary tree with root node n (possibly null).
Output: A boolean indicating whether the tree is symmetric about its center.

S1. [Call mirror helper]
    Return result of Algorithm M called with (n, n).


Algorithm M (Check If Two Subtrees Are Mirrors)
Input: Two binary tree nodes t1 and t2 (possibly null).
Output: A boolean indicating whether the two trees are mirror images of each other.

M1. [Both null]
    If t1 = null and t2 = null, return true.

M2. [Only one null]
    If t1 = null or t2 = null, return false.

M3. [Value mismatch]
    If t1.Val ≠ t2.Val, return false.

M4. [Recurse on mirror positions]
    Return the logical AND of:
        - result of Algorithm M on (t1.Left, t2.Right)
        - result of Algorithm M on (t1.Right, t2.Left)

Complexity:
    Time: O(n) — where n is the total number of nodes in the tree. Each node is visited once.
    Space: O(h) — where h is the height of the tree due to the recursion stack (best: O(log n), worst: O(n) for skewed trees).
               
```



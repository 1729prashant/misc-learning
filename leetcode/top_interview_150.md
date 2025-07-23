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


-

> **Build Binary Tree From Preorder and Inorder Traversal V1**

```
Algorithm B (Build Binary Tree from Preorder and Inorder Traversals)
Input: 
    - Preorder: a list of integers representing preorder traversal of the tree.
    - Inorder: a list of integers representing inorder traversal of the same tree.
Output: 
    - The root node of the reconstructed binary tree.

B1. [Build index map]
    Construct a map M where each value in Inorder maps to its index.
    Initialize PreIndex ← 0.

B2. [Define recursive helper]
    Define recursive function Helper(InStart, InEnd):
        If InStart > InEnd, return nil.

        [Select root]
        Set RootVal ← Preorder[PreIndex], increment PreIndex by 1.

        Create new node Root with value RootVal.

        [Divide inorder]
        Let InIndex ← M[RootVal]

        [Recurse left]
        Set Root.Left ← Helper(InStart, InIndex - 1)

        [Recurse right]
        Set Root.Right ← Helper(InIndex + 1, InEnd)

        Return Root.

B3. [Return result]
    Call Helper(0, len(Inorder) - 1) and return the result.

Complexity:
    Time: O(n), where n is the number of nodes:
			Each node is visited once.
			The inMap lookup is constant-time due to the hashmap.

    Space O(n): For the map inMap: stores n entries.
                For the recursion stack: in the worst case (skewed tree), the stack is O(n) deep.
          Space can be considered O(h) for recursion (height of the tree), but worst case h = n.
```


-


> **Build Binary Tree From Preorder and Inorder Traversal V2**

```
Algorithm B (Build Binary Tree)
Input: 
    preorder — array of length n representing the preorder traversal of a binary tree.
    inorder  — array of length n representing the inorder traversal of the same tree.
Output: 
    A pointer to the root node of the reconstructed binary tree.

B1. [Create value-index map]
    Construct a map valueToIndex such that:
        For i = 0 to n−1:
            valueToIndex[inorder[i]] ← i

B2. [Initialize preorder index]
    Let preIndex ← 0

B3. [Build the tree]
    Return result of Algorithm C called with (inLeft ← 0, inRight ← n−1)


Algorithm C (Construct Subtree from Inorder Bounds)
Input:
    inLeft, inRight — integers defining current subtree bounds in inorder traversal.
Output:
    A pointer to the root of the constructed subtree.

C1. [Base case: empty subtree]
    If inLeft > inRight, return null.

C2. [Extract root from preorder]
    Let rootVal ← preorder[preIndex]
    Increment preIndex ← preIndex + 1
    Let root ← new TreeNode(rootVal)

C3. [Find split index in inorder]
    Let inRootIndex ← valueToIndex[rootVal]

C4. [Build left subtree]
    root.Left ← result of Algorithm C on (inLeft, inRootIndex − 1)

C5. [Build right subtree]
    root.Right ← result of Algorithm C on (inRootIndex + 1, inRight)

C6. [Return constructed root]
    Return root

Complexity:
    Time: O(n) — Each node is processed once, and map lookups are O(1).
    Space: O(n) — For the value-to-index map and recursion stack (worst case: skewed tree depth n).
```

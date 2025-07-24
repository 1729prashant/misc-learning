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


-



> **Construct Binary Tree from Inorder and Postorder Traversals**

```
Algorithm B (Build Tree from Inorder and Postorder)
Input: 
    - inorder[0..n−1]: an array representing inorder traversal of a binary tree
    - postorder[0..n−1]: an array representing postorder traversal of the same tree
Output: 
    - Pointer to the root of the constructed binary tree

B1. [Preprocessing]
    Create a hashmap indexMap to store the index of each value in inorder array:
        For i from 0 to n−1:
            indexMap[inorder[i]] ← i

B2. [Initialize global postIndex]
    Set postIndex ← n − 1   // last element in postorder is the root

B3. [Define recursive builder]
    Define recursive procedure BuildSubtree(inLeft, inRight):
        Input: current inorder boundaries (inLeft to inRight)
        Output: root node of the constructed subtree

        If inLeft > inRight:
            Return null   // no nodes in this subtree

        rootVal ← postorder[postIndex]
        postIndex ← postIndex − 1

        Create new TreeNode root with value rootVal

        mid ← indexMap[rootVal]  // find root index in inorder

        root.Right ← BuildSubtree(mid + 1, inRight)   // build right subtree first
        root.Left  ← BuildSubtree(inLeft, mid − 1)    // then build left subtree

        Return root

B4. [Build tree]
    Return BuildSubtree(0, n − 1)


Complexity:
    Time: O(n) — each node is processed once; hashmap gives O(1) lookups
    Space: O(n) — for hashmap and recursion stack in worst case (unbalanced tree)

```

-

> **Connect Next Right Pointers in Each Node (Constant Space)**

```
Algorithm C (Connect Level Order Next Pointers)
Input: 
    A binary tree with root node r (possibly null). Each node has: val, left, right, next
Output:
    The same tree where each node’s next pointer is set to its next right node or null

C1. [Initialize level]
    Let level ← r   // leftmost node at current level

C2. [Loop through levels]
    While level ≠ null:
        Let curr ← level      // current node in level
        Let prev ← null       // previous node to connect via next
        Let nextLevel ← null  // leftmost node of next level

        C2.1. [Loop through nodes at current level]
            While curr ≠ null:
                For each child of curr (first left, then right):
                    If child ≠ null:
                        If prev ≠ null:
                            prev.next ← child
                        Else:
                            nextLevel ← child
                        prev ← child
                curr ← curr.next

        C2.2. [Move to next level]
            level ← nextLevel

C3. [Done]
    Return r


Complexity:
    Time: O(n) — each node is visited exactly once
    Space: O(1) — no extra space used except a few pointers (constant space)

```

-

> **Flatten Binary Tree to Linked List**

```
Algorithm F (Flatten Binary Tree to Linked List)
Input: A binary tree with root node n (possibly null).
Output: The same tree, modified in-place to become a right-skewed pre-order linked list.

F1. [Base case]
    If n is null, return.

F2. [Recurse on left and right]
    Flatten the left subtree of n using Algorithm F.
    Flatten the right subtree of n using Algorithm F.

F3. [Store flattened left and right]
    Let left = n.Left
    Let right = n.Right

F4. [Relink]
    Set n.Left = null
    Set n.Right = left

F5. [Find tail of new right subtree]
    Traverse from n along the right child until you find the last node.

F6. [Attach original right subtree]
    Set tail.Right = right

F7. [Done]
    Return (modifications are in-place)

Complexity:
    Time: O(n) — each node is visited once.
    Space: O(h) — due to recursion stack, where h is tree height.
               Worst: O(n) (skewed tree), Best: O(log n) (balanced tree)

```

-


> **Flatten Binary Tree to Linked List O(1)**

```
Algorithm F' (Flatten Binary Tree to Linked List In-Place)
Input: A binary tree with root node n (possibly null).
Output: The same tree, modified in-place to become a right-skewed pre-order linked list.

F'1. [Initialize]
     Let curr ← n

F'2. [Traverse]
     While curr ≠ null, do:

     F'2.1. [Check for left child]
            If curr.Left ≠ null:

            F'2.1.1. [Find predecessor]
                     Let prev ← curr.Left
                     While prev.Right ≠ null:
                         prev ← prev.Right

            F'2.1.2. [Rewire]
                     prev.Right ← curr.Right
                     curr.Right ← curr.Left
                     curr.Left ← null

     F'2.2. [Move to next node]
            curr ← curr.Right

F'3. [Done]
     Return (modifications are in-place)

Complexity:
    Time: O(n) — each node is visited a constant number of times.
    Space: O(1) — no recursion or stack used; only constant pointers.

```
> **Has Path Sum**

```
Algorithm H (Has Path Sum)
Input: A binary tree rooted at node n, and an integer targetSum
Output: true if there exists a root-to-leaf path such that the sum of values equals targetSum; else false

H1. [Base case – empty tree]
    If n = null:
        Return false

H2. [Leaf check]
    If n.Left = null and n.Right = null:
        If targetSum = n.Val:
            Return true
        Else:
            Return false

H3. [Recursive case]
    Return H(n.Left, targetSum - n.Val) OR H(n.Right, targetSum - n.Val)

Time Complexity: O(n) — where n is the number of nodes (must check all paths in worst case).
Space Complexity: O(h) — where h is the height of the tree (stack space due to recursion).

```

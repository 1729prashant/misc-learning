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

> **Sum Root to Leaf Numbers**

```
Algorithm SumRootToLeaf(root)
Input: A binary tree where each node contains a digit (0–9)
Output: The sum of all numbers formed from root-to-leaf paths

1. Define Recursive Function dfs(node, currentNumber)
     a. If node is null:
         i. return 0
     b. currentNumber ← currentNumber × 10 + node.val
     c. If node is a leaf (node.left = null and node.right = null):
         i. return currentNumber
     d. Return dfs(node.left, currentNumber) + dfs(node.right, currentNumber)

2. Return dfs(root, 0)

```

> **Maximum Path Sum in Binary Tree**

```
Algorithm M (Find the maximum path sum in a binary tree)
Input:
    - Root: a binary tree node representing the root of the tree
Output:
    - An integer representing the maximum path sum of any path in the tree

M1. [Initialize global maximum]
    Set MaxSum ← -∞

M2. [Define recursive helper]
    Define function MaxGain(Node):
        If Node is nil, return 0

        [Left subtree]
        Let LeftGain ← max(MaxGain(Node.Left), 0)

        [Right subtree]
        Let RightGain ← max(MaxGain(Node.Right), 0)

        [Current path sum with split at Node]
        Let PriceNewPath ← Node.Val + LeftGain + RightGain

        Update MaxSum ← max(MaxSum, PriceNewPath)

        [Return max gain continuing upward]
        Return Node.Val + max(LeftGain, RightGain)

M3. [Invoke recursion and return result]
    Call MaxGain(Root)
    Return MaxSum

Complexity:
    Time: O(n), where n is the number of nodes.
        Each node is visited once.
    Space: O(h), where h is the height of the tree (due to recursion stack).
        In worst case (skewed tree), h = n ⇒ O(n) space.

```

-

> **Binary Search Tree Iterator**

```
Algorithm I (BST Iterator using stack-based in-order traversal)
Input:
    - Root: pointer to the root of a binary search tree
Output:
    - An iterator object with Next and HasNext methods

I1. [Initialize stack]
    Let Stack ← empty list of nodes

I2. [Define helper pushLeft]
    Function pushLeft(Node):
        While Node ≠ nil:
            Push Node onto Stack
            Node ← Node.Left

I3. [Constructor]
    Call pushLeft(Root)

I4. [Next operation]
    Let Node ← Stack.Pop()
    If Node.Right ≠ nil:
        Call pushLeft(Node.Right)
    Return Node.Val

I5. [HasNext operation]
    Return whether Stack is non-empty

Complexity:
    Time:
        - Next and HasNext run in O(1) average time
        - Because each node is pushed and popped exactly once
    Space:
        - O(h), where h is the height of the tree (stack space)

```

-



> **Count nodes in a complete binary tree**

```
Algorithm C (Count nodes in a complete binary tree)
Input:
    - Root: pointer to the root of a complete binary tree
Output:
    - Integer count of total nodes in the tree

C1. [Handle empty tree]
    If Root is nil, return 0

C2. [Compute heights]
    Set LeftHeight ← ComputeHeight(Root.Left, goLeft=true)
    Set RightHeight ← ComputeHeight(Root.Right, goLeft=false)

C3. [Check if tree is perfect]
    If LeftHeight = RightHeight:
        // Tree is full ⇒ number of nodes is 2^(h+1) - 1
        Return (1 << (LeftHeight + 1)) - 1

C4. [Recurse for incomplete side]
    Return 1 + C(Root.Left) + C(Root.Right)


Helper Algorithm ComputeHeight(Node, goLeft)
Input: 
    - Node: a tree node
    - goLeft: boolean (true = follow leftmost path, false = rightmost)
Output:
    - Height of the path from Node to the bottom

H1. Set height ← 0
H2. While Node ≠ nil:
        height ← height + 1
        If goLeft: Node ← Node.Left
        Else:      Node ← Node.Right
H3. Return height

Complexity:
    Time: O(log² n)
        Each height computation takes O(log n), and recursion depth is O(log n)
    Space: O(log n) due to recursion stack

```


-


> **Find the Lowest Common Ancestor of nodes p and q in a binary tree**

```
Algorithm L (Find the Lowest Common Ancestor of nodes p and q in a binary tree)
Input:
    - root: The root of the binary tree
    - p, q: The two nodes whose lowest common ancestor we must find
Output:
    - The lowest common ancestor node of p and q

L.1: [Base case]
    If root is nil or root is p or root is q:
        Return root

L.2: [Search left subtree]
    Let left ← LCA(root.Left, p, q)

L.3: [Search right subtree]
    Let right ← LCA(root.Right, p, q)

L.4: [Check whether p and q found in both sides]
    If left ≠ nil and right ≠ nil:
        Return root   // This node is the lowest common ancestor

L.5: [Return non-nil side or nil if both nil]
    If left ≠ nil:
        Return left
    Else:
        Return right

Complexity:
    Time: O(n) – visits each node once
    Space: O(h) – where h is the height of the tree (due to recursion stack)

```

-

> **Right Side View of Binary Tree**

```
Algorithm R (Right side view of a binary tree using level-order traversal)
Input: 
    - root: pointer to the root node of a binary tree.
Output: 
    - A list of integers representing the visible nodes from the right side, top to bottom.

R.1: [Handle base case]
    If root = nil:
        Return []

R.2: [Initialize]
    Initialize an empty result list → result ← []
    Initialize a queue Q and enqueue root → Q ← [root]

R.3: [Perform level-order traversal]
    While Q is not empty:
        Let n ← length(Q)

        For i ← 0 to n - 1:
            node ← Q.pop_front()

            If i = n - 1:
                Append node.Val to result

            If node.Left ≠ nil:
                Append node.Left to Q

            If node.Right ≠ nil:
                Append node.Right to Q

R.4: [Return final result]
    Return result

Complexity:
    Time: O(n) — Every node is visited exactly once.
    Space: O(n) — In the worst case, the queue contains all nodes at the deepest level.
```


-


> **Average of Levels in Binary Tree**

```
Algorithm A (Average of node values at each level)
Input:
    - root: a pointer to the root node of a binary tree.
Output:
    - A list of float values, where the i-th element is the average of values at level i.

A1. [Initialize]
    If root is nil, return empty list.
    Create an empty list Result.
    Initialize a queue Q and enqueue root.

A2. [Process each level]
    While Q is not empty:
        Let Size ← length of Q.
        Let Sum ← 0.

        For i from 1 to Size:
            Dequeue a node N from Q.
            Add N.Val to Sum.

            If N.Left is not nil, enqueue N.Left.
            If N.Right is not nil, enqueue N.Right.

        Append (Sum / Size) to Result.

A3. [Return]
    Return Result.

Complexity:
    Time: O(n), where n is the total number of nodes. Each node is visited once.
    Space: O(w), where w is the maximum width of the tree (max number of nodes at any level).

```

> **Binary Tree Level Order Traversal**

```
Algorithm L (Level-order traversal of binary tree)  
Input:  
    - Root: pointer to the root of a binary tree  
Output:  
    - A list of lists of integers, where each sublist contains node values at a given depth level from left to right  

L1. [Handle empty tree]  
    If Root is nil, return an empty list []  

L2. [Initialize queue and result list]  
    Let Queue be an empty queue  
    Enqueue Root to Queue  
    Let Result be an empty list  

L3. [Traverse levels until queue is empty]  
    While Queue is not empty:  
        Let LevelSize ← length of Queue  
        Let Level ← empty list  
        Repeat LevelSize times:  
            Dequeue Node ← Queue.pop()  
            Append Node.Val to Level  
            If Node.Left ≠ nil, enqueue Node.Left  
            If Node.Right ≠ nil, enqueue Node.Right  
        Append Level to Result  

L4. [Return result]  
    Return Result  

Complexity:  
    Time: O(n) — Each node is visited exactly once  
    Space: O(n) — Queue may store up to n/2 nodes at the last level in the worst case  
    
```

-


> **Binary Tree Zigzag Level Order Traversal**

```
Algorithm Z (Zigzag Level Order Traversal of Binary Tree)
Input:
    - Root: pointer to the root of a binary tree
Output:
    - A list of lists of integers representing zigzag level order traversal

Z1. [Handle empty tree]
    If Root is nil, return an empty list

Z2. [Initialize structures]
    Let Result ← empty list
    Let Queue ← empty queue
    Enqueue Root into Queue
    Let LeftToRight ← true   // traversal direction flag

Z3. [Level-order traversal with direction toggle]
    While Queue is not empty:
        Let LevelSize ← length of Queue
        Let Level ← empty list of size LevelSize

        For i from 0 to LevelSize - 1:
            Dequeue Node from Queue

            If LeftToRight:
                Set Level[i] ← Node.Val
            Else:
                Set Level[LevelSize - 1 - i] ← Node.Val

            If Node.Left ≠ nil: Enqueue Node.Left
            If Node.Right ≠ nil: Enqueue Node.Right

        Append Level to Result
        Set LeftToRight ← not LeftToRight

Z4. Return Result

Complexity:
    Time: O(n), where n is the number of nodes in the tree
        - Each node is visited once and placed in a level
    Space: O(n) for the queue and result list

```


-


> **Minimum Absolute Difference in BST**

```
Algorithm D (Minimum Absolute Difference in BST)
Input:
    - Root: pointer to the root of a Binary Search Tree
Output:
    - Integer representing the minimum absolute difference between values of any two nodes

D1. [Initialize state]
    Set prev ← nil
    Set minDiff ← ∞

D2. [In-order traversal]
    Call InOrder(Root)

D3. [Return result]
    Return minDiff


Helper Algorithm InOrder(Node)
Input:
    - Node: a tree node in the BST
Uses shared variables:
    - prev: pointer to previously visited node
    - minDiff: global minimum difference

I1. [Base case]
    If Node is nil, return

I2. [Left subtree]
    Call InOrder(Node.Left)

I3. [Visit current node]
    If prev ≠ nil:
        diff ← |Node.Val - prev.Val|
        minDiff ← min(minDiff, diff)
    Set prev ← Node

I4. [Right subtree]
    Call InOrder(Node.Right)

Complexity:
    Time: O(n) — each node visited once in in-order traversal
    Space: O(h) — recursion stack depth (O(log n) for balanced BST, O(n) worst case)

```


-


> **Kth Smallest Element in a BST**

```
Algorithm E (Kth Smallest Element in BST)
Input:
    - Root: pointer to the root of a Binary Search Tree
    - k: integer representing the 1-based index of the smallest element to find
Output:
    - The kth smallest value in the BST

E1. [Initialize state]
    Set count ← 0
    Set result ← nil

E2. [In-order traversal to find kth smallest]
    Call InOrder(Root)

E3. [Return result]
    Return result


Helper Algorithm InOrder(Node)
Input:
    - Node: a tree node in the BST
Uses shared variables:
    - count: tracks how many nodes have been visited
    - result: stores the kth smallest value

I1. [Base case]
    If Node is nil or result is already set, return

I2. [Left subtree]
    Call InOrder(Node.Left)

I3. [Visit current node]
    count ← count + 1
    If count = k:
        result ← Node.Val
        Return

I4. [Right subtree]
    Call InOrder(Node.Right)

Complexity:
    Time: O(h + k) — up to k nodes may be visited along a path of height h (O(log n) in balanced, O(n) in worst case)
    Space: O(h) — recursion stack (O(log n) in balanced BST)

Follow-up Optimization:
    - Augment the BST nodes with a `size` field: the number of nodes in the subtree rooted at each node.
    - This allows finding the kth smallest in O(log n) time:
        - At each node, compare k with size of left subtree:
            - If k ≤ size(left), go left
            - If k = size(left) + 1, return current node
            - Else go right with adjusted k: k - size(left) - 1
    - This requires updating the size field on every insert/delete (O(log n) maintenance cost).

```


-

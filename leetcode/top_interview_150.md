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


> **Validate Binary Search Tree**

```
Algorithm V (Validate if a binary tree is a BST)
Input:
    - Root: pointer to the root of a binary tree
Output:
    - Boolean indicating whether the tree is a valid BST

V1. [Start validation with full range]
    Return V-Helper(Root, -∞, +∞)

Helper Algorithm V-Helper(Node, Min, Max)
Input:
    - Node: current tree node
    - Min: minimum allowed value (exclusive)
    - Max: maximum allowed value (exclusive)
Output:
    - Boolean indicating whether subtree rooted at Node is valid

H1. [Base case: empty node]
    If Node = nil, return true

H2. [Check validity of current node]
    If Node.Val ≤ Min OR Node.Val ≥ Max, return false

H3. [Recurse on left and right subtree]
    Return V-Helper(Node.Left, Min, Node.Val) AND
           V-Helper(Node.Right, Node.Val, Max)

Complexity:
    Time: O(n), where n is the number of nodes (each node is visited once)
    Space: O(h), where h is the height of the tree (recursion stack)

```


-


> **Linked List Cycle**

```
Algorithm D (Detect cycle in a singly linked list)
Input:
    - Head: pointer to the head of the linked list
Output:
    - Boolean indicating whether the list contains a cycle

D1. [Initialize two pointers]
    Set slow ← Head
    Set fast ← Head

D2. [Traverse list with two speeds]
    While fast ≠ nil AND fast.Next ≠ nil:
        slow ← slow.Next
        fast ← fast.Next.Next

        If slow = fast:
            Return true    // cycle detected

D3. [No cycle found]
    Return false

Complexity:
    Time: O(n), where n is the number of nodes in the list
    Space: O(1), uses only two pointers

Explanation:
    This is Floyd’s Tortoise and Hare algorithm.
    - `slow` moves one step at a time.
    - `fast` moves two steps at a time.
    - If there's a cycle, they will eventually meet inside the loop.
    - If there's no cycle, `fast` will reach the end of the list (nil).

```


-


>  **Add Two Numbers**

```
Algorithm A (Add two reversed linked lists representing integers)
Input:
    - l1: ListNode pointer to the head of the first linked list (least significant digit first)
    - l2: ListNode pointer to the head of the second linked list (least significant digit first)
Output:
    - A ListNode pointer to the head of a new linked list representing the sum

A1. [Initialize dummy head and carry]
    Let dummy ← new ListNode(0)
    Let curr ← dummy
    Let carry ← 0

A2. [Traverse both lists and compute digit sums]
    While l1 ≠ nil ∨ l2 ≠ nil ∨ carry ≠ 0, do:
        Let val1 ← if l1 ≠ nil then l1.Val else 0
        Let val2 ← if l2 ≠ nil then l2.Val else 0

        Let sum ← val1 + val2 + carry
        Let carry ← sum div 10
        Let digit ← sum mod 10

        curr.Next ← new ListNode(digit)
        curr ← curr.Next

        If l1 ≠ nil, then l1 ← l1.Next  
        If l2 ≠ nil, then l2 ← l2.Next

A3. [Return result skipping dummy head]
    Return dummy.Next

Complexity:
    Time: O(max(m, n)), where m and n are the lengths of the input lists  
    Space: O(max(m, n)), to store the output list (one node per digit)

```


-


> **Merge Two Sorted Lists**

```
Algorithm M (Merge two sorted linked lists)
Input:
    - list1: ListNode pointer to the head of the first sorted linked list
    - list2: ListNode pointer to the head of the second sorted linked list
Output:
    - A ListNode pointer to the head of the merged sorted linked list

M1. [Initialize dummy node and tail pointer]
    Let dummy ← new ListNode(0)
    Let tail ← dummy

M2. [Traverse both lists and pick the smaller node each time]
    While list1 ≠ nil ∧ list2 ≠ nil, do:
        If list1.Val ≤ list2.Val, then:
            tail.Next ← list1
            list1 ← list1.Next
        Else:
            tail.Next ← list2
            list2 ← list2.Next
        tail ← tail.Next

M3. [Append remaining nodes]
    If list1 ≠ nil, then:
        tail.Next ← list1
    Else if list2 ≠ nil, then:
        tail.Next ← list2

M4. [Return the merged list skipping dummy head]
    Return dummy.Next

Complexity:
    Time: O(m + n), where m and n are the lengths of the two input lists  
    Space: O(1), since merging is done in-place using existing nodes

```


-


> **Copy List with Random Pointer**
```
Algorithm D (Deep copy a linked list with next and random pointers)
Input:
    - Head: a pointer to the head of a singly linked list where each node has:
        - Val: an integer
        - Next: pointer to the next node
        - Random: pointer to any node or null

Output:
    - A new head pointer to the deep-copied linked list with identical structure

D1. [Interleave copied nodes into the original list]
    Set Curr ← Head
    While Curr ≠ null:
        Let NewNode ← new Node(Curr.Val)
        Set NewNode.Next ← Curr.Next
        Set Curr.Next ← NewNode
        Set Curr ← NewNode.Next

D2. [Assign random pointers for the copied nodes]
    Set Curr ← Head
    While Curr ≠ null:
        If Curr.Random ≠ null:
            Set Curr.Next.Random ← Curr.Random.Next
        Set Curr ← Curr.Next.Next

D3. [Separate the copied list from the original list]
    Set Curr ← Head
    Set PseudoHead ← new Node(0)
    Set CopyCurr ← PseudoHead
    While Curr ≠ null:
        Set NextOriginal ← Curr.Next.Next

        Set Copy ← Curr.Next
        Set CopyCurr.Next ← Copy
        Set CopyCurr ← Copy

        Set Curr.Next ← NextOriginal
        Set Curr ← NextOriginal

D4. [Return the head of the copied list]
    Return PseudoHead.Next

Complexity:
    Time: O(n)
        - Each node is visited a constant number of times
    Space: O(1)
        - No extra data structures used (copy is done in-place)
```


-


> **Reverse Linked List II**

```
Algorithm R (Reverse sublist from position left to right in a singly linked list)
Input:
    - head: the head of a singly linked list
    - left: integer position to start reversing (1-based)
    - right: integer position to stop reversing (inclusive, right ≥ left)
Output:
    - The head of the modified list with sublist [left, right] reversed

R1. [Handle edge cases]
    If head is nil or left = right:
        Return head

R2. [Initialize dummy node]
    Create dummy node with dummy.Next ← head
    Set prev ← dummy

R3. [Move prev to node before reversal starts]
    For i from 1 to left - 1:
        Set prev ← prev.Next

R4. [Start reversal]
    Let start ← prev.Next
    Let then ← start.Next

R5. [Iteratively reverse the sublist]
    For i from 0 to (right - left - 1):
        Set start.Next ← then.Next
        Set then.Next ← prev.Next
        Set prev.Next ← then
        Set then ← start.Next

R6. [Return the new head]
    Return dummy.Next

Complexity:
    Time: O(n), where n is the number of nodes in the list.
        We traverse the list at most twice: once to reach `left`, and once to reverse up to `right`.
    Space: O(1), in-place reversal without extra memory allocation.

```


-



> **Reverse Nodes in k-Group**

```
Algorithm K (Reverse nodes of the list in groups of size k)
Input:
    - head: the head of a singly linked list
    - k: integer size of groups to reverse
Output:
    - The head of the modified list with every k-sized group reversed

K1. [Handle edge cases]
    If head is nil or k = 1:
        Return head

K2. [Initialize dummy and pointers]
    Create dummy node with dummy.Next ← head
    Set prevGroupEnd ← dummy

K3. [Loop over the list in groups of k]
    While True:
        Set kth ← KHelper1(prevGroupEnd, k)
        If kth is nil:
            Break
        Set groupStart ← prevGroupEnd.Next
        Set nextGroupStart ← kth.Next

K4. [Reverse current group]
        Set prev ← nextGroupStart
        Set curr ← groupStart
        While curr ≠ nextGroupStart:
            Set temp ← curr.Next
            Set curr.Next ← prev
            Set prev ← curr
            Set curr ← temp

K5. [Reconnect reversed group]
        Set prevGroupEnd.Next ← kth
        Set prevGroupEnd ← groupStart

K6. [Return new head]
    Return dummy.Next

Helper Algorithm KHelper1 (Find the k-th node from current)
Input:
    - curr: current node
    - k: number of steps to move forward
Output:
    - Pointer to the k-th node, or nil if fewer than k nodes remain

KHelper1.1:
    While curr ≠ nil and k > 0:
        Set curr ← curr.Next
        Set k ← k - 1
    Return curr

Complexity:
    Time: O(n), where n is the number of nodes in the list.
        Each node is visited a constant number of times (once in group traversal, once in reversal).
    Space: O(1), in-place reversal without extra memory allocation.

```


-


> **Remove Nth Node From End of List**


```
Algorithm RNE (Remove N-th node from end in singly linked list)
Input:
    - head: the head of a singly linked list
    - n: an integer indicating which node from the end to remove (1-based)
Output:
    - The head of the modified list with the n-th node from the end removed

RNE.1: [Initialize dummy node and two pointers]
    Create dummy node with dummy.Next ← head
    Set first ← dummy
    Set second ← dummy

RNE.2: [Advance first pointer by n + 1 steps]
    For i from 0 to n:
        Set first ← first.Next

RNE.3: [Move both pointers until first reaches end]
    While first ≠ nil:
        Set first ← first.Next
        Set second ← second.Next

RNE.4: [Remove target node]
    Set second.Next ← second.Next.Next

RNE.5: [Return updated list]
    Return dummy.Next

Complexity:
    Time: O(n), where n is the length of the list — single traversal.
    Space: O(1), only constant extra pointers used.

```


-



> **Remove Duplicates from Sorted List II**

```
Algorithm D (Delete all duplicates in a sorted singly linked list, keeping only distinct numbers)
Input:
    - head: the head of a sorted singly linked list
Output:
    - The head of the modified list with only unique elements retained

D1. [Handle edge cases]
    If head is nil or head.Next is nil:
        Return head

D2. [Initialize dummy node and pointers]
    Create dummy node with dummy.Next ← head
    Set prev ← dummy
    Set curr ← head

D3. [Traverse the list]
    While curr is not nil:
        If curr.Next is not nil and curr.Val = curr.Next.Val:
            Let duplicateVal ← curr.Val
            While curr is not nil and curr.Val = duplicateVal:
                Set curr ← curr.Next
            Set prev.Next ← curr
        Else:
            Set prev ← curr
            Set curr ← curr.Next

D4. [Return the new head]
    Return dummy.Next

Complexity:
    Time: O(n), where n is the number of nodes in the list.
        We perform a single pass through the list.
    Space: O(1), all operations are done in-place using pointers.

```


-




> **Rotate List**

```
Algorithm R (Rotate a singly linked list to the right by k places)
Input:
    - head: the head of a singly linked list
    - k: number of positions to rotate the list to the right
Output:
    - The head of the rotated linked list

R1. [Handle edge cases]
    If head is nil or head.Next is nil or k = 0:
        Return head

R2. [Compute the length of the list and connect tail to head]
    Set length ← 1
    Set tail ← head
    While tail.Next ≠ nil:
        Set tail ← tail.Next
        Increment length
    Set tail.Next ← head  // Make the list circular

R3. [Compute new tail position after rotation]
    Set k ← k mod length
    Set stepsToNewTail ← length - k

R4. [Find the new tail and new head]
    Set newTail ← head
    For i from 1 to stepsToNewTail - 1:
        Set newTail ← newTail.Next
    Set newHead ← newTail.Next

R5. [Break the circular link]
    Set newTail.Next ← nil

R6. [Return the new head]
    Return newHead

Complexity:
    Time: O(n), where n is the number of nodes in the list.
        We traverse the list twice: once to find the length, once to find the new tail.
    Space: O(1), all operations are done in-place using pointers.

```


-




> **Partition List**

```
Algorithm P (Partition a linked list around a value x preserving relative order)
Input:
    - head: the head of a singly linked list
    - x: integer value to partition around
Output:
    - The head of the modified list where nodes < x come before nodes ≥ x

P1. [Handle edge cases]
    If head is nil or head.Next is nil:
        Return head

P2. [Initialize two dummy lists]
    Create dummyLess and dummyGreater as new nodes
    Set less ← dummyLess
    Set greater ← dummyGreater

P3. [Traverse and partition the original list]
    Set curr ← head
    While curr ≠ nil:
        If curr.Val < x:
            Set less.Next ← curr
            Set less ← less.Next
        Else:
            Set greater.Next ← curr
            Set greater ← greater.Next
        Set curr ← curr.Next

P4. [Connect partitions and terminate the list]
    Set greater.Next ← nil
    Set less.Next ← dummyGreater.Next

P5. [Return new head]
    Return dummyLess.Next

Complexity:
    Time: O(n), where n is the number of nodes in the list.
        We traverse the list once.
    Space: O(1), extra space is used only for constant number of pointers.

```


-




> **LRU Cache**

```
Algorithm L (Design a Least Recently Used Cache with O(1) operations)
Input:
    - capacity: maximum number of key-value pairs the cache can hold
Operations:
    - get(key): return value if key exists, else return -1
    - put(key, value): insert/update key with value and evict LRU if over capacity

Data Structures:
    - A hashmap `cache` mapping key → node for O(1) access
    - A doubly linked list with dummy head and tail to maintain usage order:
        Most recently used at head.Next, least recently used at tail.Prev

L1. [Define node structure]
    Each node stores: key, value, prev, next

L2. [Initialize the cache]
    Create:
        - hashmap `cache` ← empty
        - dummyHead and dummyTail nodes
        - Connect dummyHead.Next ← dummyTail and dummyTail.Prev ← dummyHead
        - Store `capacity` as given

L3. [Define helper: removeNode(node)]
    node.Prev.Next ← node.Next
    node.Next.Prev ← node.Prev

L4. [Define helper: insertToFront(node)]
    node.Next ← dummyHead.Next
    node.Prev ← dummyHead
    dummyHead.Next.Prev ← node
    dummyHead.Next ← node

L5. [Define get(key)]
    If key ∉ cache:
        Return -1
    Else:
        node ← cache[key]
        Call removeNode(node)
        Call insertToFront(node)
        Return node.Value

L6. [Define put(key, value)]
    If key ∈ cache:
        node ← cache[key]
        node.Value ← value
        Call removeNode(node)
        Call insertToFront(node)
    Else:
        If len(cache) = capacity:
            lru ← dummyTail.Prev
            Call removeNode(lru)
            Delete cache[lru.Key]
        Create node with key and value
        Call insertToFront(node)
        Set cache[key] ← node

Complexity:
    Time: O(1) for both get and put due to hashmap and doubly linked list.
    Space: O(capacity), storing up to capacity key-node pairs.

```


-




> **Merge Sorted Array**

```
Algorithm M (Merge two sorted arrays nums1 and nums2 into nums1 in-place)
Input:
    - nums1: integer array of size m + n, with first m elements valid, rest zeros
    - m: number of valid elements in nums1
    - nums2: integer array of size n
    - n: number of elements in nums2
Output:
    - nums1 modified in-place to hold all m + n elements in non-decreasing order

M1. [Initialize pointers]
    Set i ← m - 1      // Last valid element in nums1
    Set j ← n - 1      // Last element in nums2
    Set k ← m + n - 1  // End of nums1 array

M2. [Merge from the back]
    While i ≥ 0 and j ≥ 0:
        If nums1[i] > nums2[j]:
            Set nums1[k] ← nums1[i]
            Decrement i
        Else:
            Set nums1[k] ← nums2[j]
            Decrement j
        Decrement k

M3. [Copy remaining nums2 elements if any]
    While j ≥ 0:
        Set nums1[k] ← nums2[j]
        Decrement j
        Decrement k

M4. [No need to copy remaining nums1 elements]
    Since nums1 elements are already in place

Complexity:
    Time: O(m + n), we traverse both arrays once from the end.
    Space: O(1), done in-place without extra memory.

```


-



> **Remove Element**

```
Algorithm R (Remove all occurrences of a value from an array in-place)
Input:
    - nums: integer array
    - val: integer, value to remove
Output:
    - k: number of elements in nums not equal to val
    - nums modified in-place with first k elements not equal to val

R1. [Initialize write pointer]
    Set k ← 0  // Index to write the next valid element

R2. [Traverse the array]
    For i from 0 to length(nums) - 1:
        If nums[i] ≠ val:
            Set nums[k] ← nums[i]
            Increment k

R3. [Ignore elements beyond k]
    The elements after index k are irrelevant and can be ignored

Complexity:
    Time: O(n), where n is the length of nums — each element is visited once.
    Space: O(1), performed in-place without additional memory.

```


-



> **Remove Duplicates from Sorted Array**

```
Algorithm D (Remove duplicates from sorted array in-place)
Input:
    - nums: integer array sorted in non-decreasing order
Output:
    - k: number of unique elements in nums
    - nums modified in-place so first k elements are unique, in original order

D1. [Handle empty array]
    If length(nums) = 0:
        Return 0

D2. [Initialize write pointer]
    Set k ← 1  // Index for placing next unique element

D3. [Traverse array and copy unique elements]
    For i from 1 to length(nums) - 1:
        If nums[i] ≠ nums[i - 1]:
            Set nums[k] ← nums[i]
            Increment k

D4. [Remaining elements are irrelevant]
    Only first k elements are relevant in nums

Complexity:
    Time: O(n), where n is the length of nums — each element is visited once.
    Space: O(1), performed in-place without additional memory.

```


-



> **Remove Duplicates from Sorted Array II**

```
Algorithm D2 (Remove duplicates from sorted array, allow at most 2 occurrences)
Input:
    - nums: integer array sorted in non-decreasing order
Output:
    - k: number of elements after removing extra duplicates, max 2 per unique value
    - nums modified in-place so first k elements are valid

D2.1: [Handle short arrays]
    If length(nums) ≤ 2:
        Return length(nums)

D2.2: [Initialize write pointer]
    Set k ← 2  // First two elements are always valid

D2.3: [Traverse from the 3rd element]
    For i from 2 to length(nums) - 1:
        If nums[i] ≠ nums[k - 2]:
            Set nums[k] ← nums[i]
            Increment k

D2.4: [Return new length]
    Return k

Complexity:
    Time: O(n), single pass over nums
    Space: O(1), in-place update

```


-


> **Majority Element**

```
Algorithm M (Find majority element using Boyer-Moore Voting)
Input:
    - nums: integer array of size n
Output:
    - Majority element (appears more than ⌊n / 2⌋ times)

M.1: [Initialize candidate and count]
    Set candidate ← nums[0]
    Set count ← 1

M.2: [Iterate through array]
    For i from 1 to length(nums) - 1:
        If count == 0:
            candidate ← nums[i]
            count ← 1
        Else if nums[i] == candidate:
            count ← count + 1
        Else:
            count ← count - 1

M.3: [Return candidate]
    Return candidate

Complexity:
    Time: O(n)
    Space: O(1)

```


-


> **Rotate Array**

```
Algorithm R (Rotate Array Right by k steps using reverse)
Input:
    - nums: integer array of size n
    - k: number of steps to rotate (non-negative)
Output:
    - nums modified in-place with elements rotated right by k steps

R.1: [Normalize k]
    Set k ← k mod n

R.2: [Reverse entire array]
    Reverse(nums, 0, n - 1)

R.3: [Reverse first k elements]
    Reverse(nums, 0, k - 1)

R.4: [Reverse remaining elements]
    Reverse(nums, k, n - 1)

Helper Algorithm Reverse(A, left, right)
    While left < right:
        Swap A[left] and A[right]
        left ← left + 1
        right ← right - 1

Complexity:
    Time: O(n)
    Space: O(1)

```


-


> **Best Time to Buy and Sell Stock**

```
Algorithm M (Maximize Single Buy-Sell Stock Profit)
Input:
    - prices: array of integers representing stock prices on each day
Output:
    - Maximum profit from a single buy-sell transaction (integer)

M.1: [Initialize]
    minPrice ← ∞
    maxProfit ← 0

M.2: [Iterate through price array]
    For each price in prices:
        If price < minPrice:
            minPrice ← price
        Else if price - minPrice > maxProfit:
            maxProfit ← price - minPrice

M.3: [Return result]
    Return maxProfit

Complexity:
    Time: O(n)
    Space: O(1)

```


-


> **Best Time to Buy and Sell Stock II**

```
Algorithm MP (Maximize profit from multiple stock trades)
Input:
    - prices: array of integers where prices[i] is the stock price on day i
Output:
    - Maximum total profit achievable through multiple transactions

MP1. [Handle edge case]
    If length of prices ≤ 1:
        Return 0

MP2. [Initialize profit counter]
    Set profit ← 0

MP3. [Iterate through price array and accumulate upward differences]
    For i from 1 to length(prices) - 1:
        If prices[i] > prices[i - 1]:
            Set profit ← profit + (prices[i] - prices[i - 1])

MP4. [Return final profit]
    Return profit

Complexity:
    Time: O(n), where n is the number of days — one pass through the prices.
    Space: O(1), only a single integer is used for accumulation.

```


-


> **Jump Game**

```
Algorithm J (Check if end of array is reachable via jumps)
Input:
    - nums: array of integers where nums[i] is max jump length at index i
Output:
    - true if last index is reachable, false otherwise

J1. [Handle trivial case]
    If length of nums ≤ 1:
        Return true

J2. [Initialize farthest reachable index]
    Set farthest ← 0

J3. [Iterate through array and update farthest reachable position]
    For i from 0 to length(nums) - 1:
        If i > farthest:
            Return false
        Set farthest ← max(farthest, i + nums[i])
        If farthest ≥ length(nums) - 1:
            Return true

J4. [Final check]
    Return false

Complexity:
    Time: O(n), where n is the length of nums — single pass through the array.
    Space: O(1), only one integer is used for tracking progress.

```


-


> **Jump Game II**

```
Algorithm M (Compute minimum jumps to reach end of array)
Input:
    - nums: array of integers where nums[i] is max jump length at index i
Output:
    - Minimum number of jumps to reach last index

M1. [Initialize jump counters and window]
    Set jumps ← 0
    Set currentEnd ← 0
    Set farthest ← 0

M2. [Iterate through array to find jump boundaries]
    For i from 0 to length(nums) - 2:
        Set farthest ← max(farthest, i + nums[i])
        If i == currentEnd:
            Set jumps ← jumps + 1
            Set currentEnd ← farthest

M3. [Return total jumps needed]
    Return jumps

Complexity:
    Time: O(n), where n is the length of nums — we scan the array once.
    Space: O(1), constant space is used for tracking jumps and bounds.


```


-



> **H-Index**

```
Algorithm H (Compute h-index from citation counts)

Input:
    - citations: an array of integers of length n, where citations[i] is the number of citations for the ith paper.

Output:
    - The h-index of the researcher.

H1. [Sort the citations in descending order]  
    Sort citations in non-increasing order and store in sortedCitations.

H2. [Initialize h-index counter]  
    Set h ← 0

H3. [Scan sorted list and determine h-index]  
    For i from 0 to n − 1:
        If sortedCitations[i] ≥ i + 1:
            Set h ← i + 1
        Else:
            Break the loop

H4. [Return h-index]  
    Return h

Complexity:  
    Time: O(n log n) due to sorting  
    Space: O(n) if a new array is made during sorting; O(1) if sorting in-place

```


-






> **H-Index**

```
Algorithm H' (Compute h-index with counting)

Input:
    - citations: an array of n non-negative integers representing citation counts

Output:
    - h-index value (integer)

H'.1: [Initialize count buckets]
    Create count array of size n + 1 and initialize all to 0

H'.2: [Bucket citation counts]
    For each citation c in citations:
        If c ≥ n:
            Increment count[n]
        Else:
            Increment count[c]

H'.3: [Walk from high to low and accumulate]
    Set total ← 0
    For i from n down to 0:
        total ← total + count[i]
        If total ≥ i:
            Return i

H'.4: [Default return]
    Return 0  // This should never be reached given problem constraints

Complexity:
    Time: O(n)
    Space: O(n)

```


-



> **Insert Delete GetRandom O(1)**

```
Algorithm R (RandomizedSet with O(1) insert, remove, and getRandom)
Input:
    A dynamic stream of operations on a set of integers: Insert(val), Remove(val), GetRandom()

Output:
    Result of each operation: true/false for Insert and Remove, integer for GetRandom

Data Structures:
    nums: dynamic array of integers
    idxMap: hashmap mapping int → index in nums

R.1: Initialize the RandomizedSet
    Set nums ← empty list
    Set idxMap ← empty map

R.2: Insert(val)
    If val ∈ idxMap:
        Return false
    Append val to nums
    Let i ← length(nums) − 1
    Set idxMap[val] ← i
    Return true

R.3: Remove(val)
    If val ∉ idxMap:
        Return false
    Let i ← idxMap[val]
    Let last ← nums[length(nums) − 1]
    Replace nums[i] ← last
    Set idxMap[last] ← i
    Remove last element from nums
    Delete idxMap[val]
    Return true

R.4: GetRandom()
    Let r ← random integer in [0, length(nums) − 1]
    Return nums[r]

Time Complexity:
    Insert: O(1) average
    Remove: O(1) average
    GetRandom: O(1)

Space Complexity:
    O(n) for both nums and idxMap, where n = number of elements in set


```


-



> **Product of Array Except Self**

```
Algorithm P (Product of Array Except Self, without division, O(n) time)
Input:
    nums: an array of integers of length n

Output:
    answer: an array where answer[i] = product of all nums[j] for j ≠ i

P.1: Initialize output and prefix product array
    Let n ← length of nums
    Let answer ← array of size n initialized to 1

P.2: Compute prefix products for each index
    Let prefix ← 1
    For i ← 0 to n−1 do:
        Set answer[i] ← prefix
        Set prefix ← prefix × nums[i]

P.3: Compute suffix products and multiply with prefix products
    Let suffix ← 1
    For i ← n−1 downto 0 do:
        Set answer[i] ← answer[i] × suffix
        Set suffix ← suffix × nums[i]

P.4: Return the answer array
    Return answer

Time Complexity: O(n)
Space Complexity: O(1) extra space (excluding output array)

```


-



> **Gas Station**

```
Algorithm G (Gas Station Circuit Feasibility)
Input:
    gas: array of integers where gas[i] is amount of gas at station i
    cost: array of integers where cost[i] is gas needed to go from i to (i+1)%n

Output:
    index of the starting station if the circuit can be completed, otherwise -1

G.1: Initialize total and current tank gas
    Let n ← length of gas
    Let total ← 0     // tracks net gas over entire circuit
    Let tank ← 0      // tracks gas in tank on current path
    Let start ← 0     // candidate starting index

G.2: Traverse each station from 0 to n−1
    For i ← 0 to n−1 do:
        Let gain ← gas[i] − cost[i]
        Set tank ← tank + gain
        Set total ← total + gain
        If tank < 0 then:
            Set start ← i + 1
            Set tank ← 0

G.3: Check if total gas is non-negative
    If total < 0 then:
        Return -1
    Else:
        Return start

Time Complexity: O(n)
Space Complexity: O(1)

```


-



> **Candy**

```
Algorithm C (Candy Distribution with Rating Constraints)
Input:
    ratings: array of integers of length n, where ratings[i] is the rating of the i-th child

Output:
    Minimum total number of candies needed to satisfy all conditions

C.1: Initialize base candy distribution
    Let n ← length of ratings
    Let candies[0...n−1] ← array initialized with 1 for all positions

C.2: Left-to-right pass to ensure increasing ratings get more candies
    For i ← 1 to n−1 do:
        If ratings[i] > ratings[i−1] then:
            Set candies[i] ← candies[i−1] + 1

C.3: Right-to-left pass to ensure decreasing ratings still get more candies
    For i ← n−2 down to 0 do:
        If ratings[i] > ratings[i+1] then:
            Set candies[i] ← max(candies[i], candies[i+1] + 1)

C.4: Sum and return total candies
    Return sum of all elements in candies

Time Complexity: O(n)  
Space Complexity: O(n)

```


-



> **Trapping Rain Water**

```
Algorithm T (Trapping Rain Water with Two Pointers)
Input:
    height: array of non-negative integers of length n representing elevation heights

Output:
    Total water trapped between the bars after raining

T.1: Initialize pointers and state
    Let n ← length of height
    If n < 3 then return 0
    Let left ← 0, right ← n−1
    Let leftMax ← 0, rightMax ← 0
    Let water ← 0

T.2: Move pointers inward while computing trapped water
    While left < right do:
        If height[left] < height[right] then:
            If height[left] ≥ leftMax then:
                Set leftMax ← height[left]
            Else:
                Set water ← water + (leftMax − height[left])
            Increment left ← left + 1
        Else:
            If height[right] ≥ rightMax then:
                Set rightMax ← height[right]
            Else:
                Set water ← water + (rightMax − height[right])
            Decrement right ← right − 1

T.3: Return accumulated water
    Return water

Time Complexity: O(n)  
Space Complexity: O(1)

```


-



> **Roman to Integer**

```
Algorithm R (Roman to Integer Conversion)
Input:
    s: A Roman numeral string of length n

Output:
    The corresponding integer value

R.1: Define symbol-to-value map
    Let value be a map from rune to integer:
        'I' → 1, 'V' → 5, 'X' → 10, 'L' → 50,
        'C' → 100, 'D' → 500, 'M' → 1000

R.2: Initialize total and iterate
    Let total ← 0
    Let n ← length of s
    For i ← 0 to n−1 do:
        If i < n−1 and value[s[i]] < value[s[i+1]] then:
            total ← total − value[s[i]]  // subtraction case
        Else:
            total ← total + value[s[i]]  // normal case

R.3: Return total
    Return total

Time Complexity: O(n)  
Space Complexity: O(1)

```


-



> **Integer to Roman**

```
Algorithm T (Integer to Roman Numeral Conversion)
Input:
    num: A positive integer in range [1, 3999]

Output:
    A string representing the Roman numeral of the input integer

T.1: Prepare value-symbol pairs
    Let vals ← [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
    Let syms ← ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"]

T.2: Initialize result string
    Let res ← ""

T.3: Construct Roman numeral greedily
    For i ← 0 to 12 do:
        While num ≥ vals[i] do:
            res ← res + syms[i]
            num ← num − vals[i]

T.4: Return result
    Return res

Time Complexity: O(1)  
    (since maximum 13 iterations and Roman numerals are bounded for num ∈ [1, 3999])  
Space Complexity: O(1)

```


-


> **Length of Last Word**

```
Algorithm L (Length of Last Word)
Input:
    s: A string consisting of letters and space characters

Output:
    An integer representing the length of the last word in the string

L.1: Initialize index to end of string
    Let i ← len(s) − 1

L.2: Skip trailing spaces
    While i ≥ 0 and s[i] = ' ' do:
        i ← i − 1

L.3: Count characters of the last word
    Let length ← 0
    While i ≥ 0 and s[i] ≠ ' ' do:
        length ← length + 1
        i ← i − 1

L.4: Return the length
    Return length

Time Complexity: O(n) where n is the length of the string  
Space Complexity: O(1)

```


-

> **Longest Common Prefix**

```
Algorithm P (Longest Common Prefix)
Input:
    strs: An array of strings of length n

Output:
    A string representing the longest common prefix shared among all strings in strs

P.1: Handle empty input
    If n = 0 then:
        Return ""

P.2: Initialize prefix
    Let prefix ← strs[0]

P.3: Iterate over remaining strings
    For i ← 1 to n−1 do:
        While strs[i] does not start with prefix do:
            Remove last character from prefix
            If prefix = "" then:
                Return ""

P.4: Return prefix
    Return prefix

Time Complexity: O(S) where S is the sum of all characters in all strings (in worst case all characters are compared)
Space Complexity: O(1) extra space (modifies prefix in-place)



Algorithm P (Longest Common Prefix Without Built-in Prefix Check)
Input:
    strs: An array of strings of length n

Output:
    A string representing the longest common prefix shared among all strings in strs

P.1: Handle empty input
    If n = 0 then:
        Return ""

P.2: Initialize prefix
    Let prefix ← strs[0]

P.3: Iterate over remaining strings
    For i ← 1 to n−1 do:
        While not IsPrefix(strs[i], prefix) do:
            Remove last character from prefix
            If prefix = "" then:
                Return ""

P.4: Return prefix
    Return prefix

Helper Algorithm IsPrefix(s, p)
Input:
    s: a string
    p: a potential prefix

Output:
    true if p is a prefix of s; false otherwise

IsPrefix.1: If len(p) > len(s), return false  
IsPrefix.2: For j ← 0 to len(p)−1 do:
    If s[j] ≠ p[j], return false
IsPrefix.3: Return true

Time Complexity: O(S) where S is total characters compared across all strings  
Space Complexity: O(1) auxiliary space

```


-

> **Reverse Words in a String**

```
Algorithm W (Reverse the order of words in a string)
Input:
    - s: a string possibly containing multiple spaces and words
Output:
    - A string with words in reversed order, separated by a single space

W1. [Trim and split]
    Remove leading and trailing spaces from s
    Initialize an empty list words ← []
    Set i ← 0
    While i < len(s):
        If s[i] = ' ':
            Increment i
            Continue
        Set j ← i
        While j < len(s) and s[j] ≠ ' ':
            Increment j
        Append s[i:j] to words
        Set i ← j

W2. [Reverse the words]
    Reverse the list words in-place

W3. [Join and return]
    Return the result of joining the words list with a single space

Complexity:
    Time: O(n), where n is the length of the input string.
        - One pass to split words (skip extra spaces)
        - One pass to reverse
        - One pass to join
    Space: O(n), to store the words in a list and the resulting string.

```


-

> **Zigzag Conversion**

```
Algorithm Z (Zigzag convert string into given number of rows)

Input:
    - s: string to convert
    - numRows: number of rows in zigzag pattern

Output:
    - A string that reads the zigzag pattern line by line

Z.1: [Edge case]
    If numRows = 1 or len(s) ≤ numRows:
        Return s

Z.2: [Initialize row buffers]
    Create a list rows with numRows empty strings
    Set currentRow ← 0
    Set goingDown ← false

Z.3: [Distribute characters]
    For each character c in s:
        Append c to rows[currentRow]
        If currentRow = 0 or currentRow = numRows - 1:
            Flip goingDown
        If goingDown:
            currentRow ← currentRow + 1
        Else:
            currentRow ← currentRow - 1

Z.4: [Concatenate rows]
    Join all rows into a single string and return

Time Complexity: O(n), where n is len(s)
Space Complexity: O(n), to hold the final result and intermediate rows

```


-

> **Find the Index of the First Occurrence in a String**

```
Algorithm F (Find the index of the first occurrence of needle in haystack)

Input:
    - haystack: string to search within
    - needle: string to search for

Output:
    - Index of first occurrence of needle in haystack, or -1 if not found

F1. [Handle edge case]
    If len(needle) = 0:
        Return 0

F2. [Iterate through haystack]
    For i ← 0 to len(haystack) - len(needle):
        If haystack[i : i + len(needle)] = needle:
            Return i

F3. [Not found]
    Return -1

Time Complexity: O(n * m), where n = len(haystack), m = len(needle)
    - Worst-case comparison of m characters at n - m + 1 positions
Space Complexity: O(1)

```


-

> **Text Justification**

```
Algorithm J (Justify text with even spaces, fitting each line into maxWidth)

Input:
    - words: list of strings (non-space words)
    - maxWidth: integer specifying the width of each line

Output:
    - A list of strings, each fully justified to maxWidth

J1. [Initialize]
    result ← empty list
    i ← 0

J2. [Build lines greedily]
    While i < len(words):
        - lineLen ← len(words[i])
        - j ← i + 1
        While j < len(words) and lineLen + 1 + len(words[j]) ≤ maxWidth:
            lineLen ← lineLen + 1 + len(words[j])
            j ← j + 1

        - lineWords ← words[i:j]
        - isLast ← j = len(words)

J3. [Format current line]
    If isLast OR len(lineWords) = 1:
        - line ← join(lineWords with single space)
        - Add spaces at end to make line length = maxWidth
    Else:
        - totalSpaces ← maxWidth − totalCharsIn(lineWords)
        - slots ← len(lineWords) − 1
        - baseSpace ← totalSpaces ÷ slots
        - extra ← totalSpaces mod slots

        - For each word in lineWords:
            Add baseSpace spaces (plus one extra if leftmost slot < extra)

J4. [Add to result]
    Append line to result
    Set i ← j and repeat from J2

J5. [Return result]
    Return result

Complexity:
    Time: O(n), where n is total number of characters in all words
    Space: O(n), for the output list of lines

```


-


> **Palindrome Number**

```
Algorithm P (Check if an integer is a palindrome)
Input:
    - x: an integer
Output:
    - true if x is a palindrome; false otherwise

P1. [Handle edge cases]
    If x < 0 or (x ≠ 0 and x mod 10 = 0):
        Return false
    // Negative numbers and numbers ending in 0 (but not 0 itself) can't be palindromes

P2. [Initialize variables]
    Set reverted ← 0

P3. [Reverse half of the digits]
    While x > reverted:
        Set reverted ← reverted × 10 + x mod 10
        Set x ← x ÷ 10

P4. [Check for palindrome]
    If x = reverted or x = reverted ÷ 10:
        Return true
    Else:
        Return false

Complexity:
    Time: O(log₁₀(n)), where n is the input number x. We reverse at most half the digits.
    Space: O(1), using constant space.

```


-


> **Plus One**

```
Algorithm I (Increment large integer represented as an array)
Input:
    - digits: array of integers where each element is a digit (most significant first)
Output:
    - new array representing the integer + 1

I1. [Traverse digits from right to left]
    For i ← len(digits) - 1 downto 0:
        If digits[i] < 9:
            Increment digits[i] by 1
            Return digits
        Else:
            Set digits[i] ← 0
    // If loop completes, all digits were 9

I2. [All digits were 9, handle overflow]
    Prepend 1 to digits
    Return [1] + digits

Complexity:
    Time: O(n), where n is the length of digits
    Space: O(n), worst case when carry propagates through all digits

```


-


> **Factorial Trailing Zeroes**

```
Algorithm T (Count trailing zeroes in factorial)
Input: Integer n (where 0 ≤ n ≤ 10⁴)
Output: Number of trailing zeroes in n!

T.1: Initialize counter
    Let count ← 0

T.2: Count multiples of 5, 25, 125, ...
    While n ≥ 5 do
        n ← floor(n / 5)
        count ← count + n

T.3: Return result
    Return count

Time Complexity: O(log₅(n))
Space Complexity: O(1)

```


-


> **Sqrt(x)**

```
Algorithm S (Integer square root using binary search)
Input: Integer x (x ≥ 0)
Output: ⌊√x⌋ (square root of x rounded down)

S.1: Handle edge case
    If x < 2, return x

S.2: Initialize binary search bounds
    Let left ← 1
    Let right ← x / 2

S.3: Perform binary search
    While left ≤ right do
        Let mid ← (left + right) / 2
        Let midSquared ← mid × mid

        If midSquared = x
            Return mid

        If midSquared < x
            Let left ← mid + 1
            Let answer ← mid
        Else
            Let right ← mid - 1

S.4: Return final answer
    Return answer

Time Complexity: O(log x)  
Space Complexity: O(1)

```


-


> **Pow(x, n)**

```
Algorithm P (Power: Compute xⁿ efficiently)
Input: float x, integer n
Output: x raised to power n (i.e., xⁿ)

P.1: Handle base case
    If n = 0, return 1

P.2: Handle negative exponent
    If n < 0
        Set x ← 1 / x
        Set n ← -n

P.3: Initialize result
    Let result ← 1.0

P.4: Iteratively square x and halve n
    While n > 0 do
        If n is odd
            Set result ← result × x
        Set x ← x × x
        Set n ← n / 2

P.5: Return result
    Return result

Time Complexity: O(log |n|)  
Space Complexity: O(1)

```


-


> **Max Points on a Line**

```
Algorithm L (LineMax: Max points on a straight line)
Input: A list of points on a 2D plane, `points` where each point = [x, y]
Output: Maximum number of points that lie on the same straight line

L.1: Initialize global max
    Let n ← length of points
    If n ≤ 2, return n
    Let maxPoints ← 0

L.2: For each point as anchor
    For i from 0 to n - 1 do
        Let slopes ← empty hash map from string → int
        Let duplicates ← 1  // count of points identical to points[i]
        Let localMax ← 0

        For j from i+1 to n - 1 do
            Let dx ← points[j][0] - points[i][0]
            Let dy ← points[j][1] - points[i][1]

            If dx == 0 and dy == 0 then
                duplicates ← duplicates + 1
                Continue

            Let g ← GCD(dy, dx)
            dy ← dy / g
            dx ← dx / g

            If dx < 0 then
                dy ← -dy
                dx ← -dx  // Normalize to avoid equivalent slopes with different signs

            Let slopeKey ← string representation of (dy, dx)
            slopes[slopeKey] ← slopes.get(slopeKey, 0) + 1
            localMax ← max(localMax, slopes[slopeKey])

        maxPoints ← max(maxPoints, localMax + duplicates)

L.3: Return result
    Return maxPoints

Helper: GCD(a, b)
    While b ≠ 0
        a, b ← b, a mod b
    Return abs(a)

Time Complexity: O(n²)  
Space Complexity: O(n)

```


-


> **Valid Palindrome**

```
Algorithm P (Check if a phrase is a valid palindrome after normalization)
Input:
    - s: a string containing letters, digits, spaces, and symbols
Output:
    - true if the cleaned string is a palindrome, false otherwise

P1. [Initialize pointers]
    Set left ← 0, right ← len(s) - 1

P2. [Loop until pointers cross]
    While left < right:
        P2.1: [Skip non-alphanumeric on left]
            While left < right and s[left] is not alphanumeric:
                Increment left

        P2.2: [Skip non-alphanumeric on right]
            While left < right and s[right] is not alphanumeric:
                Decrement right

        P2.3: [Compare characters]
            If lowercase(s[left]) ≠ lowercase(s[right]):
                Return false

        P2.4: [Move pointers]
            Increment left, decrement right

P3. [If loop completes]
    Return true

Complexity:
    Time: O(n), where n is the length of the input string.
        Each character is visited at most once.
    Space: O(1), using only constant space for pointers and comparisons.

```


-


> **Is Subsequence**

```
Algorithm Q (Check if string s is a subsequence of string t)
Input:
    - s: a string of lowercase English letters
    - t: a string of lowercase English letters
Output:
    - true if s is a subsequence of t, false otherwise

Q1. [Initialize pointers]
    Set i ← 0, j ← 0

Q2. [Traverse both strings]
    While i < len(s) and j < len(t):
        Q2.1: [Match character]
            If s[i] = t[j]:
                Increment i
        Q2.2: [Advance in t regardless]
            Increment j

Q3. [Check if all of s was matched]
    If i = len(s):
        Return true
    Else:
        Return false

Complexity:
    Time: O(n), where n = len(t)
        Each character of t is visited at most once.
    Space: O(1), uses constant space.

```


-


> **Two Sum II - Input Array Is Sorted**

```
Algorithm T (Find two numbers in a sorted array that sum to a target)
Input:
    - numbers: a 1-indexed array of integers sorted in non-decreasing order
    - target: an integer representing the target sum
Output:
    - An array [index1, index2] such that numbers[index1] + numbers[index2] = target

T1. [Initialize two pointers]
    Set left ← 0, right ← len(numbers) - 1

T2. [Scan for the target sum]
    While left < right:
        Let sum ← numbers[left] + numbers[right]

        T2.1: [Check for match]
            If sum = target:
                Return [left + 1, right + 1]

        T2.2: [Adjust pointers]
            If sum < target:
                Increment left
            Else:
                Decrement right

T3. [Guaranteed solution]
    Since the problem guarantees exactly one solution, this point is never reached.

Complexity:
    Time: O(n), where n = len(numbers)
        Each element is visited at most once by either pointer.
    Space: O(1), uses only constant extra space.

```


-


> **Container With Most Water**

```
Algorithm C (Find the maximum area of water that can be contained between vertical lines)
Input:
    - height: an array of integers where height[i] represents the height of a vertical line at position i
Output:
    - The maximum amount of water a container can store

C1. [Initialize two pointers and max area]
    Set left ← 0, right ← len(height) - 1
    Set maxArea ← 0

C2. [Scan with two-pointer approach]
    While left < right:
        C2.1: [Compute area]
            Let h ← min(height[left], height[right])
            Let w ← right - left
            Let area ← h × w

        C2.2: [Update max area]
            Set maxArea ← max(maxArea, area)

        C2.3: [Move the pointer at shorter line]
            If height[left] < height[right]:
                Increment left
            Else:
                Decrement right

C3. [Return the result]
    Return maxArea

Complexity:
    Time: O(n), where n is the length of the height array
        Each index is visited at most once.
    Space: O(1), uses constant extra space.

```


-


> **3Sum**

```
Algorithm S (Find all unique triplets in an array that sum to zero)
Input:
    - nums: an array of integers
Output:
    - A list of lists, where each inner list is a unique triplet [a, b, c] such that a + b + c = 0

S1. [Sort the array]
    Sort nums in non-decreasing order

S2. [Initialize result container]
    Let result ← empty list

S3. [Iterate through nums with a fixed pointer i]
    For i from 0 to len(nums) - 3:
        S3.1: [Skip duplicates for i]
            If i > 0 and nums[i] = nums[i - 1]:
                Continue

        S3.2: [Initialize two pointers]
            Set left ← i + 1, right ← len(nums) - 1

        S3.3: [Two-pointer scan]
            While left < right:
                Let sum ← nums[i] + nums[left] + nums[right]

                S3.3.1: [Found a triplet]
                    If sum = 0:
                        Append [nums[i], nums[left], nums[right]] to result
                        Increment left while nums[left] = nums[left - 1] to skip duplicates
                        Decrement right while nums[right] = nums[right + 1] to skip duplicates
                        Increment left, decrement right

                S3.3.2: [Adjust pointers]
                    Else if sum < 0:
                        Increment left
                    Else:
                        Decrement right

S4. [Return result]
    Return result

Complexity:
    Time: O(n²), where n is the length of nums
        Sorting takes O(n log n), and each i uses a two-pointer scan in O(n) time
    Space: O(log n) for sorting if done in-place, O(k) for result with k triplets

```


-


> **Minimum Size Subarray Sum**

```
Algorithm M (Find minimal length subarray with sum ≥ target in an array of positive integers)

Input:
    - target: positive integer target sum
    - nums: array of positive integers

Output:
    - Minimal length of a contiguous subarray whose sum ≥ target, or 0 if no such subarray exists

M1. [Initialize variables]
    Set n ← length of nums
    Set left ← 0, sum ← 0, minLen ← ∞

M2. [Expand right pointer]
    For right from 0 to n - 1:
        sum ← sum + nums[right]

M3. [Shrink left pointer while condition met]
    While sum ≥ target:
        minLen ← min(minLen, right - left + 1)
        sum ← sum - nums[left]
        left ← left + 1

M4. [Check if a valid subarray was found]
    If minLen = ∞:
        Return 0

M5. [Return result]
    Return minLen

Complexity:
    Time: O(n), where n is the length of nums. Each element is visited at most twice (once by right, once by left).
    Space: O(1), constant extra space.

```


-


> **Longest Substring Without Repeating Characters**

```
Algorithm L (Find length of the longest substring without duplicate characters)

Input:
    - s: a string

Output:
    - Length of the longest substring of s without repeating characters

L1. [Initialize variables]
    Create an empty map charIndex to store the last seen index of characters
    Set left ← 0, maxLen ← 0

L2. [Expand right pointer]
    For right from 0 to length(s) - 1:
        Let c ← s[right]
        If c exists in charIndex and charIndex[c] ≥ left:
            Set left ← charIndex[c] + 1
        Set charIndex[c] ← right
        maxLen ← max(maxLen, right - left + 1)

L3. [Return result]
    Return maxLen

Complexity:
    Time: O(n), where n is the length of s. Each character is visited at most twice (once by right, once by left).
    Space: O(min(n, a)), where a is the alphabet size (number of possible characters).

```


-


> **Substring with Concatenation of All Words**

```
Algorithm C (Find starting indices of substrings that are concatenations of all words in any order)

Input:
    - s: a string
    - words: an array of strings, all of the same length

Output:
    - An array of starting indices of substrings in s that are concatenations of all the words in any order

C1. [Handle edge cases]
    If s is empty, words is empty, or words[0] is empty:
        Return empty array

C2. [Precompute parameters]
    Set wordLen ← length(words[0])
    Set wordCount ← length(words)
    Set totalLen ← wordLen × wordCount
    Create a map wordFreq to count occurrences of each word in words

C3. [Iterate over possible starting offsets]
    For offset from 0 to wordLen - 1:
        Set left ← offset, right ← offset, matched ← 0
        Create empty map seen

        While right + wordLen ≤ length(s):
            Let word ← s[right : right + wordLen]
            right ← right + wordLen

            If word exists in wordFreq:
                Increment seen[word] by 1
                matched ← matched + 1

                While seen[word] > wordFreq[word]:
                    Let leftWord ← s[left : left + wordLen]
                    Decrement seen[leftWord] by 1
                    left ← left + wordLen
                    matched ← matched - 1

                If matched = wordCount:
                    Append left to result
                    Let leftWord ← s[left : left + wordLen]
                    Decrement seen[leftWord] by 1
                    left ← left + wordLen
                    matched ← matched - 1
            Else:
                Clear seen, matched ← 0
                left ← right

C4. [Return result]
    Return result array

Complexity:
    Time: O(n × wordLen), where n is the length of s. The outer loop runs wordLen times, and inner window sliding is O(n) total per offset.
    Space: O(k), where k is the number of distinct words in words.

```


-


> **Minimum Window Substring**

```
Algorithm W (Find the minimum window substring of s containing all characters of t)

Input:
    - s: a string of length m
    - t: a string of length n

Output:
    - The smallest substring of s containing all characters of t (including duplicates), or "" if no such substring exists

W1. [Handle edge cases]
    If m < n or t is empty:
        Return ""

W2. [Build target frequency map]
    Create map targetFreq to store frequency of each character in t
    For each character c in t:
        targetFreq[c] ← targetFreq[c] + 1

W3. [Initialize window variables]
    Create empty map windowFreq
    Set required ← number of distinct keys in targetFreq
    Set formed ← 0
    Set left ← 0
    Set minLen ← ∞, minStart ← 0

W4. [Expand right pointer]
    For right from 0 to m - 1:
        c ← s[right]
        Increment windowFreq[c] by 1

        If c exists in targetFreq and windowFreq[c] = targetFreq[c]:
            formed ← formed + 1

W5. [Contract from left while valid]
    While left ≤ right and formed = required:
        If (right - left + 1) < minLen:
            minLen ← right - left + 1
            minStart ← left

        d ← s[left]
        Decrement windowFreq[d] by 1
        If d exists in targetFreq and windowFreq[d] < targetFreq[d]:
            formed ← formed - 1
        left ← left + 1

W6. [Return result]
    If minLen = ∞:
        Return ""
    Else:
        Return s[minStart : minStart + minLen]

Complexity:
    Time: O(m + n), where m = len(s), n = len(t). Each character is visited at most twice (once by right, once by left).
    Space: O(k), where k is the number of distinct characters in s and t combined.

```


-


> **Valid Sudoku**

```
Algorithm S (Validate Sudoku Board)

Input:
    - board: 9×9 array of characters ('1'-'9' or '.')

Output:
    - true if board is valid, false otherwise

S1. [Initialize tracking structures]
    Create:
        rows[9] → set of digits in each row
        cols[9] → set of digits in each column
        boxes[9] → set of digits in each 3×3 sub-box

S2. [Iterate over all cells]
    For r from 0 to 8:
        For c from 0 to 8:
            val ← board[r][c]
            If val == '.': continue

            boxIndex ← (r/3) * 3 + (c/3)   // integer division

            If val in rows[r] → return false
            If val in cols[c] → return false
            If val in boxes[boxIndex] → return false

            Add val to rows[r], cols[c], boxes[boxIndex]

S3. [If no rule violations]
    Return true

Complexity:
    Time: O(1) because the board is always 9×9 (81 cells)
    Space: O(1) constant-size tracking structures



Explanation:
For the example
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]


For every filled cell we show:

    (r,c) — zero-indexed coordinates, row then column.
    val — the digit in that cell.
    box — computed as (r/3)*3 + (c/3) (0..8).
    check — whether adding val would immediately violate a seen duplicate 
            (ok = no duplicate; violation = duplicate found).
    row[r] after / col[c] after / box[b] after — the contents of those trackers 
            after we insert the value (only if check == ok we perform the insertion).

Initial trackers are empty for all rows, cols, boxes.
Step | (r,c)  | val | box | check    | row[r] after    | col[c] after    | box[b] after   
------------------------------------------------------------------------------------------
   1 | (0,0)  | 5   |   0 | ok       | {5}             | {5}             | {5}            
   2 | (0,1)  | 3   |   0 | ok       | {3,5}           | {3}             | {3,5}          
   3 | (0,4)  | 7   |   1 | ok       | {3,5,7}         | {7}             | {7}            
   4 | (1,0)  | 6   |   0 | ok       | {6}             | {5,6}           | {3,5,6}        
   5 | (1,3)  | 1   |   1 | ok       | {1,6}           | {1}             | {1,7}          
   6 | (1,4)  | 9   |   1 | ok       | {1,6,9}         | {7,9}           | {1,7,9}        
   7 | (1,5)  | 5   |   1 | ok       | {1,5,6,9}       | {5}             | {1,5,7,9}      
   8 | (2,1)  | 9   |   0 | ok       | {9}             | {3,9}           | {3,5,6,9}      
   9 | (2,2)  | 8   |   0 | ok       | {8,9}           | {8}             | {3,5,6,8,9}    
  10 | (2,7)  | 6   |   2 | ok       | {6,8,9}         | {6}             | {6}            
  11 | (3,0)  | 8   |   3 | ok       | {8}             | {5,6,8}         | {8}            
  12 | (3,4)  | 6   |   4 | ok       | {6,8}           | {6,7,9}         | {6}            
  13 | (3,8)  | 3   |   5 | ok       | {3,6,8}         | {3}             | {3}            
  14 | (4,0)  | 4   |   3 | ok       | {4}             | {4,5,6,8}       | {4,8}          
  15 | (4,3)  | 8   |   4 | ok       | {4,8}           | {1,8}           | {6,8}          
  16 | (4,5)  | 3   |   4 | ok       | {3,4,8}         | {3,5}           | {3,6,8}        
  17 | (4,8)  | 1   |   5 | ok       | {1,3,4,8}       | {1,3}           | {1,3}          
  18 | (5,0)  | 7   |   3 | ok       | {7}             | {4,5,6,7,8}     | {4,7,8}        
  19 | (5,4)  | 2   |   4 | ok       | {2,7}           | {2,6,7,9}       | {2,3,6,8}      
  20 | (5,8)  | 6   |   5 | ok       | {2,6,7}         | {1,3,6}         | {1,3,6}        
  21 | (6,1)  | 6   |   6 | ok       | {6}             | {3,6,9}         | {6}            
  22 | (6,6)  | 2   |   8 | ok       | {2,6}           | {2}             | {2}            
  23 | (6,7)  | 8   |   8 | ok       | {2,6,8}         | {6,8}           | {2,8}          
  24 | (7,3)  | 4   |   7 | ok       | {4}             | {1,4,8}         | {4}            
  25 | (7,4)  | 1   |   7 | ok       | {1,4}           | {1,2,6,7,9}     | {1,4}          
  26 | (7,5)  | 9   |   7 | ok       | {1,4,9}         | {3,5,9}         | {1,4,9}        
  27 | (7,8)  | 5   |   8 | ok       | {1,4,5,9}       | {1,3,5,6}       | {2,5,8}        
  28 | (8,4)  | 8   |   7 | ok       | {8}             | {1,2,6,7,8,9}   | {1,4,8,9}      
  29 | (8,7)  | 7   |   8 | ok       | {7,8}           | {6,7,8}         | {2,5,7,8}      
  30 | (8,8)  | 9   |   8 | ok       | {7,8,9}         | {1,3,5,6,9}     | {2,5,7,8,9}    

How to read this and why it proves validity

- Invariant maintained: after processing step k, for every processed filled cell (r,c) 
  the trackers record the exact digits seen so far in that row, that column, and that box.
    rows[r] contains every digit that has already appeared in row r among cells processed so far.
    cols[c] contains every digit that has already appeared in column c.
    boxes[b] contains every digit that has already appeared in 3×3 box b.

 - Duplicate detection: before insertion we test val in rows[r] || val in cols[c] || val in boxes[box]. 
   If any is true, we immediately return false (invalid). That check is precisely what the check column 
   shows. Every step above is ok (no duplicate), so we insert.

 - O(1) checks and updates: each check and insertion is constant time (map lookup + map write),
   so we avoid re-scanning rows/cols/boxes repeatedly.

Conclusion for this board: 
All 30 filled cells were inspected in row-major order; 
no duplicate was detected in any row, column, or 3×3 box.
Therefore the validator would finish the scan and return true.


```


-


> **Spiral Matrix**

```
Algorithm S (Return all elements of an m × n matrix in spiral order)

Input:
    - matrix: a 2D array of integers with dimensions m × n

Output:
    - result: a 1D array containing all elements in spiral order

S1. [Handle edge case]
    If m = 0:
        Return empty result

S2. [Initialize boundaries]
    top ← 0
    bottom ← m - 1
    left ← 0
    right ← n - 1
    result ← empty list

S3. [Traverse while boundaries are valid]
    While top ≤ bottom and left ≤ right:

        S3.1. [Traverse from left to right along top row]
            For col from left to right:
                Append matrix[top][col] to result
            Set top ← top + 1

        S3.2. [Traverse from top to bottom along right column]
            For row from top to bottom:
                Append matrix[row][right] to result
            Set right ← right - 1

        S3.3. [Traverse from right to left along bottom row if in bounds]
            If top ≤ bottom:
                For col from right down to left:
                    Append matrix[bottom][col] to result
                Set bottom ← bottom - 1

        S3.4. [Traverse from bottom to top along left column if in bounds]
            If left ≤ right:
                For row from bottom down to top:
                    Append matrix[row][left] to result
                Set left ← left + 1

S4. [Return result]
    Return result

Complexity:
    Time: O(m × n), where m is number of rows and n is number of columns.
        Each element is visited exactly once.
    Space: O(1), excluding the output array.

```


-


> **Rotate Image**

```
Algorithm M (Rotate an n × n matrix 90 degrees clockwise in-place)

Input:
    - matrix: an n × n matrix of elements

Output:
    - The same matrix rotated 90 degrees clockwise

M1. [Transpose the matrix]
    For each row index i from 0 to n - 1:
        For each column index j from i + 1 to n - 1:
            Exchange the values at positions (i, j) and (j, i) in matrix.

M2. [Reverse each row]
    For each row index i from 0 to n - 1:
        Reverse the elements of row i without allocating new memory.

M3. [Termination]
    Return the modified matrix.

Complexity:
    Time: O(n²), because each element is visited a constant number of times.
    Space: O(1), since the rotation is performed in-place without extra storage.

```


-



> **Set Matrix Zeroes**

```
Algorithm Z (Set entire row and column to zero if any element is zero, in-place)

Input:
    - matrix: an m × n integer matrix

Output:
    - The same matrix with all rows and columns containing a zero set entirely to zero

Z1. [Check first row and column zero status]
    Let firstRowZero ← false
    Let firstColZero ← false
    For each column index j from 0 to n - 1:
        If matrix[0][j] = 0:
            Set firstRowZero ← true
    For each row index i from 0 to m - 1:
        If matrix[i][0] = 0:
            Set firstColZero ← true

Z2. [Mark rows and columns to be zeroed]
    For each row index i from 1 to m - 1:
        For each column index j from 1 to n - 1:
            If matrix[i][j] = 0:
                Set matrix[i][0] ← 0
                Set matrix[0][j] ← 0

Z3. [Zero marked rows]
    For each row index i from 1 to m - 1:
        If matrix[i][0] = 0:
            For each column index j from 1 to n - 1:
                Set matrix[i][j] ← 0

Z4. [Zero marked columns]
    For each column index j from 1 to n - 1:
        If matrix[0][j] = 0:
            For each row index i from 1 to m - 1:
                Set matrix[i][j] ← 0

Z5. [Zero first row if needed]
    If firstRowZero = true:
        For each column index j from 0 to n - 1:
            Set matrix[0][j] ← 0

Z6. [Zero first column if needed]
    If firstColZero = true:
        For each row index i from 0 to m - 1:
            Set matrix[i][0] ← 0

Z7. [Termination]
    Return the modified matrix

Complexity:
    Time: O(m × n), since every element is visited at most twice.
    Space: O(1), as marking is done within the existing matrix without extra storage.

```


-



> **Game of Life**

```
Algorithm G (Compute the next state of Conway’s Game of Life in-place)

Input:
    - board: m × n integer matrix (0 = dead cell, 1 = live cell)

Output:
    - The board modified in-place to its next state

G1. [Define neighbor directions]  
    Create a list directions containing the 8 possible neighbor coordinate offsets:
        (-1,-1), (-1,0), (-1,1), (0,-1), (0,1), (1,-1), (1,0), (1,1)

G2. [Iterate over the board to compute next state markers]  
    For each cell (i, j) in board:
        G2.1 [Count live neighbors]  
             Initialize liveNeighbors ← 0  
             For each direction (dx, dy) in directions:  
                 Let r ← i + dx, c ← j + dy  
                 If r and c are within bounds and abs(board[r][c]) == 1:  
                     Increment liveNeighbors  
        G2.2 [Mark changes in-place]  
             If board[i][j] == 1 and (liveNeighbors < 2 or liveNeighbors > 3):  
                 Set board[i][j] ← -1    // was live, will become dead  
             If board[i][j] == 0 and liveNeighbors == 3:  
                 Set board[i][j] ← 2     // was dead, will become live

G3. [Finalize state updates]  
    For each cell (i, j) in board:
        If board[i][j] > 0: set board[i][j] ← 1  
        Else: set board[i][j] ← 0

Complexity:
    Time: O(m × n), since each cell and its 8 neighbors are processed once.
    Space: O(1) extra, using the board itself to store transitional states.

```


-



> **Ransom Note**

```
Algorithm C (Check if ransom note can be constructed from magazine)
Input:
    - ransomNote: string of lowercase English letters
    - magazine: string of lowercase English letters
Output:
    - true if ransomNote can be constructed from magazine; false otherwise

C1. [Initialize letter counts]
    Create an array counts[26], initialized to 0
    // counts[i] stores the available count of the (i-th) lowercase letter

C2. [Count magazine letters]
    For each character ch in magazine:
        Let index ← (code of ch) - (code of 'a')    // maps 'a'..'z' to 0..25
        Increment counts[index] by 1

C3. [Check ransom note letters]
    For each character ch in ransomNote:
        Let index ← (code of ch) - (code of 'a')
        Decrement counts[index] by 1
        If counts[index] < 0:
            Return false    // magazine has fewer of this letter than needed

C4. [All letters available]
    Return true

Complexity:
    Time: O(m + n), where m = length of magazine, n = length of ransomNote.
          We scan both strings once.
    Space: O(1), fixed 26-element array regardless of input size.

```


-



> **Isomorphic Strings**

```
Algorithm I (Check if two strings are isomorphic)
Input:
    - s: string of length n (ASCII characters)
    - t: string of length n (ASCII characters)
Output:
    - true if s and t are isomorphic; false otherwise

I1. [Initialize mappings]
    Create an empty map mapST to store mapping from characters in s to characters in t
    Create an empty map mapTS to store mapping from characters in t to characters in s

I2. [Iterate over both strings]
    For i from 0 to n - 1:
        Let chS ← s[i]
        Let chT ← t[i]

        I2.1. [Check s → t mapping]
            If chS exists in mapST:
                If mapST[chS] ≠ chT:
                    Return false
            Else:
                Set mapST[chS] ← chT

        I2.2. [Check t → s mapping]
            If chT exists in mapTS:
                If mapTS[chT] ≠ chS:
                    Return false
            Else:
                Set mapTS[chT] ← chS

I3. [All checks passed]
    Return true

Complexity:
    Time: O(n), where n = length of s (same as t), since each character is processed once.
    Space: O(1) if ASCII fixed-size array is used for mapping; O(k) where k is the number of distinct characters otherwise.

```


-





> **Word Pattern**

```
Algorithm W (Check if a string follows a given pattern with bijection between pattern letters and words)

Input:
- pattern: string containing only lowercase letters
- s: string containing words separated by single spaces

Output:
- true if there exists a bijection between characters in pattern and words in s
- false otherwise

W1. [Split input string]  
    Split s by spaces into list words.  
    If length(words) ≠ length(pattern), return false.

W2. [Initialize mappings]  
    Create empty map charToWord (key: character from pattern, value: corresponding word).  
    Create empty map wordToChar (key: word from s, value: corresponding character from pattern).

W3. [Iterate and validate mapping]  
    For i from 0 to length(pattern) - 1:  
        Let c ← pattern[i], w ← words[i]  
        If c exists in charToWord:  
            If charToWord[c] ≠ w, return false  
        Else:  
            If w exists in wordToChar, return false  
            Set charToWord[c] ← w  
            Set wordToChar[w] ← c

W4. [Return result]  
    Return true

Complexity:  
- Time: O(n), where n is the number of words (same as length of pattern). Each lookup and insert in maps is O(1) on average.  
- Space: O(k), where k ≤ n is the number of unique pattern characters and words stored in maps.

```


-






> **Valid Anagram**

```
Algorithm A' (Check if one string is an anagram of another - arbitrary characters)

Input:
- s: string containing any characters (letters, digits, symbols, spaces, Unicode)
- t: string containing any characters

Output:
- true if t is an anagram of s
- false otherwise

A'.1 [Check lengths]  
     If length(s) ≠ length(t), return false.

A'.2 [Initialize frequency map]  
     Create an empty map freq.

A'.3 [Count characters in s]  
     For each character ch in s:  
         freq[ch] ← freq[ch] + 1

A'.4 [Subtract characters in t]  
     For each character ch in t:  
         If ch not in freq, return false  
         freq[ch] ← freq[ch] - 1  
         If freq[ch] < 0, return false

A'.5 [Return result]  
     Return true (all counts balanced).

Complexity:  
- Time: O(n), where n is the length of s (and t).  
- Space: O(k), where k is the number of distinct characters.

```


-






> **Group Anagrams**

```
Algorithm G (Group anagrams from a list of strings)

Input:
    - strs: array of strings

Output:
    - A list of groups, where each group contains strings that are anagrams of each other

G1. [Initialize data structure]
    Create an empty hash map groups, mapping from a "canonical form" to a list of strings.

G2. [Iterate over each string]
    For each string word in strs:
        G2.1. [Compute canonical form]
              Convert word into a character array chars.
              Sort chars in ascending order.
              Convert sorted chars back into a string key.
        G2.2. [Add to group]
              Append word to groups[key].

G3. [Collect results]
    Return the list of values from groups as the grouped anagrams.

Complexity:
    Let n be the number of strings and k be the maximum length of a string.
    Time: O(n·k·log k), sorting each string takes O(k·log k).
    Space: O(n·k), to store the hash map and the sorted keys.

```


-






> **Two Sum**

```
Algorithm T (Find two indices whose values sum to a target)

Input:
    - nums: an array of integers
    - target: an integer

Output:
    - A pair [i, j] such that nums[i] + nums[j] = target and i ≠ j

T1. [Initialize map]
    Create an empty map M to store (value → index) pairs.

T2. [Iterate over nums]
    For i from 0 to length(nums) - 1:
        Let complement ← target - nums[i]

T3. [Check if complement exists]
        If complement is in M:
            Return [M[complement], i]

T4. [Store current value in map]
        Set M[nums[i]] ← i

T5. [No solution found]
    (Given constraints guarantee exactly one solution, so this step is never reached.)

Complexity:
    Time: O(n), where n is the number of elements in nums. Each lookup and insert in M takes O(1) on average.
    Space: O(n), for storing at most n elements in the map.

```


-






> **Happy Number**

```
Algorithm H (Check if a number is happy)

Input:
    - n: a positive integer

Output:
    - true if n is a happy number, false otherwise

H1. [Initialize set]
    Create an empty set S to store numbers already seen.

H2. [Iterate]
    While n is not 1 and n is not in S:
        Insert n into S.
        Set n ← SumOfSquares(n).

H3. [Check result]
    If n = 1:
        Return true.
    Else:
        Return false.

---

Algorithm SumOfSquares (Compute sum of squares of digits)

Input:
    - num: a positive integer

Output:
    - sum: integer sum of the squares of digits of num

SS1. [Initialize sum]
    sum ← 0.

SS2. [Process digits]
    While num > 0:
        digit ← num mod 10
        sum ← sum + digit × digit
        num ← num ÷ 10

SS3. [Return result]
    Return sum.

---

Complexity:
    Time: O(k) per iteration, where k is the number of digits in n; the number of iterations is bounded because values repeat or reach 1, so overall O(1) with a small constant.
    Space: O(m), where m is the number of unique values before repetition (bounded by possible sums of squares of digits).

```


-






> **Contains Duplicate II**

```
Algorithm C (Check for nearby duplicate elements)

Input:
    - nums: array of integers
    - k: integer distance threshold

Output:
    - true if there exist distinct indices i and j such that nums[i] = nums[j] and |i - j| ≤ k
    - false otherwise

C1. [Initialize map]
    Create an empty map M to store element → last seen index.

C2. [Iterate over array]
    For i from 0 to len(nums) - 1:
        If nums[i] exists in M and (i - M[nums[i]]) ≤ k:
            Return true
        Set M[nums[i]] ← i

C3. [Return false]
    Return false.

---

Complexity:
    Time: O(n), where n = length of nums (single pass).
    Space: O(min(n, k)), since at most k elements are stored in the map.

```


-






> **Longest Consecutive Sequence**

```
Algorithm L (Longest consecutive sequence in O(n) time)

Input:
    - nums: array of integers

Output:
    - length of the longest sequence of consecutive integers

L1. [Handle empty input]
    If nums is empty, return 0.

L2. [Insert into set]
    Create a hash set S containing all elements of nums.

L3. [Initialize maximum length]
    Set longest ← 0.

L4. [Iterate through numbers]
    For each num in S:
        If (num - 1) is not in S:       // num is the start of a sequence
            current ← num
            length ← 1
            While (current + 1) is in S:
                current ← current + 1
                length ← length + 1
            longest ← max(longest, length)

L5. [Return result]
    Return longest.

---

Complexity:
    Time: O(n), because each number is visited at most twice.
    Space: O(n) for the hash set.

```


-






> ****

```
Algorithm R (Summary Ranges)

Input:
    nums — a sorted array of unique integers.

Output:
    A list of strings representing ranges, where each string is either:
        "a"     if the range contains a single number,
        "a->b"  if the range covers consecutive integers from a to b.

R1: Initialize result container
    Let result ← empty list.
    If nums is empty, return result.

R2: Start first range
    Let start ← nums[0].
    Let end ← nums[0].

R3: Traverse array
    For i from 1 to length(nums) - 1 do:
        If nums[i] = end + 1:
            Update end ← nums[i].
        Else:
            If start = end:
                Append string(start) to result.
            Else:
                Append string(start) + "->" + string(end) to result.
            Reset start ← nums[i], end ← nums[i].

R4: Finalize last range
    After loop ends:
        If start = end:
            Append string(start) to result.
        Else:
            Append string(start) + "->" + string(end) to result.

R5: Return result
    Return result.

Complexity:
    Time:  O(n), where n = length(nums).
    Space: O(1) extra (excluding result storage).

```


-






> **Merge Intervals**

```
Algorithm M (Merge Intervals)

Input:
    intervals — an array of intervals [start, end].

Output:
    A merged array of non-overlapping intervals covering all input intervals.

M1: Handle empty input
    If intervals is empty, return empty list.

M2: Sort intervals
    Sort intervals by start coordinate in ascending order.

M3: Initialize result
    Let result ← empty list.
    Let current ← intervals[0].

M4: Traverse and merge
    For each interval in intervals[1:]:
        If interval.start ≤ current.end:
            // overlap found
            current.end ← max(current.end, interval.end)
        Else:
            Append current to result.
            current ← interval.

M5: Append last interval
    Append current to result.

M6: Return result
    Return result.

Complexity:
    Time:  O(n log n), due to sorting n intervals.
    Space: O(n) for the result (O(1) extra beyond output).

```


-






> **Insert Interval**

```
Algorithm I (Insert Interval)

Input:
    intervals — sorted array of non-overlapping intervals [start, end].
    newInterval — an interval [start, end].

Output:
    An array of intervals after inserting newInterval, still sorted and non-overlapping.

I1: Initialize result
    Let result ← empty list.

I2: Traverse intervals before overlap
    While intervals is not empty AND intervals[0].end < newInterval.start:
        Append intervals[0] to result.
        Remove intervals[0] from intervals.

I3: Merge overlaps with newInterval
    While intervals is not empty AND intervals[0].start ≤ newInterval.end:
        newInterval.start ← min(newInterval.start, intervals[0].start)
        newInterval.end   ← max(newInterval.end,   intervals[0].end)
        Remove intervals[0] from intervals.

I4: Append merged newInterval
    Append newInterval to result.

I5: Append remaining intervals
    Append all remaining intervals to result.

I6: Return result
    Return result.

Complexity:
    Time:  O(n), each interval processed once.
    Space: O(n), for the result.

```


-






> **Minimum Number of Arrows to Burst Balloons**

```
Algorithm B (Minimum Number of Arrows to Burst Balloons)

Input:
    points — a list of intervals [x_start, x_end] representing balloon diameters.

Output:
    An integer representing the minimum number of arrows required to burst all balloons.

B.1: Handle empty input
    If points is empty, return 0.

B.2: Sort intervals
    Sort points by their x_end values in ascending order.

B.3: Initialize variables
    arrows ← 1
    arrowPos ← points[0].x_end   // place first arrow at the end of the first interval

B.4: Traverse intervals
    For each interval [start, end] in points starting from index 1:
        If start > arrowPos:
            arrows ← arrows + 1
            arrowPos ← end

B.5: Return result
    Return arrows.

Time Complexity:
    O(n log n) due to sorting, where n = number of intervals.
Space Complexity:
    O(1) extra space (ignoring sort cost).

```


-






> **Number of Islands**

```
Algorithm I (Count number of islands in a 2D grid using DFS exploration)

Input:
    - grid: m × n matrix of characters ('1' for land, '0' for water)

Output:
    - Integer count representing the number of islands

I1. [Initialize island counter]
    Set count ← 0

I2. [Iterate over grid]
    For each row i from 0 to m - 1:
        For each column j from 0 to n - 1:
            If grid[i][j] = '1':
                Increment count ← count + 1
                Call SinkIsland(i, j, grid)

I3. [Return result]
    Return count


Algorithm SinkIsland (Mark entire island starting at (i, j) as visited)

Input:
    - i, j: coordinates of the current land cell
    - grid: m × n matrix

Output:
    - None (grid is modified in-place)

S1. [Check boundaries]
    If i < 0 or j < 0 or i ≥ m or j ≥ n:
        Return

S2. [Check if water or already visited]
    If grid[i][j] ≠ '1':
        Return

S3. [Mark current cell as visited]
    Set grid[i][j] ← '0'

S4. [Explore neighbors]
    Call SinkIsland(i+1, j, grid)
    Call SinkIsland(i-1, j, grid)
    Call SinkIsland(i, j+1, grid)
    Call SinkIsland(i, j-1, grid)


Complexity:
    Time: O(m × n), where m is rows and n is columns.
        Each cell is visited at most once.
    Space: O(m × n) in worst case due to recursive call stack (all land).

```


-







> **Surrounded Regions**

```
Algorithm C (Capture surrounded regions in a 2D board of 'X' and 'O')

Input:
    - board: m × n matrix of characters, where each cell is either 'X' or 'O'

Output:
    - The board modified in-place such that all 'O' regions fully surrounded by 'X' are converted to 'X'

C1. [Handle edge case]
    If board is empty:
        Return

C2. [Mark boundary-connected 'O']
    For each cell (i, j) on the boundary of board:
        If board[i][j] = 'O':
            Run DFS-Mark(i, j) to mark all connected 'O's as temporary marker (say 'T')

C3. [Convert surrounded 'O's]
    For each cell (i, j) in the board:
        If board[i][j] = 'O':
            Set board[i][j] ← 'X'   // These are surrounded and must be captured

C4. [Restore safe 'O's]
    For each cell (i, j) in the board:
        If board[i][j] = 'T':
            Set board[i][j] ← 'O'   // Restore boundary-connected regions


Helper Algorithm DFS-Mark(i, j) (Mark all 'O's connected to (i, j))
Input:
    - i, j: coordinates of a cell
Effect:
    - Marks connected 'O' cells as 'T'

M1. [Mark cell]
    Set board[i][j] ← 'T'

M2. [Recurse on neighbors]
    For each neighbor (ni, nj) of (i, j) in {up, down, left, right}:
        If (ni, nj) is within bounds and board[ni][nj] = 'O':
            Call DFS-Mark(ni, nj)


Complexity:
    Time: O(m × n), since each cell is visited at most once during marking and conversion.
    Space: O(m × n) in worst case recursion depth (if all cells are 'O'), though iterative stack or BFS can reduce risk.

```


-







> **Clone Graph**

```
Algorithm C (Clone an undirected connected graph using DFS)

Input:
    - node: reference to a node in a connected undirected graph
        Each node contains:
            - val: integer identifier
            - neighbors: list of adjacent nodes

Output:
    - A deep copy (clone) of the graph, starting from the given node reference

C1. [Handle empty input]  
    If node = nil:  
        Return nil  

C2. [Initialize data structures]  
    Create a hash map `visited` to store mapping (original node → cloned node).  

C3. [Define recursive DFS clone procedure]  
    Procedure DFS(current):  
        If current ∈ visited:  
            Return visited[current]  

        Create clone ← new Node(current.Val, empty neighbors list).  
        Store visited[current] ← clone.  

        For each neighbor in current.Neighbors:  
            Append DFS(neighbor) to clone.Neighbors.  

        Return clone  

C4. [Start DFS cloning]  
    Return DFS(node)  


Complexity:  
    Time: O(V + E), where V is the number of nodes and E is the number of edges.  
        Each node and edge is visited once during the DFS.  
    Space: O(V), for recursion stack (worst case in deep graph) and the hash map storing clones.  

```


-







> **Evaluate Division**

```
Algorithm E (Evaluate Division using graph traversal)

Input:
- equations: list of pairs [Ai, Bi] representing equations
- values: list of floats such that Ai / Bi = values[i]
- queries: list of pairs [Cj, Dj] to evaluate

Output:
- List of floats representing answers for each query; -1.0 if not computable

E1. [Build graph]  
    Create adjacency list graph G mapping string → list of (neighbor, weight).  
    For each equation (Ai, Bi) with value v:  
        Add edge Ai → Bi with weight v  
        Add edge Bi → Ai with weight 1/v  

E2. [Initialize result list]  
    Create empty list results  

E3. [Process each query]  
    For each query (C, D):  
        If C or D not in G: Append -1.0 to results and continue  
        If C = D: Append 1.0 to results and continue  

E4. [Graph traversal to compute ratio]  
    Perform DFS from C to D:  
        Maintain visited set  
        DFS(node, target, accum):  
            If node = target: return accum  
            For each (neighbor, weight) of node:  
                If neighbor not visited:  
                    result ← DFS(neighbor, target, accum × weight)  
                    If result ≠ -1.0: return result  
            Return -1.0  

E5. [Store query result]  
    Run DFS(C, D, 1.0) and append result to results  

E6. [Return results]  
    Return results list

Complexity:  
- Time: O(Q × (V + E)), where Q = number of queries, V = number of variables, E = number of equations. Each query may explore the entire graph in worst case.  
- Space: O(V + E) for the graph, O(V) recursion stack for DFS.  

```


-







> **Course Schedule**

```
Algorithm C (Check if all courses can be finished using prerequisites)  
Input:  
- numCourses: integer, total number of courses labeled 0 to numCourses-1  
- prerequisites: list of pairs [a, b], meaning course b must be taken before course a  

Output:  
- true if all courses can be finished, false otherwise  

C1. [Initialize graph structures]  
    Create adjacency list graph of size numCourses, initially empty  
    Create inDegree array of size numCourses, initialized to 0  

C2. [Build graph]  
    For each pair [a, b] in prerequisites:  
        Append a to graph[b]  
        Increment inDegree[a] by 1  

C3. [Initialize queue with zero in-degree courses]  
    Create empty queue  
    For each course i in [0..numCourses-1]:  
        If inDegree[i] = 0: enqueue(i)  

C4. [Process courses using BFS topological sort]  
    Initialize count ← 0  
    While queue not empty:  
        Dequeue course curr  
        Increment count by 1  
        For each neighbor next in graph[curr]:  
            Decrement inDegree[next] by 1  
            If inDegree[next] = 0: enqueue(next)  

C5. [Check feasibility]  
    If count = numCourses: return true  
    Else: return false  

Complexity:  
- Time: O(V + E), where V = numCourses and E = number of prerequisites. Each edge and vertex is processed once.  
- Space: O(V + E) for adjacency list and inDegree array.  

```


-






> **Course Schedule II**

```
Algorithm O (Find a valid order to finish all courses using topological sort)

Input:
- numCourses: integer, number of courses labeled 0 … numCourses-1
- prerequisites: list of pairs [a, b], meaning b → a (must take b before a)

Output:
- A list of courses in a valid order, or [] if impossible

O1. [Initialize structures]
    Create adjacency list graph for numCourses.
    Create indegree array indeg[0..numCourses-1] initialized to 0.
    For each (a, b) in prerequisites:
        Append a to graph[b].
        Increment indeg[a].

O2. [Find starting courses]
    Create queue Q.
    For i from 0 to numCourses-1:
        If indeg[i] = 0, enqueue(Q, i).

O3. [Process queue]
    Initialize result list order ← [].
    While Q not empty:
        Dequeue u from Q.
        Append u to order.
        For each v in graph[u]:
            Decrement indeg[v].
            If indeg[v] = 0, enqueue(Q, v).

O4. [Check completion]
    If length(order) = numCourses:
        Return order
    Else:
        Return []   // cycle exists

Complexity:
- Time: O(V + E), where V = numCourses, E = number of prerequisites.
- Space: O(V + E) for graph, indegree, and queue.

```


-







> **Snakes and Ladders**

```
Algorithm S (Minimum dice rolls to reach the end in Snakes and Ladders)

Input:
- board: n × n matrix with values -1 (empty) or destination square number (snake/ladder)
Output:
- Minimum number of dice rolls to reach square n², or -1 if unreachable

S1. [Flatten board mapping]
    Let n ← length of board
    Create array moves[1..n²] initialized to -1
    Set idx ← 1
    For row from n-1 down to 0:
        If (n-1 - row) is even:
            For col from 0 to n-1:
                If board[row][col] ≠ -1: moves[idx] ← board[row][col]
                idx ← idx + 1
        Else:
            For col from n-1 down to 0:
                If board[row][col] ≠ -1: moves[idx] ← board[row][col]
                idx ← idx + 1

S2. [Initialize BFS]
    Create queue Q and enqueue (1, 0)   // (square, rolls)
    Create visited set with 1

S3. [BFS traversal]
    While Q not empty:
        Dequeue (curr, rolls)
        If curr = n²: return rolls
        For step from 1 to 6:
            next ← curr + step
            If next > n²: break
            If moves[next] ≠ -1: next ← moves[next]
            If next not in visited:
                Add next to visited
                Enqueue (next, rolls + 1)

S4. [No path]
    Return -1

Complexity:
- Time: O(n²), since we process at most n² board cells and each has up to 6 edges.
- Space: O(n²), for moves array and visited set.

```


-







> **Minimum Genetic Mutation**

```
Algorithm M (Minimum genetic mutations using BFS)  
Input:  
- startGene: string of length 8 representing the starting gene  
- endGene: string of length 8 representing the target gene  
- bank: list of strings of length 8 representing valid mutations  

Output:  
- Minimum number of mutations from startGene to endGene, or -1 if impossible  

M1. [Convert bank to set]  
    Let valid ← set(bank).  
    If endGene ∉ valid: return -1.  

M2. [Initialize BFS]  
    Create queue Q and enqueue (startGene, 0).  
    Create visited set V and add startGene.  

M3. [Process queue]  
    While Q is not empty:  
        (gene, steps) ← dequeue(Q).  

        If gene = endGene: return steps.  

        For each position i in 0..7:  
            For each char in ['A','C','G','T']:  
                If char ≠ gene[i]:  
                    newGene ← gene with gene[i] replaced by char.  

                    If newGene ∈ valid and newGene ∉ V:  
                        Add newGene to V.  
                        Enqueue (newGene, steps + 1).  

M4. [If end not reached]  
    Return -1.  

Complexity:  
- Time: O(N * L * Σ), where N = |bank|, L = 8 (gene length), Σ = 4 (alphabet size).  
  In practice O(N * 32).  
- Space: O(N) for visited set and queue.  

```


-





> ****

```
Algorithm L (Shortest Word Transformation Sequence using BFS)

Input:
- beginWord: starting word (string)
- endWord: target word (string)
- wordList: list of valid dictionary words (array of strings)

Output:
- Minimum number of words in the transformation sequence from beginWord to endWord,
  or 0 if no such sequence exists.

L1. [Preprocess word list]  
    Convert wordList into a set wordSet for O(1) membership lookup.  
    If endWord ∉ wordSet: Return 0.

L2. [Initialize BFS]  
    Create a queue Q and enqueue the pair (beginWord, 1), where 1 is the current path length.  
    Create a visited set and insert beginWord.

L3. [BFS traversal]  
    While Q is not empty:  
        Dequeue (word, steps).  
        If word = endWord: Return steps.  

L4. [Generate neighbors]  
    For each position i from 0 to len(word) - 1:  
        For each character c from 'a' to 'z':  
            If c ≠ word[i]:  
                newWord ← word with word[i] replaced by c.  
                If newWord ∈ wordSet and newWord ∉ visited:  
                    Enqueue (newWord, steps + 1).  
                    Insert newWord into visited.

L5. [No path found]  
    If BFS completes without finding endWord: Return 0.

Complexity:  
- Time: O(L × 26 × N), where L = word length, N = number of words in wordList.  
  For each word dequeued, we generate up to 26L possible neighbors and check set membership.  
- Space: O(N), for storing the dictionary set, visited set, and queue.

```


-







> **Implement Trie (Prefix Tree)**

```
Algorithm T (Trie with insert, search, and prefix check)

Input:
- Operations: insert(word), search(word), startsWith(prefix)
- Words/prefixes consist of lowercase English letters

Output:
- Trie structure supporting word insertion, full word search, and prefix existence check

T1. [Define Trie node structure]
    Each node contains:
    - children: map[byte]*TrieNode (mapping character → next node)
    - isEnd: boolean flag (true if node marks the end of a word)

T2. [Constructor]
    Create a new Trie with a root node (empty, not an end of word).

T3. [Insert(word)]
    Set current ← root
    For each character c in word:
        If current.children[c] does not exist:
            Create new TrieNode and assign to current.children[c]
        Move current ← current.children[c]
    After loop, set current.isEnd ← true

T4. [Search(word)]
    Set current ← root
    For each character c in word:
        If current.children[c] does not exist:
            Return false
        Move current ← current.children[c]
    Return current.isEnd

T5. [StartsWith(prefix)]
    Set current ← root
    For each character c in prefix:
        If current.children[c] does not exist:
            Return false
        Move current ← current.children[c]
    Return true

T6. [Delete(word)]
    Goal: remove word without breaking other words that share its prefix.

    T6.1: Define recursive helper DeleteHelper(node, word, depth) → bool
        Input: current TrieNode, word, and current depth
        Output: bool (true if this node should be deleted by parent)

        T6.1.1: [Base case: end of word]
            If depth == len(word):
                If node.isEnd == false:
                    Return false   // word does not exist
                Set node.isEnd = false
                If node has no children:
                    Return true    // safe to delete this node
                Return false       // cannot delete because children exist

        T6.1.2: [Recursive case]
            c ← word[depth]
            If node.children[c] does not exist:
                Return false   // word does not exist
            childShouldDelete ← DeleteHelper(node.children[c], word, depth+1)
            If childShouldDelete:
                delete(node.children, c)
                If node has no children and node.isEnd == false:
                    Return true
            Return false

    T6.2: Call DeleteHelper(root, word, 0)

Complexity:
- Insert: O(L)
- Search: O(L)
- StartsWith: O(L)
- Delete: O(L)
  where L is the length of the word/prefix
- Space: O(N * Σ) for N words and alphabet size Σ


examples

1. 
add the below to a new trie
"international"
"internet"
"interval"
"into"

so the trie becomes
root
 └─ i
    └─ n
       └─ t
          └─ e
             └─ r
                └─ n
                   └─ a
                      └─ t
                         └─ i
                            └─ o
                               └─ n
                                  └─ a
                                     └─ l   [end of "international"]
                └─ e
                   └─ t   [end of "internet"]
                └─ v
                   └─ a
                      └─ l   [end of "interval"]
          └─ o
             └─   [end of "into"]


2. 
insert("apple") → creates /a/p/p/l/e path and marks e as end. 
search("apple") → checks if /a/p/p/l/e exists and is marked. 
startsWith("app") → checks if /a/p/p path exists, regardless of whether it ends a word.
delete("apple") → unmarks /a/p/p/l/e as end, and prunes nodes upward if they’re no longer needed.

```


-




> **Design Add and Search Words Data Structure**

```
Algorithm W (Word Dictionary with Wildcard Search)
Input: 
    - Operations: addWord(word), search(word)
    - word may contain lowercase letters and '.' where '.' matches any letter
Output:
    - Boolean result for search queries

W.1: Constructor
    Create a root node:
        - children ← empty dictionary (key = character, value = node)
        - isEnd ← false

W.2: AddWord(word)
    node ← root
    For each character ch in word:
        If ch not in node.children:
            node.children[ch] ← new node with empty children and isEnd = false
        node ← node.children[ch]
    node.isEnd ← true

W.3: Search(word)
    Return DFS(word, idx = 0, root)

Algorithm DFS (Recursive Search with Wildcard)
Input:
    - word: string
    - idx: current index in word
    - node: current Trie node
Output:
    - Boolean indicating whether word can be matched

DFS.1: If idx == length(word):
           Return node.isEnd

DFS.2: ch ← word[idx]

DFS.3: If ch == '.':
           For each child in node.children:
               If DFS(word, idx + 1, child) = true:
                   Return true
           Return false

DFS.4: Else:
           If ch not in node.children:
               Return false
           Else:
               Return DFS(word, idx + 1, node.children[ch])

Complexity:
    - AddWord: O(L), where L = length of word
    - Search (worst case with '.' at every position): O(26^L) 
      but usually prunes quickly in practice
    - Space: O(N * L) for N words of average length L


```


-






> **Word Search II**

```
Algorithm W (Find all words in board using Trie + DFS)

Input:
- board: m × n grid of characters
- words: list of strings
Output:
- list of words from `words` that can be formed on the board

W1. [Build Trie]
    Initialize root Trie node with empty children and word field = ""
    For each word in words:
        node ← root
        For each character ch in word:
            If node.children[ch] does not exist:
                create new Trie node
            node ← node.children[ch]
        At end of word: store word in node.word (marks terminal)

W2. [DFS Search Initialization]
    Initialize result set = ∅
    For i = 0 to m-1:
        For j = 0 to n-1:
            Call DFS(i, j, root)

Algorithm DFS (Explore board from cell (i, j) with current Trie node)
Input:
- i, j: current coordinates
- node: current Trie node
Output:
- Marks found words in result set

DFS1. [Bounds and pruning]
    If (i, j) is out of bounds or board[i][j] = '#' (visited marker) or board[i][j] not in node.children:
        Return

DFS2. [Advance in Trie]
    ch ← board[i][j]
    nextNode ← node.children[ch]

DFS3. [Word found]
    If nextNode.word ≠ "":
        Add nextNode.word to result set
        Set nextNode.word = ""   // prevent duplicate results

DFS4. [Mark visited]
    Set board[i][j] = '#'

DFS5. [Explore neighbors]
    For each direction (up, down, left, right):
        DFS(new_i, new_j, nextNode)

DFS6. [Unmark visited]
    Restore board[i][j] = ch

W3. [Return result]
    Convert result set to list and return

Complexity:
- Building Trie: O(ΣL), where ΣL is total length of all words
- DFS traversal: O(m × n × 4^L), but with Trie pruning, worst-case bounded by O(m × n × maxWordLength)
- Space: O(ΣL) for Trie + O(L) recursion stack

```


-







> **Letter Combinations of a Phone Number**

```
Algorithm L (Generate all possible letter combinations from digits using backtracking)

Input:
- digits: string of digits from '2' to '9'
Output:
- list of all possible letter combinations

L1. [Handle edge case]
    If digits is empty: return empty list

L2. [Mapping initialization]
    Define digitToLetters map:
        '2' → "abc", '3' → "def", '4' → "ghi", '5' → "jkl",
        '6' → "mno", '7' → "pqrs", '8' → "tuv", '9' → "wxyz"

L3. [Initialize result]
    result ← empty list

L4. [Recursive backtracking]
    Call Backtrack(idx=0, path="")

Algorithm Backtrack (Recursive construction of combinations)
Input:
- idx: current index in digits
- path: partial combination built so far
Output:
- Append valid combinations to result

B1. [Base case]
    If idx == len(digits):
        Append path to result
        Return

B2. [Expand current digit]
    currentDigit ← digits[idx]
    For each letter in digitToLetters[currentDigit]:
        Call Backtrack(idx+1, path+letter)

L5. [Return result]
    Return result

Complexity:
- Time: O(4^n), where n = length of digits (each digit maps to up to 4 letters, and we explore all combinations).
- Space: O(n) recursion depth + O(4^n) output storage.

```


-




> **Combinations**

```
Algorithm C (Generate all k-combinations from numbers 1..n using backtracking)

Input:
- n: integer, upper bound of the range [1..n]
- k: integer, size of each combination
Output:
- list of all possible k-combinations

C1. [Initialize result]
    result ← empty list

C2. [Recursive backtracking]
    Call Backtrack(start=1, path=[])

Algorithm Backtrack (Recursive construction of combinations)
Input:
- start: current number to consider in the range [1..n]
- path: partial combination built so far
Output:
- Append valid combinations to result

B1. [Base case: full combination]
    If length(path) == k:
        Append copy of path to result
        Return

B2. [Recursive expansion]
    For num from start to n:
        Append num to path
        Call Backtrack(num+1, path)
        Remove last element from path   // backtrack

C3. [Return result]
    Return result

Complexity:
- Time: O(C(n,k) * k), since there are C(n,k) combinations and each requires O(k) to build/append.
- Space: O(k) recursion depth (excluding output).

```


-


> **Permutations**

```
Algorithm P (Generate all permutations of distinct integers using backtracking)

Input:
- nums: list of distinct integers
Output:
- list of all possible permutations

P1. [Initialize result]
    result ← empty list

P2. [Recursive backtracking]
    Call Backtrack(path=[], used[0..len(nums)-1] initialized to false)

Algorithm Backtrack (Recursive construction of permutations)
Input:
- path: current partial permutation
- used: boolean array marking which elements are already in path
Output:
- Append valid permutations to result

B1. [Base case: full permutation]
    If length(path) == length(nums):
        Append copy of path to result
        Return

B2. [Recursive expansion]
    For i from 0 to len(nums)-1:
        If used[i] == false:
            Mark used[i] = true
            Append nums[i] to path
            Call Backtrack(path, used)
            Remove last element from path   // backtrack
            Mark used[i] = false

P3. [Return result]
    Return result

Complexity:
- Time: O(n × n!), because there are n! permutations and copying each path of size n costs O(n).
- Space: O(n) recursion depth + O(n) used array (excluding output).

```


-



> **Combination Sum**

```
Algorithm C (Backtracking with sorting + pruning for Combination Sum)
Input:
    - candidates: array of distinct positive integers
    - target: integer, the required sum
Output:
    - A list of unique combinations (each combination is a list of integers) that sum to target

C1. [Sort candidates for pruning]
    Sort candidates in non-decreasing order.

C2. [Prepare result container]
    Let result ← empty list

C3. [Define backtracking function]
    Define Backtrack(start, comb, remain):
        // Invariant: every value in comb is >= candidates[start-1] (comb is non-decreasing)
        If remain = 0:
            Append a copy of comb to result
            Return
        For i from start to len(candidates) - 1:
            If candidates[i] > remain:
                Break    // no later candidate can fit because array is sorted
            Append candidates[i] to comb
            Call Backtrack(i, comb, remain - candidates[i])   // allow reuse of candidates[i]
            Remove last element from comb (backtrack)

C4. [Kick off recursion]
    Call Backtrack(0, [], target)

C5. [Return]
    Return result

Complexity:
    Time: Exponential in the worst case. A useful upper-bound is O(N^(T/min)), where
        N = number of candidates, T = target, min = smallest candidate.
        Sorting adds O(N log N). Pruning (candidates[i] > remain) reduces the explored space when candidates are positive.
    Space: O(T/min) recursion depth (call stack) plus O(k * L) for the output,
        where k = number of solutions and L = average length of a solution.



```


-





> **N-Queens II**

```
Algorithm N (Backtracking with pruning for N-Queens count)
Input:
    - n: size of the chessboard (n x n)
Output:
    - The number of distinct ways to place n queens so that none attack each other

N1. [Initialize bookkeeping sets]
    Create three sets (or boolean arrays) to track conflicts:
        - cols: columns under attack
        - diag1: "\" diagonals under attack, indexed by (row - col)
        - diag2: "/" diagonals under attack, indexed by (row + col)
    Let count ← 0

N2. [Define backtracking function]
    Define Backtrack(row):
        If row = n:
            Increment count by 1
            Return
        For col from 0 to n - 1:
            If col ∈ cols or (row - col) ∈ diag1 or (row + col) ∈ diag2:
                Continue
            Insert col into cols
            Insert (row - col) into diag1
            Insert (row + col) into diag2
            Call Backtrack(row + 1)
            Remove col from cols
            Remove (row - col) from diag1
            Remove (row + col) from diag2

N3. [Kick off recursion]
    Call Backtrack(0)

N4. [Return result]
    Return count

Complexity:
    Time: O(n!), since in the worst case each row tries many positions, but pruning cuts this significantly.
        The actual search tree size is far smaller than n! for typical n.
    Space: O(n) for recursion depth and bookkeeping arrays.

```


-



> ****

```

```


-




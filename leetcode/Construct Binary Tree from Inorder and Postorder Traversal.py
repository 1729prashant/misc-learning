"""
106. Construct Binary Tree from Inorder and Postorder Traversal

Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.

 

Example 1:


Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
Output: [3,9,20,null,null,15,7]
Example 2:

Input: inorder = [-1], postorder = [-1]
Output: [-1]
 

Constraints:

1 <= inorder.length <= 3000
postorder.length == inorder.length
-3000 <= inorder[i], postorder[i] <= 3000
inorder and postorder consist of unique values.
Each value of postorder also appears in inorder.
inorder is guaranteed to be the inorder traversal of the tree.
postorder is guaranteed to be the postorder traversal of the tree.

Algorithm Explanation:
1. Root Identification: The last element in the postorder traversal is always the root of the tree.
2. Subtree Division: Find this root in the inorder traversal. Everything to the left of this root in inorder corresponds to the left subtree, and everything to the right corresponds to the right subtree.
3. Recursive Construction:
  a. Use the elements from the start of postorder up to (but not including) the position of the root in inorder for the left subtree.
  b. Use the elements from after this position in postorder up to the current root for the right subtree.

"""
from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def buildTree(self, inorder: list[int], postorder: list[int]) -> Optional[TreeNode]:
        if not inorder or not postorder:
            return None
        
        # The last element in postorder is the root
        root = TreeNode(postorder[-1])
        
        # Find the position of the root in inorder to determine the size of left subtree
        mid = inorder.index(postorder[-1])
        
        # Recursively build left and right subtrees
        root.left = self.buildTree(inorder[:mid], postorder[:mid])
        root.right = self.buildTree(inorder[mid+1:], postorder[mid:-1])
        
        return root
    
solution = Solution()

# Helper function to convert tree to list (for assertion)
def tree_to_list(root):
    if not root:
        return []
    result = []
    queue = [root]
    while queue:
        node = queue.pop(0)
        if node:
            result.append(node.val)
            queue.append(node.left if node.left else None)
            queue.append(node.right if node.right else None)
        else:
            result.append(None)
    # Remove trailing None values for simplicity
    while result and result[-1] is None:
        result.pop()
    return result

# Test Case 1
inorder1 = [9, 3, 15, 20, 7]
postorder1 = [9, 15, 7, 20, 3]
tree1 = solution.buildTree(inorder1, postorder1)
assert tree_to_list(tree1) == [3, 9, 20, None, None, 15, 7]

# Test Case 2 - Single Node
inorder2 = [-1]
postorder2 = [-1]
tree2 = solution.buildTree(inorder2, postorder2)
assert tree_to_list(tree2) == [-1]

# Test Case 3 - More complex tree
inorder3 = [2, 1, 3]
postorder3 = [2, 3, 1]
tree3 = solution.buildTree(inorder3, postorder3)
assert tree_to_list(tree3) == [1, 2, 3]

# Test Case 4 - Empty Tree
inorder4 = []
postorder4 = []
tree4 = solution.buildTree(inorder4, postorder4)
assert tree_to_list(tree4) == []
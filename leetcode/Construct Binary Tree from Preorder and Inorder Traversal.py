"""
105. Construct Binary Tree from Preorder and Inorder Traversal

Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

 

Example 1:


Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
Example 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]
 

Constraints:

1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder and inorder consist of unique values.
Each value of inorder also appears in preorder.
preorder is guaranteed to be the preorder traversal of the tree.
inorder is guaranteed to be the inorder traversal of the tree.



Algorithm Explanation:
1. Root Identification: The first element in the preorder traversal is always the root of the tree.
2. Subtree Division: Find this root in the inorder traversal. Everything to the left of this root in inorder corresponds to the left subtree, and everything to the right corresponds to the right subtree.
3. Recursive Construction:
  a. Use the elements from the start of preorder up to (but not including) the current root for the left subtree.
  b. Use the elements from after the current root in preorder for the right subtree.
"""


from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def buildTree(self, preorder: list[int], inorder: list[int]) -> Optional[TreeNode]:
        if not preorder or not inorder:
            return None
        
        # The first element in preorder is the root
        root = TreeNode(preorder[0])
        
        # Find the position of the root in inorder to determine the size of left subtree
        mid = inorder.index(preorder[0])
        
        # Recursively build left and right subtrees
        root.left = self.buildTree(preorder[1:mid+1], inorder[:mid])
        root.right = self.buildTree(preorder[mid+1:], inorder[mid+1:])
        
        return root


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

# Note: The assertions above would be used in actual test scenarios but aren't executable here.

solution = Solution()

# Test Case 1
preorder1 = [3,9,20,15,7]
inorder1 = [9,3,15,20,7]
tree1 = solution.buildTree(preorder1, inorder1)
# You would need a helper function to convert tree to list for assertion, here's a conceptual one:
assert tree_to_list(tree1) == [3,9,20,None,None,15,7]

# Test Case 2
preorder2 = [-1]
inorder2 = [-1]
tree2 = solution.buildTree(preorder2, inorder2)
assert tree_to_list(tree2) == [-1]

# Test Case 3 - More complex tree
preorder3 = [1,2,4,5,3,6,7]
inorder3 = [4,2,5,1,6,3,7]
tree3 = solution.buildTree(preorder3, inorder3)
assert tree_to_list(tree3) == [1,2,3,4,5,6,7]


"""
104. Maximum Depth of Binary Tree

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

 

Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: 3
Example 2:

Input: root = [1,null,2]
Output: 2
 

Constraints:

The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100
"""

import unittest
from typing import Optional

# Assume TreeNode is defined elsewhere or within the test file for completeness
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        left_depth = self.maxDepth(root.left)
        right_depth = self.maxDepth(root.right)
        return max(left_depth, right_depth) + 1

class TestMaxDepth(unittest.TestCase):

    def test_max_depth_example1(self):
        # Example 1: [3,9,20,null,null,15,7]
        root = TreeNode(3)
        root.left = TreeNode(9)
        root.right = TreeNode(20)
        root.right.left = TreeNode(15)
        root.right.right = TreeNode(7)
        
        solution = Solution()
        self.assertEqual(solution.maxDepth(root), 3)

    def test_max_depth_example2(self):
        # Example 2: [1,null,2]
        root = TreeNode(1)
        root.right = TreeNode(2)
        
        solution = Solution()
        self.assertEqual(solution.maxDepth(root), 2)

    def test_max_depth_empty_tree(self):
        # Empty tree case
        solution = Solution()
        self.assertEqual(solution.maxDepth(None), 0)

    def test_max_depth_single_node(self):
        # Single node tree
        root = TreeNode(1)
        
        solution = Solution()
        self.assertEqual(solution.maxDepth(root), 1)

    def test_max_depth_unbalanced_tree(self):
        # Unbalanced tree: [1,2,null,3,null,4,null,5]
        root = TreeNode(1)
        root.left = TreeNode(2)
        root.left.left = TreeNode(3)
        root.left.left.left = TreeNode(4)
        root.left.left.left.left = TreeNode(5)
        
        solution = Solution()
        self.assertEqual(solution.maxDepth(root), 5)

if __name__ == '__main__':
    unittest.main()
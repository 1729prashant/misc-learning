"""
100. Same Tree

Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

 

Example 1:


Input: p = [1,2,3], q = [1,2,3]
Output: true
Example 2:


Input: p = [1,2], q = [1,null,2]
Output: false
Example 3:


Input: p = [1,2,1], q = [1,1,2]
Output: false
 

Constraints:

The number of nodes in both trees is in the range [0, 100].
-104 <= Node.val <= 104
"""
import unittest
from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        if not p and not q:
            return True
        if not p or not q:
            return False
        if p.val != q.val:
            return False
        return self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)

class TestIsSameTree(unittest.TestCase):

    def create_tree(self, values):
        if not values:
            return None
        root = TreeNode(values[0])
        queue = [root]
        i = 1
        while queue and i < len(values):
            node = queue.pop(0)
            if i < len(values) and values[i] is not None:
                node.left = TreeNode(values[i])
                queue.append(node.left)
            i += 1
            if i < len(values) and values[i] is not None:
                node.right = TreeNode(values[i])
                queue.append(node.right)
            i += 1
        return root

    def test_same_tree_example1(self):
        p = self.create_tree([1, 2, 3])
        q = self.create_tree([1, 2, 3])
        self.assertTrue(Solution().isSameTree(p, q))

    def test_same_tree_example2(self):
        p = self.create_tree([1, 2])
        q = self.create_tree([1, None, 2])
        self.assertFalse(Solution().isSameTree(p, q))

    def test_same_tree_example3(self):
        p = self.create_tree([1, 2, 1])
        q = self.create_tree([1, 1, 2])
        self.assertFalse(Solution().isSameTree(p, q))

    def test_same_tree_both_none(self):
        self.assertTrue(Solution().isSameTree(None, None))

    def test_same_tree_one_none(self):
        p = self.create_tree([1])
        self.assertFalse(Solution().isSameTree(p, None))

    def test_same_tree_different_structure(self):
        p = self.create_tree([1, 2, 3, None, None, 4])
        q = self.create_tree([1, 2, 3, 4])
        self.assertFalse(Solution().isSameTree(p, q))

    def test_same_tree_deep_tree(self):
        p = self.create_tree([1, 2, 3, 4, 5, 6, 7])
        q = self.create_tree([1, 2, 3, 4, 5, 6, 7])
        self.assertTrue(Solution().isSameTree(p, q))

if __name__ == '__main__':
    unittest.main()
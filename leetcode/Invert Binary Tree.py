"""
226. Invert Binary Tree

Given the root of a binary tree, invert the tree, and return its root.

 

Example 1:


Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
Example 2:


Input: root = [2,1,3]
Output: [2,3,1]
Example 3:

Input: root = []
Output: []
 

Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
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
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        if root is None:
            return None

        root.left, root.right = root.right, root.left

        self.invertTree(root.left)
        self.invertTree(root.right)

        return root

class TestInvertTree(unittest.TestCase):

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

    def tree_to_list(self, root):
        if not root:
            return []
        result = []
        queue = [root]
        
        while queue:
            node = queue.pop(0)
            if node:
                result.append(node.val)
                queue.append(node.left)
                queue.append(node.right)
            else:
                result.append(None)
        
        # Clean up trailing Nones
        while result and result[-1] is None:
            result.pop()
        
        return result

    def test_invert_tree_example1(self):
        # Example 1: [4,2,7,1,3,6,9] -> [4,7,2,9,6,3,1]
        root = self.create_tree([4, 2, 7, 1, 3, 6, 9])
        inverted = Solution().invertTree(root)
        self.assertEqual(self.tree_to_list(inverted), [4, 7, 2, 9, 6, 3, 1])

    def test_invert_tree_example2(self):
        # Example 2: [2,1,3] -> [2,3,1]
        root = self.create_tree([2, 1, 3])
        inverted = Solution().invertTree(root)
        self.assertEqual(self.tree_to_list(inverted), [2, 3, 1])

    def test_invert_tree_example3(self):
        # Example 3: [] -> []
        root = self.create_tree([])
        inverted = Solution().invertTree(root)
        self.assertEqual(self.tree_to_list(inverted), [])

    def test_invert_tree_single_node(self):
        # Single node tree
        root = self.create_tree([1])
        inverted = Solution().invertTree(root)
        self.assertEqual(self.tree_to_list(inverted), [1])

    def test_invert_tree_unbalanced(self):
        # Unbalanced tree: [1,2,null,3,null,4] -> [1,null,2,null,3,null,4]
        root = self.create_tree([1, 2, None, 3, None, 4])
        inverted = Solution().invertTree(root)
        self.assertEqual(self.tree_to_list(inverted), [1, None, 2, None, 3, None, 4])


if __name__ == '__main__':
    unittest.main()
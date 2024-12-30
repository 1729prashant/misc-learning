"""
101. Symmetric Tree

Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

 

Example 1:


Input: root = [1,2,2,3,4,4,3]
Output: true
Example 2:


Input: root = [1,2,2,null,3,null,3]
Output: false
 

Constraints:

The number of nodes in the tree is in the range [1, 1000].
-100 <= Node.val <= 100
 

Follow up: Could you solve it both recursively and iteratively?
"""

from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        def is_mirror(t1: Optional[TreeNode], t2: Optional[TreeNode]) -> bool:
            if not t1 and not t2:
                return True
            if not t1 or not t2:
                return False
            return (t1.val == t2.val) and is_mirror(t1.left, t2.right) and is_mirror(t1.right, t2.left)

        return is_mirror(root, root)
    


def create_tree(values):
    if not values:
        return None
    root = TreeNode(values[0])
    queue = [root]
    i = 1
    while i < len(values):
        node = queue.pop(0)
        if node:
            if i < len(values) and values[i] is not None:
                node.left = TreeNode(values[i])
                queue.append(node.left)
            i += 1
            if i < len(values) and values[i] is not None:
                node.right = TreeNode(values[i])
                queue.append(node.right)
            i += 1
    return root

# Instantiate Solution
solution = Solution()

# Test case 1: Symmetric tree
root1 = create_tree([1, 2, 2, 3, 4, 4, 3])
print(solution.isSymmetric(root1))  # Expected: True

# Test case 2: Not symmetric
root2 = create_tree([1, 2, 2, None, 3, None, 3])
print(solution.isSymmetric(root2))  # Expected: False

# Test case 3: Single node tree
root3 = create_tree([1])
print(solution.isSymmetric(root3))  # Expected: True

# Test case 4: Empty tree
root4 = create_tree([])
print(solution.isSymmetric(root4))  # Expected: True

# Test case 5: Unbalanced symmetric tree
root5 = create_tree([1, 2, 2, 3, None, None, 3])
print(solution.isSymmetric(root5))  # Expected: True

# Test case 6: Tree with negative values
root6 = create_tree([1, -2, -2, -3, None, None, -3])
print(solution.isSymmetric(root6))  # Expected: True

# Test case 7: Large tree with symmetry
root7 = create_tree([1, 2, 2, 3, 4, 4, 3, 5, None, None, 5, 5, None, None, 5])
print(solution.isSymmetric(root7))  # Expected: True

# Test case 8: Large tree without symmetry
root8 = create_tree([1, 2, 2, 3, 4, 4, 3, 5, None, None, 6, 5, None, None, 5])
print(solution.isSymmetric(root8))  # Expected: False

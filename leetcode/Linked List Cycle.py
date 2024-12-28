"""
141. Linked List Cycle

Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

 

Example 1:


Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
Example 2:


Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
Example 3:


Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
 

Constraints:

The number of the nodes in the list is in the range [0, 104].
-105 <= Node.val <= 105
pos is -1 or a valid index in the linked-list.
 

Follow up: Can you solve it using O(1) (i.e. constant) memory?
"""


def createList(values, pos):
    if not values:
        return None
    head = ListNode(values[0])
    current = head
    nodes = [head]
    for value in values[1:]:
        current.next = ListNode(value)
        current = current.next
        nodes.append(current)
    if pos != -1:
        nodes[-1].next = nodes[pos]
    return head

# Test case 1
head = createList([3,2,0,-4], 1)
solution = Solution()
assert solution.hasCycle(head) == True, "Test case 1 failed"

# Test case 2
head = createList([1,2], 0)
assert solution.hasCycle(head) == True, "Test case 2 failed"

# Test case 3
head = createList([1], -1)
assert solution.hasCycle(head) == False, "Test case 3 failed"

# Test case 4 - Large list without cycle
head = createList([i for i in range(10000)], -1)
assert solution.hasCycle(head) == False, "Test case 4 failed"

# Test case 5 - Large list with cycle at the end
head = createList([i for i in range(10000)], 9999)
assert solution.hasCycle(head) == True, "Test case 5 failed"

print("All test cases passed!")

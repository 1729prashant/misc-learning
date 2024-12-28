/*
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
*/

package main

import (
    "testing"
)

// ListNode represents a node in a singly linked list
type ListNode struct {
    Val  int
    Next *ListNode
}

// hasCycle checks if a linked list has a cycle
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    slow := head
    fast := head.Next

    for slow != fast {
        if fast == nil || fast.Next == nil {
            return false
        }
        slow = slow.next
        fast = fast.Next.Next
    }
    return true
}

// Helper function to create a linked list with a cycle
func createList(values []int, pos int) *ListNode {
    if len(values) == 0 {
        return nil
    }
    head := &ListNode{Val: values[0]}
    var tail *ListNode
    nodes := make([]*ListNode, len(values))
    current := head
    nodes[0] = head

    for i := 1; i < len(values); i++ {
        current.Next = &ListNode{Val: values[i]}
        current = current.Next
        nodes[i] = current
    }
    tail = current

    if pos != -1 {
        tail.Next = nodes[pos]
    }
    return head
}

// Test cases
func TestHasCycle(t *testing.T) {
    tests := []struct {
        name   string
        values []int
        pos    int
        want   bool
    }{
        {"Cycle at 1", []int{3, 2, 0, -4}, 1, true},
        {"Cycle at 0", []int{1, 2}, 0, true},
        {"No cycle", []int{1}, -1, false},
        {"Large list no cycle", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, -1, false},
        {"Large list with cycle at end", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            head := createList(tt.values, tt.pos)
            if got := hasCycle(head); got != tt.want {
                t.Errorf("hasCycle() = %v, want %v", got, tt.want)
            }
        })
    }
}

func main() {
    // Run test cases
    TestHasCycle(&testing.T{})
}

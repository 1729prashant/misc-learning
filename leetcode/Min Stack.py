"""
155. Min Stack

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
You must implement a solution with O(1) time complexity for each function.

 

Example 1:

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2
 

Constraints:

-231 <= val <= 231 - 1
Methods pop, top and getMin operations will always be called on non-empty stacks.
At most 3 * 104 calls will be made to push, pop, top, and getMin.
"""


class MinStack:

    def __init__(self):
        self.stack = []
        self.min_stack = []

    def push(self, val: int) -> None:
        self.stack.append(val)
        # If min_stack is empty or val is less than or equal to the current minimum, push val
        if not self.min_stack or val <= self.min_stack[-1]:
            self.min_stack.append(val)

    def pop(self) -> None:
        if self.stack:
            if self.stack[-1] == self.min_stack[-1]:
                self.min_stack.pop()
            self.stack.pop()

    def top(self) -> int:
        if self.stack:
            return self.stack[-1]

    def getMin(self) -> int:
        if self.min_stack:
            return self.min_stack[-1]




def test_min_stack():
    minStack = MinStack()
    
    # Case 1: Regular operations
    minStack.push(-2)
    minStack.push(0)
    minStack.push(-3)
    assert minStack.getMin() == -3, "Failed to get min after pushing -3"
    minStack.pop()
    assert minStack.top() == 0, "Failed to get top after pop"
    assert minStack.getMin() == -2, "Failed to get min after pop"

    # Case 2: Check for correct handling of equal elements
    minStack = MinStack()
    minStack.push(1)
    minStack.push(1)
    assert minStack.getMin() == 1, "Failed with equal elements"
    minStack.pop()
    assert minStack.getMin() == 1, "Failed min after pop with equal elements"

    # Case 3: Check with large numbers and negative numbers
    minStack = MinStack()
    minStack.push(1000)
    minStack.push(-1000)
    minStack.push(-1000)
    minStack.push(1000)
    assert minStack.getMin() == -1000, "Failed with large and negative numbers"
    minStack.pop()
    assert minStack.getMin() == -1000, "Failed after pop with large and negative numbers"

    # Case 4: Empty stack
    minStack = MinStack()
    try:
        minStack.top()
        assert False, "Should raise IndexError on empty stack for top"
    except IndexError:
        pass
    try:
        minStack.getMin()
        assert False, "Should raise IndexError on empty stack for getMin"
    except IndexError:
        pass

    print("All tests passed!")

test_min_stack()

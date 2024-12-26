/*
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
*/

import (
    "testing"
    "reflect"
)

type MinStack struct {
    stack    []int
    minStack []int
}

func Constructor() MinStack {
    return MinStack{
        stack:    make([]int, 0),
        minStack: make([]int, 0),
    }
}

func (this *MinStack) Push(val int) {
    this.stack = append(this.stack, val)
    // If minStack is empty or val is less than or equal to the current minimum, push val
    if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
        this.minStack = append(this.minStack, val)
    }
}

func (this *MinStack) Pop() {
    if len(this.stack) > 0 {
        if this.stack[len(this.stack)-1] == this.minStack[len(this.minStack)-1] {
            this.minStack = this.minStack[:len(this.minStack)-1]
        }
        this.stack = this.stack[:len(this.stack)-1]
    }
}

func (this *MinStack) Top() int {
    if len(this.stack) > 0 {
        return this.stack[len(this.stack)-1]
    }
    return 0 // Return 0 or handle error for empty stack
}

func (this *MinStack) GetMin() int {
    if len(this.minStack) > 0 {
        return this.minStack[len(this.minStack)-1]
    }
    return 0 // Return 0 or handle error for empty stack
}






func TestMinStack(t *testing.T) {
    minStack := Constructor()

    // Test case 1: Regular operations
    minStack.Push(-2)
    minStack.Push(0)
    minStack.Push(-3)
    if minStack.GetMin() != -3 {
        t.Errorf("Expected -3, got %d", minStack.GetMin())
    }
    minStack.Pop()
    if minStack.Top() != 0 {
        t.Errorf("Expected 0, got %d", minStack.Top())
    }
    if minStack.GetMin() != -2 {
        t.Errorf("Expected -2, got %d", minStack.GetMin())
    }

    // Test case 2: Check for correct handling of equal elements
    minStack = Constructor()
    minStack.Push(1)
    minStack.Push(1)
    if minStack.GetMin() != 1 {
        t.Errorf("Expected 1, got %d for equal elements", minStack.GetMin())
    }
    minStack.Pop()
    if minStack.GetMin() != 1 {
        t.Errorf("Expected 1, got %d after pop with equal elements", minStack.GetMin())
    }

    // Test case 3: Check with large numbers and negative numbers
    minStack = Constructor()
    minStack.Push(1000)
    minStack.Push(-1000)
    minStack.Push(-1000)
    minStack.Push(1000)
    if minStack.GetMin() != -1000 {
        t.Errorf("Expected -1000, got %d for large and negative numbers", minStack.GetMin())
    }
    minStack.Pop()
    if minStack.GetMin() != -1000 {
        t.Errorf("Expected -1000, got %d after pop with large and negative numbers", minStack.GetMin())
    }

    // Test case 4: Empty stack operations
    minStack = Constructor()
    if minStack.Top() != 0 {
        t.Errorf("Expected 0 for empty stack Top, got %d", minStack.Top())
    }
    if minStack.GetMin() != 0 {
        t.Errorf("Expected 0 for empty stack GetMin, got %d", minStack.GetMin())
    }
}

func TestMinStackPopEmpty(t *testing.T) {
    minStack := Constructor()
    minStack.Push(1)
    minStack.Pop()
    minStack.Pop()  // Should not panic; does nothing for an empty stack
    if len(minStack.stack) != 0 || len(minStack.minStack) != 0 {
        t.Errorf("Stack should be empty after pop on empty stack")
    }
}

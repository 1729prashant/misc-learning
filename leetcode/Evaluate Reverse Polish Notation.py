"""
150. Evaluate Reverse Polish Notation

You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:

The valid operators are '+', '-', '*', and '/'.
Each operand may be an integer or another expression.
The division between two integers always truncates toward zero.
There will not be any division by zero.
The input represents a valid arithmetic expression in a reverse polish notation.
The answer and all the intermediate calculations can be represented in a 32-bit integer.
 

Example 1:

Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
Example 2:

Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
Example 3:

Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
 

Constraints:

1 <= tokens.length <= 104
tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].
"""

def evalRPN(tokens):
    stack = []
    for t in tokens:
        if t in {'+', '-', '*', '/'}:
            b = stack.pop()
            a = stack.pop()
            if t == '+':
                stack.append(a + b)
            elif t == '-':
                stack.append(a - b)
            elif t == '*':
                stack.append(a * b)
            else:  # Division
                # Python 3 division returns float, so we need to truncate towards zero
                stack.append(int(float(a) / b))
        else:
            # Convert string to integer
            stack.append(int(t))
    return stack[0]

# Test cases
def test_evalRPN():
    assert evalRPN(["2","1","+","3","*"]) == 9, "Test case 1 failed"
    assert evalRPN(["4","13","5","/","+"]) == 6, "Test case 2 failed"
    assert evalRPN(["10","6","9","3","+","-11","*","/","*","17","+","5","+"]) == 22, "Test case 3 failed"
    
    # Additional test cases for edge cases
    assert evalRPN(["15","7","1","1","+","-","/","3","*","2","1","1","+","+","-"]) == 5, "Test case 4 failed"
    assert evalRPN(["4","3","-"]) == 1, "Test case 5 failed"
    assert evalRPN(["3","-4","+"]) == -1, "Test case 6 failed"
    assert evalRPN(["78","3","4","+","*","-8","/","-"]) == -6, "Test case 7 failed"
    
    print("All test cases passed!")

# Run the test cases
test_evalRPN()

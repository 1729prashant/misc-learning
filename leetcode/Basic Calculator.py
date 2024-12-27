"""
224. Basic Calculator

Given a string s representing a valid expression, implement a basic calculator to evaluate it, and return the result of the evaluation.

Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().

 

Example 1:

Input: s = "1 + 1"
Output: 2
Example 2:

Input: s = " 2-1 + 2 "
Output: 3
Example 3:

Input: s = "(1+(4+5+2)-3)+(6+8)"
Output: 23
 

Constraints:

1 <= s.length <= 3 * 105
s consists of digits, '+', '-', '(', ')', and ' '.
s represents a valid expression.
'+' is not used as a unary operation (i.e., "+1" and "+(2 + 3)" is invalid).
'-' could be used as a unary operation (i.e., "-1" and "-(2 + 3)" is valid).
There will be no two consecutive operators in the input.
Every number and running calculation will fit in a signed 32-bit integer.

"""

class Solution:
    def calculate(self, s: str) -> int:
        stack = []
        number = 0
        sign = 1
        result = 0
        
        for char in s:
            if char.isdigit():
                # Building the number
                number = number * 10 + int(char)
            elif char in ['+', '-']:
                # Adding the previous number with its sign to the result
                result += sign * number
                # Reset number and set new sign
                number = 0
                sign = 1 if char == '+' else -1
            elif char == '(':
                # Push the current result and sign onto the stack
                stack.append(result)
                stack.append(sign)
                # Reset result and sign for the sub-expression
                result = 0
                sign = 1
            elif char == ')':
                # Calculate the current number with the sign and add it to the result
                result += sign * number
                # Pop the sign and previous result from the stack
                result *= stack.pop()  # multiply by the sign before the parenthesis
                result += stack.pop()  # add to the result before the parenthesis
                number = 0
            # Skip spaces
        
        # Add the last number to the result
        return result + sign * number

# Test cases
def test_calculate():
    solution = Solution()
    
    assert solution.calculate("1 + 1") == 2, "Failed for '1 + 1'"
    assert solution.calculate(" 2-1 + 2 ") == 3, "Failed for ' 2-1 + 2 '"
    assert solution.calculate("(1+(4+5+2)-3)+(6+8)") == 23, "Failed for '(1+(4+5+2)-3)+(6+8)'"
    assert solution.calculate(" 2-1 + 2 ") == 3, "Failed for ' 2-1 + 2 '"
    assert solution.calculate("-(3 + (1 - 2))") == 2, "Failed for '-(3 + (1 - 2))'"
    assert solution.calculate("0") == 0, "Failed for '0'"
    assert solution.calculate(" -1") == -1, "Failed for ' -1'"
    assert solution.calculate("1-1+1") == 1, "Failed for '1-1+1'"
    assert solution.calculate(" (3)") == 3, "Failed for ' (3)'"
    assert solution.calculate("2-(5-6)") == 3, "Failed for '2-(5-6)'"
    
    print("All test cases passed!")

# Run the test
test_calculate()

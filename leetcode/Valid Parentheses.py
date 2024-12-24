"""
20. Valid Parentheses

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
 

Example 1:

Input: s = "()"

Output: true

Example 2:

Input: s = "()[]{}"

Output: true

Example 3:

Input: s = "(]"

Output: false

Example 4:

Input: s = "([])"

Output: true

 

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
"""

class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        mapping = {")": "(", "]": "[", "}": "{"}
        
        for char in s:
            if char in mapping:
                # If the character is a closing bracket
                top_element = stack.pop() if stack else '#'
                if mapping[char] != top_element:
                    return False
            else:
                # If the character is an opening bracket, push it onto the stack
                stack.append(char)
        
        # If the stack is empty, all brackets were matched
        return not stack



# Test cases for the isValid function
def test_isValid():
    solution = Solution()

    # Test case 1: Balanced simple brackets
    assert solution.isValid("()") == True, "Failed for '()'"

    # Test case 2: Balanced multiple types of brackets
    assert solution.isValid("()[]{}") == True, "Failed for '()[]{}'"

    # Test case 3: Unbalanced brackets
    assert solution.isValid("(]") == False, "Failed for '(]'"

    # Test case 4: Nested brackets, balanced
    assert solution.isValid("([])") == True, "Failed for '([])'"

    # Test case 5: Mismatched types of brackets
    assert solution.isValid("([)]") == False, "Failed for '([)]'"

    # Test case 6: Too many closing brackets
    assert solution.isValid("((())))") == False, "Failed for '((())))'"

    # Test case 7: Empty string
    assert solution.isValid("") == True, "Failed for empty string"

    # Test case 8: Only opening brackets
    assert solution.isValid("(((") == False, "Failed for '((('"

    # Test case 9: Only closing brackets
    assert solution.isValid(")))") == False, "Failed for ')))'"

    # Test case 10: Complex nested brackets
    assert solution.isValid("{[]()}") == True, "Failed for '{[]()}'"

    # Test case 11: Very complex but balanced
    assert solution.isValid("{[()(){}]}") == True, "Failed for '{[()(){}]}'"

    print("All test cases passed!")

# Run the test cases
test_isValid()

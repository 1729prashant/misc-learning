/*
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
*/

package main

import "fmt"

type Solution struct{}

func (s *Solution) isValid(str string) bool {
    stack := make([]rune, 0)
    bracketMap := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, char := range str {
        if char == '(' || char == '[' || char == '{' {
            stack = append(stack, char)
        } else if char == ')' || char == ']' || char == '}' {
            if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}

func main() {
    solution := &Solution{}
    testCases(solution)
}

func testCases(s *Solution) {
    tests := []struct {
        input string
        want  bool
    }{
        {"()", true},
        {"()[]{}", true},
        {"(]", false},
        {"([])", true},
        {"([)]", false},
        {"((())))", false},
        {"", true},
        {"(((", false},
        {")))", false},
        {"{[()()]}", true},
        {"{[()(){}]}", true},
    }

    for _, tt := range tests {
        if got := s.isValid(tt.input); got != tt.want {
            fmt.Printf("isValid(%q) = %v, want %v\n", tt.input, got, tt.want)
        }
    }
    fmt.Println("All tests passed!")
}

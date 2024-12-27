/*
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
*/

package main

import (
    "fmt"
    "strconv"
    "testing"
)

// evalRPN evaluates a Reverse Polish Notation expression
func evalRPN(tokens []string) int {
    stack := []int{}
    for _, token := range tokens {
        switch token {
        case "+":
            b, a := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-2]
            stack = append(stack, a+b)
        case "-":
            b, a := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-2]
            stack = append(stack, a-b)
        case "*":
            b, a := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-2]
            stack = append(stack, a*b)
        case "/":
            b, a := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-2]
            // Division truncates towards zero in Go for integers
            stack = append(stack, a/b)
        default:
            num, _ := strconv.Atoi(token)
            stack = append(stack, num)
        }
    }
    return stack[0]
}

// Test cases
func TestEvalRPN(t *testing.T) {
    tests := []struct {
        name     string
        tokens   []string
        expected int
    }{
        {"Example 1", []string{"2", "1", "+", "3", "*"}, 9},
        {"Example 2", []string{"4", "13", "5", "/", "+"}, 6},
        {"Example 3", []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
        {"Complex", []string{"15", "7", "1", "1", "+", "-", "/", "3", "*", "2", "1", "1", "+", "+", "-"}, 5},
        {"Simple Subtraction", []string{"4", "3", "-"}, 1},
        {"Negative Number", []string{"3", "-4", "+"}, -1},
        {"Division and Negative", []string{"78", "3", "4", "+", "*", "-8", "/", "-"}, -6},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := evalRPN(tt.tokens)
            if result != tt.expected {
                t.Errorf("evalRPN() = %v, want %v", result, tt.expected)
            }
        })
    }
}

func main() {
    fmt.Println("Running tests...")
}

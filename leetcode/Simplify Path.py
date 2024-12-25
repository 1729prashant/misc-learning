"""
71. Simplify Path

You are given an absolute path for a Unix-style file system, which always begins with a slash '/'. Your task is to transform this absolute path into its simplified canonical path.

The rules of a Unix-style file system are as follows:

A single period '.' represents the current directory.
A double period '..' represents the previous/parent directory.
Multiple consecutive slashes such as '//' and '///' are treated as a single slash '/'.
Any sequence of periods that does not match the rules above should be treated as a valid directory or file name. For example, '...' and '....' are valid directory or file names.
The simplified canonical path should follow these rules:

The path must start with a single slash '/'.
Directories within the path must be separated by exactly one slash '/'.
The path must not end with a slash '/', unless it is the root directory.
The path must not have any single or double periods ('.' and '..') used to denote current or parent directories.
Return the simplified canonical path.

 

Example 1:

Input: path = "/home/"

Output: "/home"

Explanation:

The trailing slash should be removed.

Example 2:

Input: path = "/home//foo/"

Output: "/home/foo"

Explanation:

Multiple consecutive slashes are replaced by a single one.

Example 3:

Input: path = "/home/user/Documents/../Pictures"

Output: "/home/user/Pictures"

Explanation:

A double period ".." refers to the directory up a level (the parent directory).

Example 4:

Input: path = "/../"

Output: "/"

Explanation:

Going one level up from the root directory is not possible.

Example 5:

Input: path = "/.../a/../b/c/../d/./"

Output: "/.../b/d"

Explanation:

"..." is a valid name for a directory in this problem.

 

Constraints:

1 <= path.length <= 3000
path consists of English letters, digits, period '.', slash '/' or '_'.
path is a valid absolute Unix path.
"""

class Solution:
    def simplifyPath(self, path: str) -> str:
        stack = []
        for directory in path.split('/'):
            if directory == '..' and stack:
                stack.pop()
            elif directory not in {'', '.', '..'}:
                stack.append(directory)
        return '/' + '/'.join(stack)

def test_simplifyPath():
    solution = Solution()

    # Example 1
    assert solution.simplifyPath("/home/") == "/home", "Failed for '/home/'"

    # Example 2
    assert solution.simplifyPath("/home//foo/") == "/home/foo", "Failed for '/home//foo/'"

    # Example 3
    assert solution.simplifyPath("/home/user/Documents/../Pictures") == "/home/user/Pictures", "Failed for '/home/user/Documents/../Pictures'"

    # Example 4
    assert solution.simplifyPath("/../") == "/", "Failed for '/../'"

    # Example 5
    assert solution.simplifyPath("/.../a/../b/c/../d/./") == "/.../b/d", "Failed for '/.../a/../b/c/../d/./'"

    # Additional test cases
    assert solution.simplifyPath("/a/./b/../../c/") == "/c", "Failed for '/a/./b/../../c/'"
    assert solution.simplifyPath("/a/../.././../b") == "/b", "Failed for '/a/../.././../b'"
    assert solution.simplifyPath("/") == "/", "Failed for '/'"
    assert solution.simplifyPath("/.../..../") == "/.../....", "Failed for '/.../..../'"
    assert solution.simplifyPath("/./././") == "/", "Failed for '/./././'"

    print("All test cases passed!")

# Run the test cases
test_simplifyPath()

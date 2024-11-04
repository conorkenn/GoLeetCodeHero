## 20. Valid Parentheses

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

### Example 1:
```
Input: s = "()"
Output: true
```
### Example 2:
```
Input: s = "()[]{}"
Output: true
```
### Example 3:
```
Input: s = "(]"
Output: false
```

### Example 4:
```
Input: s = "([])"
Output: true
```
##
### Constraints:

- `1 <= s.length <= 10^4`
- `s consists of parentheses only '()[]{}'.`


## Answer
```
func isValid(s string) bool {
    
    lparen := "[{("
    rparen := "]})"
    stack := []rune{}
    for _, paren := range s {
        if strings.ContainsRune(lparen, paren){
            stack = append(stack, paren)
        } else {
            if len(stack) == 0{
                return false
            }
            t := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if strings.IndexRune(lparen, t) != strings.IndexRune(rparen, paren){
                return false
            }
        }
    }
    return len(stack) == 0
}
```

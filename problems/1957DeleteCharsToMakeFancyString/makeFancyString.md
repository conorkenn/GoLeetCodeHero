## 1957. Delete Characters to Make Fancy String

A fancy string is a string where no three consecutive characters are equal.

Given a string `s`, delete the minimum possible number of characters from `s` to make it fancy.

Return *the final string after the deletion*. It can be shown that the answer will always be unique.
##

### Example 1:
```
Input: s = "leeetcode"
Output: "leetcode"
Explanation:
Remove an 'e' from the first group of 'e's to create "leetcode".
No three consecutive characters are equal, so return "leetcode".
```
### Example 2:
```
Input: s = "aaabaaaa"
Output: "aabaa"
Explanation:
Remove an 'a' from the first group of 'a's to create "aabaaaa".
Remove two 'a's from the second group of 'a's to create "aabaa".
No three consecutive characters are equal, so return "aabaa"
```
### Example 3:
```
Input: s = "aab"
Output: "aab"
Explanation: No three consecutive characters are equal, so return "aab".
```
##
### Constraints:

- `1 <= s.length <= 105`
- `s` consists only of lowercase English letters.


## Answer
```
func makeFancyString(s string) string {
    var stack[]byte

    for i:=0; i < len(s); i++ {
        if len(stack) >= 2 && stack[len(stack)-1] == s[i] && stack[len(stack)-2] == s[i] {
            continue        
        }
        stack = append(stack, s[i])
    }

    return string(stack)
}
```

Initialize a stack and loop over the string. Putting every char into the stack after checking if the last two elements are the same char and continuing if they are. Return the stack as a string
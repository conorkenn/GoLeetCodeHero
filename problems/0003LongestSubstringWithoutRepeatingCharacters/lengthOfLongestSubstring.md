## 3. Longest Substring Without Repeating Characters

Given a string `s`, find the length of the longest substring without repeating characters.
##

### Example 1:
```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```
### Example 2:
```
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```
### Example 3:
```
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
```
##
### Constraints:

- 0 <= s.length <= 5 * 104
- s consists of English letters, digits, symbols and spaces.

## Answer
```
func lengthOfLongestSubstring(s string) int {
    substring := []rune{}
    ans := 0
    runes := []rune(s)
    for _, c := range runes {
        if !slices.Contains(substring, c) {
            substring = append(substring, c)
        } else {
            for slices.Contains(substring, c) {
                substring = substring[1:]
            }
            substring = append(substring, c)
        }
        if len(substring) > ans {
            ans = len(substring)
        }
    }

    return ans
}
```

1. Create a slice to hold substring characters. 
2. We loop over all the characters in the string adding to the substring slice. 
3. When we find a duplicate value we pop elements out of the substring slice until we remove the duplicate. 
5. With each iteration we check see if the slice we currently have is larger than `ans`
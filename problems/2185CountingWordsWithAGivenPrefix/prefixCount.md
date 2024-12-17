## 2185. Counting Words With a Given Prefix

You are given an array of strings words and a string pref.

Return the number of strings in words that contain pref as a prefix.

A prefix of a string s is any leading contiguous substring of s.
##

### Example 1:
```
Input: words = ["pay","attention","practice","attend"], pref = "at"
Output: 2
Explanation: The 2 strings that contain "at" as a prefix are: "attention" and "attend".
```
### Example 2:
```
Input: words = ["leetcode","win","loops","success"], pref = "code"
Output: 0
Explanation: There are no strings that contain "code" as a prefix.
```
### Example 3:
```
Input: nums = [1,1,2,2], n = 2
Output: [1,2,1,2]
```
##
### Constraints:

- 1 <= words.length <= 100
- 1 <= words[i].length, pref.length <= 100
- words[i] and pref consist of lowercase English letters.


## Answer
```
func prefixCount(words []string, pref string) int {
    count := 0

    for _, word := range words {
        if strings.HasPrefix(word, pref){
            count += 1
        }
    }

    return count
}
```
## 1832. Check if the Sentence Is Pangram

A pangram is a sentence where every letter of the English alphabet appears at least once.

Given a string `sentence` containing only lowercase English letters, return `true` if `sentence` is a pangram, or `false` otherwise.
##

### Example 1:
```
Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
Output: true
Explanation: sentence contains at least one of every letter of the English alphabet.
```
### Example 2:
```
Input: sentence = "leetcode"
Output: false
```
##
### Constraints:

- 1 <= sentence.length <= 1000
- sentence consists of lowercase English letters.


## Answer
```
func checkIfPangram(sentence string) bool {
    d := make(map[rune]int)

    for _, c := range sentence{
        if _, ok := d[c]; ok {
            d[c] += 1
        } else {
            d[c] = 1
        }
    }
    return len(d) == 26
}
```


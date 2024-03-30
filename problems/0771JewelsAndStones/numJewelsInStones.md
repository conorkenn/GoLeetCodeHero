## 771. Jewels and Stones


You're given strings `jewels` representing the types of `stones` that are jewels, and stones representing the stones you have. Each character in `stones` is a type of stone you have. You want to know how many of the stones you have are also jewels.

Letters are case sensitive, so "a" is considered a different type of stone from "A".
##

### Example 1:
```
Input: jewels = "aA", stones = "aAAbbbb"
Output: 3
```
### Example 2:
```
Input: jewels = "z", stones = "ZZ"
Output: 0
```
##
### Constraints:

- `1 <= jewels.length, stones.length <= 50`
- `jewels` and `stones` consist of only English letters.
- All the characters of `jewels` are unique.


## Answer
```
func numJewelsInStones(jewels string, stones string) int {
    j := map[rune]bool{}

    for _, jewel := range jewels {
        j[jewel] = true
    }
    ans := 0

    for _,stone := range stones {
        if j[stone]{
            ans += 1
        }
    }

    return ans
}
```

1. Loop over the `jewels` and add them to a map
2. Loop over `stones` and check is the stone exist in the map
3. Return number of times we found a `stone` in the map
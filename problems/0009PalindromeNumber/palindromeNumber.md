## 9. Palindrome Number

Given an integer `x`, return `true` if `x` is a 
palindrome
, and `false` otherwise.
##

### Example 1:
```
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

```
### Example 2:
```
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```
### Example 3:
```
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```
##
### Constraints:

- -2^31 <= x <= 2^31 - 1


## Answer
```
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    num := x
    reversed := 0

    for x > 0 {
        reversed = reversed * 10 + x % 10
        x /= 10
    }

    return reversed == num
}
```

We know negative numbers will always be false 

For the postive numbers:
1. store the input
2. reverse the input like the problem 7. Reverse Integer
3. compare the reversed number with the initial value
## 7. Reverse Integer

Given a signed 32-bit integer `x`, return `x` with its digits reversed. If reversing `x` causes the value to go outside the signed 32-bit integer range `[-231, 231 - 1]`, then return `0`.

**Assume the environment does not allow you to store 64-bit integers (signed or unsigned).**
##

### Example 1:
```
Input: x = 123
Output: 321
```
### Example 2:
```
Input: x = -123
Output: -321
```
### Example 3:
```
Input: x = 120
Output: 21
```
##
### Constraints:

- `-2^31 <= x <= 2^31 - 1`


## Answer
```
func reverse(x int) int {
    ans := 0
    isNegative := 1
    if x < 0 {
        isNegative = -1
        x *= isNegative
    }

    for x > 0{
        ans = ans * 10 + x % 10
         if ans > 2147483647  || ans < -2147483648{
            return 0
        }
        x /= 10
    }
    return ans * isNegative
}
```

1. Initialize a variable to hold the answer to zero, `ans`
2. Check if the input is negative, if so we can make the value postive and update back to negative before we return
3. Loop while x > 0, we can multiple by 10 and add the last digit of x which we can get with `% 10` 
4. There is a constraint if the reversed number gets too large or small we return `0`
5. Return ans
## 1. Two Sum

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have **exactly one solution**, and you may not use the same element twice.

You can return the answer in any order.
##

### Example 1:
```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```
### Example 2:
```
Input: nums = [3,2,4], target = 6
Output: [1,2] 
```
### Example 3:
```
Input: nums = [3,3], target = 6
Output: [0,1]
```
##
### Constraints:

- 2 <= nums.length <= 104
- -109 <= nums[i] <= 109
- -109 <= target <= 109
- Only one valid answer exists.


## Answer
```
func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for i, num := range nums{ 
        value, ok := seen[target - num]
        if ok {
            return []int{i, value}
        }
        seen[num] = i
    }
    return nil
}
```

The optimal solution to `twoSum` is to create a map that we will use to track the numbers we have seen and the index where we saw them. 

Each iteration we will check if `target - num` exists in our map if it does we will return that index and the current index in a slice. 
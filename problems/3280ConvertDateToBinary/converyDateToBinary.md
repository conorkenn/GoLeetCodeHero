## 3280. Convert Date to Binary

You are given a string `date` representing a Gregorian calendar date in the `yyyy-mm-dd` format.

The date can be written in its binary representation obtained by converting the year, month, and day to their binary representations without any leading zeroes and writing them down in `year-month-day` format.

Return the binary representation of the date.


### Example 1:
Input: date = "2080-02-29"
 Output: "100000100000-10-11101"

Explanation:

100000100000 is the binary representation of 2080.
10 is the binary representation of 02 (2).
11101 is the binary representation of 29.

### Example 2:
Input: date = "1900-01-01" 
Output: "11101101100-1-1"

Explanation:

11101101100 is the binary representation of 1900.
1 is the binary representation of 01 (1).
1 is the binary representation of 01 (1).

### Constraints:

- `date.length == 10`
- `date[4] == date[7] == '-'`, and all other `date[i]`'s are digits.
- The input is generated such that `date` represents a valid Gregorian calendar date between Jan 1st, 1900 and Dec 31st, 2100 (both inclusive).


## Answer
```
func convertDateToBinary(date string) string {
    parts := strings.Split(date, "-")
    var binaryParts []string
    for _, part := range parts {
        num, _ := strconv.Atoi(part)
        binaryParts = append(binaryParts, strconv.FormatInt(int64(num),2))
    }
    return strings.Join(binaryParts, "-")
}
```
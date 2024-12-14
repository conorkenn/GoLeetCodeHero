package leetcode

import (
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	parts := strings.Split(date, "-")
	var binaryParts []string
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		binaryParts = append(binaryParts, strconv.FormatInt(int64(num), 2))
	}
	return strings.Join(binaryParts, "-")

}

package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	jewels string
	stones string
	result int
}

func TestTwoSum(t *testing.T) {

	test1 := TestCase{"aA", "aAAbbbb", 3}

	test2 := TestCase{"z", "ZZ", 0}

	t.Run("Test case 1", func(t *testing.T) {
		got := numJewelsInStones(test1.jewels, test1.stones)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := numJewelsInStones(test2.jewels, test2.stones)
		want := test2.result
		checkResult(t, got, want)

	})

}

func checkResult(t testing.TB, got, want int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

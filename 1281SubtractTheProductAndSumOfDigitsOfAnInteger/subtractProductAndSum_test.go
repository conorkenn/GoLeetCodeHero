package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	n      int
	result int
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	test1 := TestCase{234, 15}
	test2 := TestCase{4421, 21}

	t.Run("Test case 1", func(t *testing.T) {
		got := subtractProductAndSum(test1.n)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := subtractProductAndSum(test2.n)
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

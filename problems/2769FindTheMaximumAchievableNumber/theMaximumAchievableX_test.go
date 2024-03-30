package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	num    int
	t      int
	result int
}

func TestTheMaximumAchievableX(t *testing.T) {
	test1 := TestCase{4, 1, 6}

	test2 := TestCase{3, 2, 7}

	t.Run("Test case 1", func(t *testing.T) {
		got := theMaximumAchievableX(test1.num, test1.t)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := theMaximumAchievableX(test2.num, test2.t)
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

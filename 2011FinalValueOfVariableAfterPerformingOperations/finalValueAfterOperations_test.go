package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	operations []string
	result     int
}

func TestBuildArray(t *testing.T) {
	test1 := TestCase{
		[]string{"--X", "X++", "X++"},
		1,
	}

	test2 := TestCase{
		[]string{"++X", "++X", "X++"},
		3,
	}

	test3 := TestCase{
		[]string{"X++", "++X", "--X", "X--"},
		0,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := finalValueAfterOperations(test1.operations)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := finalValueAfterOperations(test2.operations)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := finalValueAfterOperations(test3.operations)
		want := test3.result
		checkResult(t, got, want)

	})

}

func checkResult(t testing.TB, got, want int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

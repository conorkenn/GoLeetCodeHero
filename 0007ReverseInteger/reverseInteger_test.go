package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	x      int
	result int
}

func TestReverseInteger(t *testing.T) {

	test1 := TestCase{123, 321}
	test2 := TestCase{-123, -321}
	test3 := TestCase{120, 21}

	t.Run("Test case 1", func(t *testing.T) {
		got := reverse(test1.x)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := reverse(test2.x)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := reverse(test3.x)
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

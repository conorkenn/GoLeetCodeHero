package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	s      string
	result int
}

func TestCanAliceWin(t *testing.T) {

	test1 := TestCase{
		"B0B6G0R6R0R6G9",
		1,
	}

	test2 := TestCase{
		"B0R0G0R9R0B0G0",
		1,
	}

	test3 := TestCase{
		"G4",
		0,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := countPoints(test1.s)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := countPoints(test2.s)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := countPoints(test3.s)
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

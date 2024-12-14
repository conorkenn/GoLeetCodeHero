package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	date   string
	result string
}

func TestCanAliceWin(t *testing.T) {

	test1 := TestCase{
		"2080-02-29",
		"100000100000-10-11101",
	}

	test2 := TestCase{
		"1900-01-01",
		"11101101100-1-1",
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := convertDateToBinary(test1.date)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := convertDateToBinary(test2.date)
		want := test2.result
		checkResult(t, got, want)

	})
}

func checkResult(t testing.TB, got, want string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

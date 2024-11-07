package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	num    string
	result bool
}

func TestCircularSentence(t *testing.T) {

	test1 := TestCase{
		"1234",
		false,
	}

	test2 := TestCase{
		"24123",
		true,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := isBalanced(test1.num)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := isBalanced(test2.num)
		want := test2.result
		checkResult(t, got, want)

	})

}

func checkResult(t testing.TB, got, want bool) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

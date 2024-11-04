package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	s      string
	result bool
}

func TestContainsDuplicates(t *testing.T) {

	test1 := TestCase{"()", true}
	test2 := TestCase{"()[]{}", true}
	test3 := TestCase{"(]", false}
	test4 := TestCase{"([])", true}

	t.Run("Test case 1", func(t *testing.T) {
		got := isValid(test1.s)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := isValid(test2.s)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := isValid(test3.s)
		want := test3.result
		checkResult(t, got, want)

	})

	t.Run("Test case 4", func(t *testing.T) {
		got := isValid(test4.s)
		want := test4.result
		checkResult(t, got, want)

	})

}

func checkResult(t testing.TB, got, want bool) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

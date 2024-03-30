package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	sentence string
	result   bool
}

func TestMostWordsFound(t *testing.T) {
	test1 := TestCase{"thequickbrownfoxjumpsoverthelazydog", true}
	test2 := TestCase{"leetcode", false}

	t.Run("Test case 1", func(t *testing.T) {
		got := checkIfPangram(test1.sentence)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := checkIfPangram(test2.sentence)
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

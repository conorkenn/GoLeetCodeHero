package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	words  []string
	pref   string
	result int
}

func TestCanAliceWin(t *testing.T) {

	test1 := TestCase{
		[]string{"pay", "attention", "practice", "attend"},
		"at",
		2,
	}

	test2 := TestCase{
		[]string{"leetcode", "win", "loops", "success"},
		"code",
		0,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := prefixCount(test1.words, test1.pref)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := prefixCount(test2.words, test2.pref)
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

package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	sentences []string
	result    int
}

func TestMostWordsFound(t *testing.T) {
	test1 := TestCase{[]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}, 6}
	test2 := TestCase{[]string{"please wait", "continue to fight", "continue to win"}, 3}

	t.Run("Test case 1", func(t *testing.T) {
		got := mostWordsFound(test1.sentences)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := mostWordsFound(test2.sentences)
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

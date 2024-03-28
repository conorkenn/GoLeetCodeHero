package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	n      string
	result int
}

func TestMinPartitions(t *testing.T) {

	test1 := TestCase{"32", 3}
	test2 := TestCase{"82734", 8}
	test3 := TestCase{"27346209830709182346", 9}

	t.Run("Test case 1", func(t *testing.T) {
		got := minPartitions(test1.n)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := minPartitions(test2.n)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := minPartitions(test3.n)
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

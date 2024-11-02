package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	prices []int
	result int
}

func TestCircularSentence(t *testing.T) {

	test1 := TestCase{
		[]int{7, 1, 5, 3, 6, 4},
		5,
	}

	test2 := TestCase{
		[]int{7, 6, 4, 3, 1},
		0,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := maxProfit(test1.prices)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := maxProfit(test2.prices)
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

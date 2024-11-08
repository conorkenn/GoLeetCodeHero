package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	flowerbed []int
	n         int
	result    bool
}

func TestCircularSentence(t *testing.T) {

	test1 := TestCase{
		[]int{1, 0, 0, 0, 1},
		1,
		true,
	}

	test2 := TestCase{
		[]int{1, 0, 0, 0, 1},
		2,
		false,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := canPlaceFlowers(test1.flowerbed, test1.n)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := canPlaceFlowers(test2.flowerbed, test2.n)
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

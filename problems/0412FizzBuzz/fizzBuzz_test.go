package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	n      int
	result []string
}

func TestMostWordsFound(t *testing.T) {

	test1 := TestCase{3, []string{"1", "2", "Fizz"}}
	test2 := TestCase{5, []string{"1", "2", "Fizz", "4", "Buzz"}}
	test3 := TestCase{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}}

	t.Run("Test case 1", func(t *testing.T) {
		got := fizzBuzz(test1.n)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := fizzBuzz(test2.n)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := fizzBuzz(test3.n)
		want := test3.result
		checkResult(t, got, want)

	})

}

func checkResult(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

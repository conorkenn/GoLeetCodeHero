package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	celsius float64
	result  []float64
}

func TestConvertTemperature(t *testing.T) {
	test1 := TestCase{36.50, []float64{309.65000, 97.70000}}
	test2 := TestCase{122.11, []float64{395.26000, 251.79800}}

	t.Run("Test case 1", func(t *testing.T) {
		got := convertTemperature(test1.celsius)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := convertTemperature(test2.celsius)
		want := test2.result
		checkResult(t, got, want)

	})
}

func checkResult(t testing.TB, got, want []float64) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

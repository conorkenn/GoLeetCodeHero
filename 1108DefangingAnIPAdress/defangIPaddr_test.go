package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	address string
	result  string
}

func TestDefangIPaddr(t *testing.T) {
	test1 := TestCase{"1.1.1.1", "1[.]1[.]1[.]1"}
	test2 := TestCase{"255.100.50.0", "255[.]100[.]50[.]0"}

	t.Run("Test case 1", func(t *testing.T) {
		got := defangIPaddr(test1.address)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := defangIPaddr(test2.address)
		want := test2.result
		checkResult(t, got, want)

	})

}

func checkResult(t testing.TB, got, want string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s want %s", got, want)
	}
}

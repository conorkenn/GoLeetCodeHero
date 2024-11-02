package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	head   *ListNode
	result *ListNode
}

func createLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	for _, v := range vals[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

func getNodeAt(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		head = head.Next
	}
	return head
}

func TestCircularSentence(t *testing.T) {

	list1 := createLinkedList([]int{1, 2, 3, 4, 5})
	middle1 := getNodeAt(list1, 2) // middle node is 3
	test1 := TestCase{
		head:   list1,
		result: middle1,
	}

	list2 := createLinkedList([]int{1, 2, 3, 4, 5, 6})
	middle2 := getNodeAt(list2, 3) // middle node is 4
	test2 := TestCase{
		head:   list2,
		result: middle2,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := middleNode(test1.head)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := middleNode(test2.head)
		want := test2.result
		checkResult(t, got, want)

	})

}

func checkResult(t testing.TB, got, want *ListNode) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

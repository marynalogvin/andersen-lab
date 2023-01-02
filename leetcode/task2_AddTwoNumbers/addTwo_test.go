package addTwo

import (
	"testing"
)

func TestAddTwo(t *testing.T) {
	var testCases = []struct {
		list1 *Node
		list2 *Node
		want  *Node
	}{
		{
			list1: createList([]int{1, 2, 3}),
			list2: createList([]int{4, 5, 6}),
			want:  createList([]int{5, 7, 9}),
		},
		{
			list1: createList([]int{0, 2, 3}),
			list2: createList([]int{5, 6}),
			want:  createList([]int{5, 8, 3}),
		}, {
			list1: createList([]int{1, 2}),
			list2: createList([]int{4, 5, 6}),
			want:  createList([]int{5, 7, 6}),
		}, {
			list1: createList([]int{1, 2, 6}),
			list2: createList([]int{4, 5, 6}),
			want:  createList([]int{5, 7, 2, 1}),
		},
	}
	for _, test := range testCases {
		result := addTwo(test.list1, test.list2)
		ok := compareResults(result, test.want)
		if !ok {
			t.Errorf("Input: %v, %v => %v, want %v", test.list1, test.list2, result, test.want)
		}
	}
}

func compareResults(result *Node, want *Node) bool {
	currOk := true
	nextOk := true
	if result.next != nil || want.next != nil {
		nextOk = compareResults(result.next, want.next)
		currOk = result.value == want.value
	}
	return nextOk && currOk

}

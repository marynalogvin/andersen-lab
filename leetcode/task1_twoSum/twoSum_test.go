package twoSum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	var testCases = []struct {
		nums   []int
		target int
		want   [2]int
	}{
		{
			nums:   []int{1, 2, 3, 4},
			target: 3,
			want:   [2]int{0, 1},
		},
		{
			nums:   []int{0, 0, 3, 4},
			target: 3,
			want:   [2]int{0, 2},
		},
		{
			nums:   []int{1, 1, 3, 4},
			target: 5,
			want:   [2]int{0, 3},
		},
		{
			nums:   []int{4, 2, 3, 4},
			target: 7,
			want:   [2]int{0, 2},
		},
	}

	for _, test := range testCases {
		result := twoSum(test.nums, test.target)
		if result != test.want {
			t.Errorf("%v,target %v => %v, want %v", test.nums, test.target, result, test.want)
		}
	}
}

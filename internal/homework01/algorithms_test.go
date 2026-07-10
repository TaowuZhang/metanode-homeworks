package homework01

import (
	"reflect"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	if got := SingleNumber([]int{4, 1, 2, 1, 2}); got != 4 {
		t.Fatalf("got %d", got)
	}
	if got := SingleNumber([]int{-1}); got != -1 {
		t.Fatalf("single negative got %d", got)
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := map[int]bool{121: true, -121: false, 10: false, 0: true, 1221: true}
	for in, want := range cases {
		if got := IsPalindrome(in); got != want {
			t.Fatalf("%d got %v want %v", in, got, want)
		}
	}
}

func TestIsValidParentheses(t *testing.T) {
	cases := map[string]bool{"()[]{}": true, "([{}])": true, "(]": false, "([)]": false, "": true}
	for in, want := range cases {
		if got := IsValidParentheses(in); got != want {
			t.Fatalf("%q got %v want %v", in, got, want)
		}
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{nil, ""},
		{[]string{"same"}, "same"},
	}
	for _, tc := range cases {
		if got := LongestCommonPrefix(tc.in); got != tc.want {
			t.Fatalf("%v got %q want %q", tc.in, got, tc.want)
		}
	}
}

func TestPlusOne(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{9}, []int{1, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
	}
	for _, tc := range cases {
		if got := PlusOne(tc.in); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("%v got %v want %v", tc.in, got, tc.want)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 2, 3, 3}
	n := RemoveDuplicates(nums)
	if n != 4 || !reflect.DeepEqual(nums[:n], []int{0, 1, 2, 3}) {
		t.Fatalf("got n=%d nums=%v", n, nums[:n])
	}
	if got := RemoveDuplicates(nil); got != 0 {
		t.Fatalf("empty got %d", got)
	}
}

func TestMergeIntervals(t *testing.T) {
	got := MergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	want := [][]int{{1, 6}, {8, 10}, {15, 18}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
	got = MergeIntervals([][]int{{1, 4}, {4, 5}})
	want = [][]int{{1, 5}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("touching got %v want %v", got, want)
	}
}

func TestTwoSum(t *testing.T) {
	got := TwoSum([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual(got, []int{0, 1}) {
		t.Fatalf("got %v", got)
	}
	if got := TwoSum([]int{1, 2, 3}, 7); got != nil {
		t.Fatalf("no solution got %v", got)
	}
}

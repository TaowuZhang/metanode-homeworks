package homework01

import "sort"

func SingleNumber(nums []int) int {
	counts := make(map[int]int, len(nums))
	for _, n := range nums {
		counts[n]++
	}
	for _, n := range nums {
		if counts[n] == 1 {
			return n
		}
	}
	return 0
}

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	origin, reversed := x, 0
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return origin == reversed
}

func IsValidParentheses(s string) bool {
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := make([]rune, 0, len(s))
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}
	return len(stack) == 0
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, s := range strs[1:] {
		for len(prefix) > 0 && (len(s) < len(prefix) || s[:len(prefix)] != prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

func PlusOne(digits []int) []int {
	out := append([]int(nil), digits...)
	for i := len(out) - 1; i >= 0; i-- {
		if out[i] < 9 {
			out[i]++
			return out
		}
		out[i] = 0
	}
	return append([]int{1}, out...)
}

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{append([]int(nil), intervals[0]...)}
	for _, interval := range intervals[1:] {
		last := merged[len(merged)-1]
		if interval[0] <= last[1] {
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
			continue
		}
		merged = append(merged, append([]int(nil), interval...))
	}
	return merged
}

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}

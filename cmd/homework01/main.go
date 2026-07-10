package main

import (
	"fmt"

	"github.com/TaowuZhang/metanode-go-homeworks/internal/homework01"
)

func main() {
	fmt.Println("single number:", homework01.SingleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println("is palindrome:", homework01.IsPalindrome(1221))
	fmt.Println("valid parentheses:", homework01.IsValidParentheses("([{}])"))
	fmt.Println("longest common prefix:", homework01.LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println("plus one:", homework01.PlusOne([]int{9, 9, 9}))
	nums := []int{1, 1, 2, 3, 3}
	n := homework01.RemoveDuplicates(nums)
	fmt.Println("remove duplicates:", nums[:n])
	fmt.Println("merge intervals:", homework01.MergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}}))
	fmt.Println("two sum:", homework01.TwoSum([]int{2, 7, 11, 15}, 9))
}

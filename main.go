package main

import (
	"fmt"
	"linkedlist"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	// 49字母异位词分组
	var ans [][]string
	dict := make(map[string][]string)
	for _, s := range strs {
		t := []byte(s)
		slices.Sort(t)
		key := string(t)
		dict[key] = append(dict[key], s)
	}
	for _, cur := range dict {
		ans = append(ans, cur)
	}
	return ans
}

func main() {
	pre := &linkedlist.ListNode{Val: 666}
	fmt.Println(pre.Val)
}

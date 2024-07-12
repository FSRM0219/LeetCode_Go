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

func findAnagrams(s string, p string) []int {
	ls, lp := len(s), len(p)
	if ls < lp {
		return []int{}
	}
	var ans []int
	var cnts, cntp [26]int
	for i, c := range p {
		cntp[c-'a']++
		cnts[s[i]-'a']++
	}
	if cntp == cnts {
		ans = append(ans, 0)
	}
	for i := 0; i < ls-lp; i++ {
		cnts[s[i]-'a']--
		cnts[s[i+lp]-'a']++
		if cntp == cnts {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func main() {
	pre := &linkedlist.ListNode{Val: 666}
	fmt.Println(pre.Val)
}

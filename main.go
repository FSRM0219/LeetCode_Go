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

func countPrimes(n int) int {
	var ans int
	isPrime := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			ans++
		}
		for j := 2 * i; j < n; j += i {
			isPrime[j] = false
		}
	}
	return ans
}

func main() {
	pre := &linkedlist.ListNode{Val: 666}
	fmt.Println(pre.Val)
}

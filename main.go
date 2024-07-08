package main

import (
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

func longestConsecutive(nums []int) int {
	var ans int
	dict := make(map[int]bool)
	for _, num := range nums {
		dict[num] = true
	}
	for _, num := range nums {
		cnt := 0
		if dict[num] {
			cnt++
			dict[num] = false
			for cur := num - 1; dict[cur]; cur-- {
				dict[cur] = false
				cnt++
			}
			for cur := num + 1; dict[cur]; cur++ {
				dict[cur] = false
				cnt++
			}
		}
		ans = max(ans, cnt)
	}
	return ans
}

func main() {

}

package main

import (
	"sort"
	"strings"
)

func makeKey(val string) string {
	letters := strings.Split(val, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func checkSlice(s string, sl []string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}
	return false
}

func sortMap(NSmap map[string][]string) map[string][]string {
	result := make(map[string][]string)
	for _, v := range NSmap {
		result[v[0]] = v
		sort.Strings(v)
	}
	return result
}

func FindQuantities(s []string) map[string][]string {
	result := make(map[string][]string)
	var falseKeys []string
	for i := 0; i < len(s); i++ {
		val := strings.ToLower(s[i])
		val = makeKey(val)
		if checkSlice(val, falseKeys) {
			result[val] = append(result[val], s[i])
		} else {
			falseKeys = append(falseKeys, val)
			result[val] = []string{s[i]}
		}
	}
	result = sortMap(result)
	return result
}

package strutil

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func difference(a, b string) string {
	mb := make(map[rune]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = diff + string(x)
		}
	}
	return diff
}

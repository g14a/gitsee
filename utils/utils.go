package utils

import (
	"sort"
)

type SortMap struct {
	Key   string
	Value int
}

type ArraySortMap []SortMap

func (a ArraySortMap) Len() int {
	return len(a)
}

func (a ArraySortMap) Less(i, j int) bool {
	return a[i].Value < a[j].Value
}

func (a ArraySortMap) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func GetSortedMap(m map[string]int) map[string]int {
	var sortMap ArraySortMap
	
	for k, v := range m {
		sortMap = append(sortMap, SortMap{
			Key:   k,
			Value: v,
		})
	}

	sort.Sort(sort.Reverse(sortMap))

	resultSortedMap := make(map[string]int)

	i := 0
	for _, elem := range sortMap {
		if i <= 10 {
			resultSortedMap[elem.Key] = elem.Value
			i++
		} else {
			break
		}
	}

	return resultSortedMap
}

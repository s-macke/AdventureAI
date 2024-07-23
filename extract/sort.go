package main

import "sort"

type kvSliceStruct struct {
	k string
	v float64
}

func SortMap(data map[string]*StoryProgress) []string {
	var kvSlice []kvSliceStruct

	for k, v := range data {
		kvSlice = append(kvSlice, kvSliceStruct{k, v.GetAverage(len(v.avg) - 1)})
	}

	// Sort the slice based on the value
	sort.Slice(kvSlice, func(i, j int) bool {
		return kvSlice[i].v > kvSlice[j].v
	})
	var result []string
	for _, kv := range kvSlice {
		result = append(result, kv.k)
	}
	return result
}

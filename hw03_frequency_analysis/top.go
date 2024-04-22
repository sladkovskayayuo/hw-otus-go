package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	if str == "" {
		return []string{}
	}
	collection := strings.Fields(str)
	m := make(map[string]int)
	for _, word := range collection {
		m[word]++
	}

	type kv struct {
		Key   string
		Value int
	}

	var sm []kv
	for k, v := range m {
		sm = append(sm, kv{k, v})
	}

	sort.Slice(sm, func(i, j int) bool {
		if sm[i].Value == sm[j].Value {
			return sm[i].Key < sm[j].Key
		}
		return sm[i].Value > sm[j].Value
	})

	var keys []string
	for i, v := range sm {
		if i < 10 || sm[9].Value == sm[i].Value {
			keys = append(keys, v.Key)
		} else {
			break
		}
	}

	return keys[:10]
}

package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var keys []int

	for k, _ := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		result = append(result, input[k])
	}

	return
}

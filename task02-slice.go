package homework

func reverse(input []int64) (result []int64) {
	for i := 0; i < len(input); i++ {
		result = append(result, int64(input[len(input)-1-i]))
	}

	return
}

package tools

func Counter(word string) map[string]int {
	count_word := make(map[string]int)
	for _, value := range word {
		count_word[string(value)]++
	}
	return count_word
}
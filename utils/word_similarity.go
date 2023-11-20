package utils

func Word_similarity_Levenshtein(word1 string, word2 string) int {
	pre := make([]int, len(word2)+1)
	cur := make([]int, len(word2)+1)

	for i := 1; i <= len(word1); i++ {
		cur[0] = 1
		for j := 1; j < len(pre); j++ {
			if word1[i-1] != word2[j-1] {
				cur[j] = min(cur[j-1], pre[j-1], pre[j]) + 1
			} else {
				cur[j] = pre[j-1]
			}
		}

		tmp := make([]int, len(cur))
		copy(tmp, cur)
		pre = tmp
	}

	return pre[len(word2)]
}

func Distance_words_ratio(word1 string, word2 string) float64 {
	distance := Word_similarity_Levenshtein(word1, word2)
	maxLen := max(len(word1), len(word2))

	return 1 - float64(distance)/float64(maxLen)
}

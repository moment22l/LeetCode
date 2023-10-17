package main

import (
	"sort"
	"strings"
)

func topStudents(positive_feedback []string, negative_feedback []string,
	report []string, student_id []int, k int) []int {
	words := make(map[string]int)
	for _, word := range positive_feedback {
		words[word] = 3
	}
	for _, word := range negative_feedback {
		words[word] = -1
	}
	type pair struct {
		score, id int
	}
	scores := make([]pair, len(student_id))
	for i, sentence := range report {
		scores[i].id = student_id[i]
		bs := strings.Split(sentence, " ")
		for _, word := range bs {
			if score, ok := words[word]; ok {
				scores[i].score += score
			}
		}
	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score || (scores[i].score == scores[j].score && scores[i].id < scores[j].id)
	})
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, scores[i].id)
	}
	return ans
}

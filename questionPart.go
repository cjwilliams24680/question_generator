package main

import "math/rand"

type QuestionPart struct {
	variants []string
}

func (q *QuestionPart) getRandomVariant() string {
	length := len(q.variants)
	return q.variants[rand.Intn(length)]
}

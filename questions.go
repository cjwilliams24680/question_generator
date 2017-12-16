package main

import "bytes"

type Question struct {
	parts []*QuestionPart
}

func (q *Question) generateQuestion() string {
	var buffer bytes.Buffer

	for i, part := range q.parts {
		if i > 0 {
			buffer.WriteString(" ")
		}

		buffer.WriteString(part.getRandomVariant())
	}

	buffer.WriteString("?")

	return buffer.String()
}

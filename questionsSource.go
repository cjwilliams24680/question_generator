package main

type QuestionSource struct {

}

var howDidYou = &QuestionPart{[]string{"How did you", "Why did you", "What made you"}}

var getInterestedIn = &QuestionPart{[]string{"get interested in", "choose", "become interested in"}}

var yourResearch = &QuestionPart{[]string{"your research", "your research area", "your dissertation topic"}}

var psychology = &QuestionPart{[]string{"psychology", "therapy", "pediatric psychology"}}

var whatWouldYouDoIf = &QuestionPart{[]string{"What would you do if", "How would you respond if", "How would you handle a situation where"}}

var what = &QuestionPart{[]string{"In your opinion, what", "What"}}

var areYourStrengthsAsACandidate = &QuestionPart{[]string{"is your main strength", "are your greatest strengths", "are your top 3 best attributes as a candidate", "is the biggest thing you'd bring to the team"}}

var areYourWeaknessesAsACandidate = &QuestionPart{[]string{"is your main weakness", "are your greatest weaknesses", "is an area you'd like to improve in during internship", "would you need to most help with from your supervisor"}}



func (source *QuestionSource) buildQuestionsList() []*Question {
	var questions []*Question

	questions = append(questions, &Question{[]*QuestionPart{howDidYou, getInterestedIn, yourResearch}})

	questions = append(questions, &Question{[]*QuestionPart{howDidYou, getInterestedIn, psychology}})

	questions = append(questions, &Question{[]*QuestionPart{what, areYourStrengthsAsACandidate}})

	questions = append(questions, &Question{[]*QuestionPart{what, areYourWeaknessesAsACandidate}})

	return questions
}

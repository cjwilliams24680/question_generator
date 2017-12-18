package main

var howDidYou = &QuestionPart{[]string{"How did you", "Why did you", "What made you"}}

var getInterestedIn = &QuestionPart{[]string{"get interested in", "choose", "become interested in"}}

var yourResearch = &QuestionPart{[]string{"your research", "your research area", "your dissertation topic"}}

var psychology = &QuestionPart{[]string{"psychology", "therapy", "pediatric psychology"}}

var whatWouldYouDoIf = &QuestionPart{[]string{"What would you do if", "How would you respond if", "How would you handle a situation where"}}

var what = &QuestionPart{[]string{"In your opinion, what", "What"}}

var areYourStrengthsAsACandidate = &QuestionPart{[]string{"is your main strength", "are your greatest strengths", "are your top 3 best attributes as a candidate", "is the biggest thing you'd bring to the team"}}

var areYourWeaknessesAsACandidate = &QuestionPart{[]string{"is your main weakness", "are your greatest weaknesses", "is an area you'd like to improve in during internship", "would you need to most help with from your supervisor"}}

var whatIs = &QuestionPart{[]string{"What is"}}

var aPsychologists = &QuestionPart{[]string{"a psychologist's", "a therapist's"}}

var role = &QuestionPart{[]string{"role", "duty", "responsiblity"}}

var onATeam = &QuestionPart{[]string{"on a multidisciplinary team", "within their multidisciplinary team"}}

var listYour = &QuestionPart{[]string{"Please describe your", "List your", "What are your"}}

var careerGoals = &QuestionPart{[]string{"career goals", "internship goals", "goals for the next year? How about the next 10"}}

var defines = &QuestionPart{[]string{"defines", "makes"}}

var aGood = &QuestionPart{[]string{"a good", "your ideal", "the best possible", "great"}}

var supervisor = &QuestionPart{[]string{"supervisor", "mentor", "teacher"}}

var internshipSite = &QuestionPart{[]string{"site", "internship", "internship site"}}

var describe = &QuestionPart{[]string{"Describe", "Talk about", "Walk through"}}

var aCase = &QuestionPart{[]string{"a case of yours", "one of your patients", "a treatment case"}}

var thatWasUnsuccessful = &QuestionPart{[]string{"where treatment was unsuccessful.", "in which treatment didn't go well.", "that didn't meet treatment goals."}}

var howCouldItHaveGoneDifferently = &QuestionPart{[]string{"How could it have gone differently", "Why did it not go well", "Why did it end that way", "How did you respond to the situation"}}

func buildQuestionsList() []*Question {
	var questions []*Question

	questions = append(questions, &Question{[]*QuestionPart{howDidYou, getInterestedIn, yourResearch}})

	questions = append(questions, &Question{[]*QuestionPart{howDidYou, getInterestedIn, psychology}})

	questions = append(questions, &Question{[]*QuestionPart{what, areYourStrengthsAsACandidate}})

	questions = append(questions, &Question{[]*QuestionPart{what, areYourWeaknessesAsACandidate}})

	questions = append(questions, &Question{[]*QuestionPart{whatIs, aPsychologists, role, onATeam}})

	questions = append(questions, &Question{[]*QuestionPart{listYour, careerGoals}})

	questions = append(questions, &Question{[]*QuestionPart{what, defines, aGood, supervisor}})

	questions = append(questions, &Question{[]*QuestionPart{what, defines, aGood, internshipSite}})

	questions = append(questions, &Question{[]*QuestionPart{describe, aCase, thatWasUnsuccessful, howCouldItHaveGoneDifferently}})

	return questions
}

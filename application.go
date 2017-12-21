package main

import (
	"time"
	"math/rand"
	"net/http"
	"fmt"
	"bytes"
	"log"
)

var questions []*Question
var numberOfQuestions int

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	questions = buildQuestionsList()

	// Since the questions can't change during runtime, I can count it once and leave it.
	// If I decide to allow runtime modifications, I'll need to revisit this however
	numberOfQuestions = len(questions)

	http.HandleFunc("/", generateQuestion)

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}

	//todo add another one where they can specify an internship site
}

func generateQuestion(w http.ResponseWriter, r *http.Request) {
	index := rand.Intn(numberOfQuestions)
	fmt.Fprintln(w, questions[index].generateQuestion())
}

type QuestionPart struct {
	variants []string
}

func (q *QuestionPart) getRandomVariant() string {
	length := len(q.variants)
	return q.variants[rand.Intn(length)]
}

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

var howDidYou = &QuestionPart{[]string{"How did you", "Why did you", "What made you"}}

var getInterestedIn = &QuestionPart{[]string{"get interested in", "choose", "become interested in"}}

var yourResearch = &QuestionPart{[]string{"your research", "your research area", "your dissertation topic"}}

var psychology = &QuestionPart{[]string{"psychology", "therapy", "pediatric psychology"}}

var whatWouldYouDoIf = &QuestionPart{[]string{"What would you do if", "How would you respond if", "How would you handle a situation where"}}

var what = &QuestionPart{[]string{"In your opinion, what", "What"}}

var areYourStrengthsAsACandidate = &QuestionPart{[]string{"is your main strength", "are your greatest strengths", "are your top 3 best attributes as a candidate", "is the biggest thing you'd bring to the team"}}

var areYourWeaknessesAsACandidate = &QuestionPart{[]string{"is your main weakness", "are your greatest weaknesses", "is an area you'd like to improve in during internship", "would you need the most help with from your supervisor"}}

var whatIs = &QuestionPart{[]string{"What is"}}

var aPsychologists = &QuestionPart{[]string{"a psychologist's", "a therapist's"}}

var role = &QuestionPart{[]string{"role", "duty", "responsiblity"}}

var onATeam = &QuestionPart{[]string{"on a multidisciplinary team", "within their multidisciplinary team"}}

var listYour = &QuestionPart{[]string{"Please describe your", "List your", "What are your"}}

var careerGoals = &QuestionPart{[]string{"career goals", "internship goals", "goals for the next year? How about the next 10"}}

var defines = &QuestionPart{[]string{"defines", "makes"}}

var aGood = &QuestionPart{[]string{"a good", "your ideal", "the best possible", "a great", "a positive"}}

var aBad = &QuestionPart{[]string{"a bad", "a poor", "a negative"}}

var supervisor = &QuestionPart{[]string{"supervisor", "mentor", "teacher"}}

var internshipSite = &QuestionPart{[]string{"internship", "internship site"}}

var describe = &QuestionPart{[]string{"Describe", "Talk about", "Walk through", "Tell me about"}}

var aCase = &QuestionPart{[]string{"a case of yours", "one of your patients", "a treatment case"}}

var thatWasUnsuccessful = &QuestionPart{[]string{"where treatment was unsuccessful.", "in which treatment didn't go well.", "that didn't meet treatment goals."}}

var howCouldItHaveGoneDifferently = &QuestionPart{[]string{"How could it have gone differently", "Why did it not go well", "Why did it end that way", "How did you respond to the situation"}}

var anEthicalDilemna = &QuestionPart{[]string{"an ethical dilemna", "an moral problem"}}

var thatYouFaced = &QuestionPart{[]string{"that you faced.", "which you ran into.", "that you had to deal with."}}

var howDidYouHandleIt = &QuestionPart{[]string{"How did you handle it"}}

var underWhatConditions = &QuestionPart{[]string{"Under what conditions", "In what situation"}}

var can = &QuestionPart{[]string{"can", "should"}}

var aPsychologist = &QuestionPart{[]string{"a psychologist", "psychologists", "a therapist", "therapists"}}

var breakConfidentiality = &QuestionPart{[]string{"break confidentiality"}}

var anAssessment = &QuestionPart{[]string{"an assessment", "an instrument", "a test"}}

var withWhichYouFeelCompetent = &QuestionPart{[]string{"with which you feel competent", "which you can confidently use", "with which you extremely confident", "which you are confident using"}}

var whatIsYourOpinion = &QuestionPart{[]string{"What is you opinion"}}

var onProjectiveTests = &QuestionPart{[]string{"on projective tests", "of dynamic approaches"}} //todo are there any other things which would also be applicable to these questions?

var psychologicalAssessments = &QuestionPart{[]string{"psychological assessments", "psychological tests", "therapeutic assessments", "empirically validated treatments"}}

var areYouFamiliarWith = &QuestionPart{[]string{"are you familiar with", "are you confident in"}}

var whatTypeOfClientIsMostDifficultForYouToWorkWithWhy = &QuestionPart{[]string{"What type of client is most difficult for you to work with? Why?", "What type of client is most difficult for you to work with? What type of feelings do you have towards such clients? How do these feelings interfere with treatment?"}}

var yourOrientation = &QuestionPart{[]string{"your orientation", "your preferred orientation", "your favorite orientation"}}

var inTherapy = &QuestionPart{[]string{"in therapy"}}

var thatWasDifficult = &QuestionPart{[]string{"that was difficult", "which was challenging"}}

var supervisoryExperience = &QuestionPart{[]string{"supervisory experience"}}

var about = &QuestionPart{[]string{"about", "on"}}

var psychologists = &QuestionPart{[]string{"psychologists", "therapists"}}

var havingPrescriptionPrivileges = &QuestionPart{[]string{"havingPrescriptionPrivileges"}}

var whatKindsOf = &QuestionPart{[]string{"What kinds of", "What types of"}}

var supervisors = &QuestionPart{[]string{"supervisors", "mentors", "teachers"}}

var practicums = &QuestionPart{[]string{"practicums", "practicum sites", "training sites"}}

var haveYouHad = &QuestionPart{[]string{"have you had?"}}

var whatWorkedWell = &QuestionPart{[]string{"What worked well and what didn't", "Which was your most and least favorite", "What was the best and worst"}}

var haveYouHadAnyExperienceIn = &QuestionPart{[]string{"Have you had any experience in", "Do you have any exposure with", "Have you done", "Have you participated in"}}

var groupTherapy = &QuestionPart{[]string{"with group therapy", "with co therapy"}}

var yourTherapeuticStyle = &QuestionPart{[]string{"your therapeutic style"}}

var whyDidYouChoose = &QuestionPart{[]string{"Why did you choose", "What made you pick"}}

var toEnrollIn = &QuestionPart{[]string{"to enroll in", "to go to", "to study at", "to get your degree from"}}

var yourGraduateProgram = &QuestionPart{[]string{"your graduate program", "your school", "your training site"}}

var whatAreTheLimitationsOf = &QuestionPart{[]string{"What are some limitations of", "What are some weaknesses of"}}

var whatAreTheStrengthsOf = &QuestionPart{[]string{"What are some strengths of"}}

var whatTypesOfClients = &QuestionPart{[]string{"What types of clients", "What kinds of patients"}}

var areYouMostComfortableWithWhy = &QuestionPart{[]string{"are you most comfortable with. Why", "do you have the most familiarity with? Why"}}

var describeAMinorityGroupWithWhichYouHaveALotOfExperienceWorkingWith = &QuestionPart{[]string{"What is a minority group with which you have a lot of experience working with?"}}

var whatAreSomeChallenges = &QuestionPart{[]string{"What are some challenges you've faced when working with this group"}}

var howDoYouWorkWithAndUnderstandPatientsWithDifferentBackgrounds = &QuestionPart{[]string{"How do you work with and understand patients with different backgrounds?"}}

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

	questions = append(questions, &Question{[]*QuestionPart{describe, anEthicalDilemna, thatYouFaced, howDidYouHandleIt}})

	questions = append(questions, &Question{[]*QuestionPart{underWhatConditions, can, aPsychologist, breakConfidentiality}})

	questions = append(questions, &Question{[]*QuestionPart{describe, anAssessment, withWhichYouFeelCompetent}})

	questions = append(questions, &Question{[]*QuestionPart{whatIsYourOpinion, onProjectiveTests}})

	questions = append(questions, &Question{[]*QuestionPart{what, psychologicalAssessments, areYouFamiliarWith}})

	questions = append(questions, &Question{[]*QuestionPart{whatTypeOfClientIsMostDifficultForYouToWorkWithWhy}})

	questions = append(questions, &Question{[]*QuestionPart{whatIs, yourOrientation, inTherapy}})

	questions = append(questions, &Question{[]*QuestionPart{describe, aCase, thatWasDifficult}})

	questions = append(questions, &Question{[]*QuestionPart{describe, aGood, supervisoryExperience}})

	questions = append(questions, &Question{[]*QuestionPart{describe, aBad, supervisoryExperience}})

	questions = append(questions, &Question{[]*QuestionPart{whatIsYourOpinion, about, psychologists, havingPrescriptionPrivileges}})

	questions = append(questions, &Question{[]*QuestionPart{whatKindsOf, supervisors, haveYouHad, whatWorkedWell}})

	questions = append(questions, &Question{[]*QuestionPart{haveYouHadAnyExperienceIn, groupTherapy}})

	questions = append(questions, &Question{[]*QuestionPart{describe, yourTherapeuticStyle}})

	questions = append(questions, &Question{[]*QuestionPart{whyDidYouChoose, toEnrollIn, yourGraduateProgram}})

	questions = append(questions, &Question{[]*QuestionPart{whatAreTheLimitationsOf, yourGraduateProgram}})

	questions = append(questions, &Question{[]*QuestionPart{whatAreTheStrengthsOf, yourGraduateProgram}})

	questions = append(questions, &Question{[]*QuestionPart{whatTypesOfClients, areYouMostComfortableWithWhy}})

	questions = append(questions, &Question{[]*QuestionPart{describeAMinorityGroupWithWhichYouHaveALotOfExperienceWorkingWith, whatAreSomeChallenges}})

	questions = append(questions, &Question{[]*QuestionPart{howDoYouWorkWithAndUnderstandPatientsWithDifferentBackgrounds}})

	questions = append(questions, &Question{[]*QuestionPart{whatKindsOf, practicums, haveYouHad, whatWorkedWell}})

	return questions
}

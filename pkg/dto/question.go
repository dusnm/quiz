package dto

type Answer struct {
	Answer  string `form:"answer"`
	Correct bool   `form:"correct"`
}

type Question struct {
	Question string   `form:"question"`
	Answers  []Answer `form:"answer[]"`
}

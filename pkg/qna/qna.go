package qna

// Question structure
type Question struct {
	Body string `json:"body" binding:"required"`
}

// Answer structure
type Answer struct {
	Body string `json:"body" binding:"required"`
}

// Ask q question
func Ask(q Question) (Answer, error) {
	return Answer{
		Body: "I dont know " + q.Body,
	}, nil
}

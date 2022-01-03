package models

type Idf struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

// idf list

type Idfs []*Idf

func New() *Idf {
	return &Idf{
		ID:   "ABC00EFG",
		Body: "fname=test",
	}
}

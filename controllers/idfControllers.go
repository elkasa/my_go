package controllers

import "fmt"

// idfs
type Idf struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

type Idfs []*Idf

// ListIdf returm all idfs from the database limite 50 idfs tbd

func ListIdf() Idfs {
	return listIdf
}

func findIndexByIdf(id string) int {
	for i, p := range listIdf {
		if p.ID == id {
			return i
		}
	}

	return -1
}

// GetProducts returns all products from the database
func GetIdfs() Idfs {
	return listIdf
}

func GetIdfID(id string) (*Idf, error) {
	i := findIndexByIdf(id)
	if i == -1 {
		return nil, fmt.Errorf("Idf not found")
	}

	return listIdf[i], nil
}

// AddIdf adds a new idf to the database
func AddIdf(idf Idf) error {
	listIdf = append(listIdf, &idf)
	return nil
}

// UpdateIdf

func UpdateIdf(idf Idf) (*Idf, error) {
	i := findIndexByIdf(idf.ID)
	if i == -1 {
		return nil, fmt.Errorf("Idf not found")
	}
	listIdf[i] = &idf
	return listIdf[i], nil

}

func DeleteIdf(id string) (*Idf, error) {
	i := findIndexByIdf(id)
	if i == -1 {
		return nil, fmt.Errorf("Idf not found")
	}
	listIdf = append(listIdf[:i], listIdf[i+1:]...)
	return listIdf[i], nil
}

var listIdf = []*Idf{
	{
		ID:   "ABC00EFG",
		Body: "fname=test",
	},

	{
		ID:   "xyz00jkl",
		Body: "fname=test",
	},
}

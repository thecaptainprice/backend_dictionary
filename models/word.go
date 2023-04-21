package models

type Word struct {
	ID      uint64 `json:"id"`
	Word    string `json:"word"`
	Meaning string `json:"meaning"`
}

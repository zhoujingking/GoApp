package models

type Book struct {
	Name string `json:"name" binding:"required"`
	ISBN string `json:"isbn" binding:"required"`
}
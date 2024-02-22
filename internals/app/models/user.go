package models


type User struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Rank string `json:"rank"`
}
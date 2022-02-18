package models

type Token struct {
	Code         int    `json:"code"`
	UserID       string `json:"userId"`
	Token        string `json:"token"`
	ErrorMessage string `json:"errorMessage"`
}

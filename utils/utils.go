package utils

import "github.com/parnurzeal/gorequest"

func NewRequest(url string) {
	gorequest.New().Set("Content-Type", "application/json")
}

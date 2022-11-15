package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type Utils struct {
}

// Intersection find the same elements of two array
func Intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

// FindEmailFromText return email mentioned in text
func FindEmailFromText(text string) []string {

	regex := regexp.MustCompile(EmailValidationRegex)

	emailChain := regex.FindAllString(text, -1)

	emails := make([]string, len(emailChain))

	for index, emailCharacter := range emailChain {
		emails[index] = emailCharacter
	}
	return emails
}

// GetReceiverID get slice of receiver id
func GetReceiverID(a, b []int) (c []int) {

	sameElements := Intersection(a, b)

	for _, v := range sameElements {
		a = removeIndex(a, indexOf(v, a))
	}

	return a
}

// removeIndex remove index in slice
func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

// indexOf get index of value in slice
func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

// UniqueSlice remove duplicate element in slice
func UniqueSlice(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func LogRequest(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request path: %s", r.URL.Path)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

type RespError struct {
	Code    int
	Message string
}

type Response struct {
	Message string `json:"message"`
}

func ResponseJson(w http.ResponseWriter, httpCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(data)
}

func (r RespError) Error() string {
	return fmt.Sprintf("%d: %s", r.Code, r.Message)
}

func InternalServerError(msg string) error {
	return RespError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func JsonResponseError(w http.ResponseWriter, err error) {
	if _, ok := err.(RespError); !ok {
		err = InternalServerError(err.Error())
	}
	error, _ := err.(RespError)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(error.Code)
	res := Response{
		Message: error.Message,
	}
	json.NewEncoder(w).Encode(res)
}

package utils

import (
	"log"
	"net/http"
	"regexp"
)

type Utils struct {
}
type UtilsInf interface {
	Intersection(a, b []int) (c []int)
	FindEmailFromText(text string) []string
	GetReceiverID(a, b []int) (c []int)
	UniqueSlice(intSlice []int) []int
}

// Intersection find the same elements of two array
func (u Utils) Intersection(a, b []int) (c []int) {
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
func (u Utils) FindEmailFromText(text string) []string {

	regex := regexp.MustCompile(EmailValidationRegex)

	emailChain := regex.FindAllString(text, -1)

	emails := make([]string, len(emailChain))

	for index, emailCharacter := range emailChain {
		emails[index] = emailCharacter
	}
	return emails
}

// GetReceiverID get slice of receiver id
func (u Utils) GetReceiverID(a, b []int) (c []int) {

	sameElements := u.Intersection(a, b)

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
func (u Utils) UniqueSlice(intSlice []int) []int {
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

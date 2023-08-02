package utils

import (
	"errors"
	"regexp"

	"github.com/abdul-rehman-d/library-api/models"
)

func validateNameAndRegex(str string, regex string) bool {
	if len(str) < 3 {
		return false
	}

	return regexp.MustCompile(regex).MatchString(str)
}

func ValidateBookTitle(title string) bool {
	return validateNameAndRegex(title, `^[A-Za-z0-9\s\-_,\.;:()]+$`)
}

func ValidateAuthorName(name string) bool {
	return validateNameAndRegex(name, `^\s*([A-Za-z]{1,}([\.,] |[-']| ))+[A-Za-z]+\.?\s*$`)
}

func ValidateNewBook(book models.Book) error {
	if !ValidateBookTitle(book.Title) {
		return errors.New("invalid book title")
	}

	if !ValidateAuthorName(book.Author) {
		return errors.New("invalid author name")
	}

	return nil
}

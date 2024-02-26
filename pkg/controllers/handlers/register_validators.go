package handlers

import (
	"regexp"
	"time"

	"bitbucket.com/finease/backend/pkg/controllers/api"
)

type Validator func(user api.User) bool

func ValidateEmail(user api.User) bool {
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(user.Email)
}

func ValidateDOB(user api.User) bool {
	layout := "02/01/2006"
	_, err := time.Parse(layout, user.DateOfBirth)
	return err == nil
}

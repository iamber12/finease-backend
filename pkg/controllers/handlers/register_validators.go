package handlers

import (
	"bitbucket.com/finease/backend/pkg/controllers/api"
	"regexp"
)

type Validator func(user api.User) bool

func ValidateEmail(user api.User) bool {
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(user.Email)
}

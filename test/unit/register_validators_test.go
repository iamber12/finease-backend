package unit

import (
	"bitbucket.com/finease/backend/pkg/controllers/api"
	"bitbucket.com/finease/backend/pkg/controllers/handlers"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidationEmail(t *testing.T) {
	expectedValidEmailAddress := []string{
		"example@example.com",
		"example@mail.example.com",
		"example+tag@example.com",
		"example123@example.com",
		"example_name@example.com",
	}

	expectedInvalidEmailAddress := []string{
		"example.com",
		"example@.com",
		"example@",
		"example@@example.com",
		"",
	}

	sampleUser := api.User{}
	for _, email := range expectedValidEmailAddress {
		sampleUser.Email = email
		validEmail := handlers.ValidateEmail(sampleUser)
		assert.True(t, validEmail, fmt.Sprintf("Email being tested: %s", email))
	}

	for _, email := range expectedInvalidEmailAddress {
		sampleUser.Email = email
		validEmail := handlers.ValidateEmail(sampleUser)
		assert.False(t, validEmail, fmt.Sprintf("Email being tested: %s", email))
	}

}

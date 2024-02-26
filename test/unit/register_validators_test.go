package unit

import (
	"fmt"
	"testing"

	"bitbucket.com/finease/backend/pkg/controllers/api"
	"bitbucket.com/finease/backend/pkg/controllers/handlers"
	"github.com/stretchr/testify/assert"
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

func TestValidationDOB(t *testing.T) {
	expectedValidDOB := []string{
		"23/09/2022",
		"01/01/2000",
		"31/12/1999",
		"29/02/2004",
		"15/08/1947",
	}

	expectedInvalidDOB := []string{
		"1234567",
		"23=11-2022",
		"11/000/1122",
		"00/00/0000",
	}

	sampleUser := api.User{}
	for _, dob := range expectedValidDOB {
		sampleUser.DateOfBirth = dob
		validEmail := handlers.ValidateDOB(sampleUser)
		assert.True(t, validEmail, fmt.Sprintf("DOB being tested: %s", dob))
	}

	for _, dob := range expectedInvalidDOB {
		sampleUser.DateOfBirth = dob
		validEmail := handlers.ValidateDOB(sampleUser)
		assert.False(t, validEmail, fmt.Sprintf("DOB being tested: %s", dob))
	}

}

package mocks

import (
	"bitbucket.com/finease/backend/pkg/models"
)

var (
	SampleUuid     = "88dfb4b9-a14e-474f-810f-89778137a9f9"
	SampleName     = "alex"
	SampleDob      = "28/02/1998"
	SampleAddress  = "200 University Ave W, Ontario, Waterloo"
	SampleRole     = "lender"
	SampleEmail    = "alex.bob@gmail.com"
	SamplePassword = "foobar123"
	SampleActive   = true
	SampleUser     = models.User{
		Generic: models.Generic{
			Uuid: SampleUuid,
		},
		Name:        SampleName,
		DateOfBirth: SampleDob,
		Address:     SampleAddress,
		PrimaryRole: SampleRole,
		Email:       SampleEmail,
		Password:    SamplePassword,
		Active:      &SampleActive,
	}
)

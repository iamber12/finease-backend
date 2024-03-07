package integration

import (
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"bitbucket.com/finease/backend/test/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthRegister(t *testing.T) {
	userDaoMock := mocks.NewUserDaoMock()

	authService := services.NewAuthService(userDaoMock)

	ctx := context.Background()

	sampleUser := mocks.SampleUser
	// GOOD: valid user should register successfully
	_, err := authService.Register(ctx, sampleUser)
	assert.Nil(t, err)

	// BAD: should fail to register the same user
	_, err = authService.Register(ctx, sampleUser)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "existing user with the same email found")
}

func TestAuthLogin(t *testing.T) {
	userDaoMock := mocks.NewUserDaoMock()

	authService := services.NewAuthService(userDaoMock)

	ctx := context.Background()

	sampleUser := mocks.SampleUser
	_, err := authService.Register(ctx, sampleUser)
	assert.Nil(t, err)

	// BAD: should fail to login the user with non-existent email
	_, _, err = authService.Login(ctx, "some-email", "some-password")
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "user not found")

	// BAD: should fail to login the user with right email but wrong password
	_, _, err = authService.Login(ctx, sampleUser.Email, "wrong-password")
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "invalid password")

	token, loggedInUser, err := authService.Login(ctx, sampleUser.Email, sampleUser.Password)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, loggedInUser.Email, sampleUser.Email)
}

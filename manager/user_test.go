package manager_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	fuzz "github.com/google/gofuzz"
	"github.com/kanca-studio/palang/manager"
	"github.com/kanca-studio/palang/mock"
	"github.com/kanca-studio/palang/service/user"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mock.NewAuthService(ctrl)
	userService := mock.NewUserService(ctrl)

	t.Run("register user", func(t *testing.T) {
		userService.
			EXPECT().
			IdentifierTypeToConst(gomock.Eq("Username")).
			Return(user.Username)
		authService.
			EXPECT().
			HashPassword(gomock.Eq("secret")).
			Return("hashed", nil)
		userService.
			EXPECT().
			CreateUser(gomock.Eq(user.Username), gomock.Eq("ndaimen"), gomock.Eq("hashed")).
			Return(user.Model{}, nil)

		um := manager.NewUser(userService, authService)
		assert.NotNil(t, um)

		err := um.Register("Username", "ndaimen", "secret")
		assert.Nil(t, err)
	})

	t.Run("register user failed", func(t *testing.T) {
		userService.
			EXPECT().
			IdentifierTypeToConst(gomock.Eq("Username")).
			Return(user.Username)
		authService.
			EXPECT().
			HashPassword(gomock.Eq("secret")).
			Return("hashed", nil)
		userService.
			EXPECT().
			CreateUser(gomock.Eq(user.Username), gomock.Eq("ndaimen"), gomock.Eq("hashed")).
			Return(user.Model{}, errors.New("fail create user"))

		um := manager.NewUser(userService, authService)
		assert.NotNil(t, um)

		err := um.Register("Username", "ndaimen", "secret")
		assert.NotNil(t, err)
	})
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mock.NewAuthService(ctrl)
	userService := mock.NewUserService(ctrl)

	sampleUser := user.Model{}
	f := fuzz.New()
	f.Fuzz(&sampleUser)

	t.Run("login user", func(t *testing.T) {
		sampleUser.Verified = true
		userService.
			EXPECT().
			IdentifierTypeToConst(gomock.Eq("Username")).
			Return(user.Username)
		userService.
			EXPECT().
			GetUserByIdentifier(gomock.Eq(user.Username), gomock.Eq(sampleUser.Username)).
			Return(sampleUser, nil)
		authService.
			EXPECT().
			CheckPasswordHash(gomock.Eq("secret"), gomock.Eq(sampleUser.Password)).
			Return(true)
		authService.
			EXPECT().
			CreateToken(gomock.Eq(sampleUser.ID)).
			Return("token", nil)

		um := manager.NewUser(userService, authService)
		assert.NotNil(t, um)

		token, err := um.Login("Username", sampleUser.Username, "secret")
		assert.Nil(t, err)
		assert.Equal(t, token, "token")
	})
}

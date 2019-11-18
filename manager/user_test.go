package manager_test

import (
	"errors"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	fuzz "github.com/google/gofuzz"
	"github.com/kanca-studio/palang/manager"
	"github.com/kanca-studio/palang/mock"
	"github.com/kanca-studio/palang/service/user"
	"github.com/stretchr/testify/assert"
)

func TestUserManagerRegister(t *testing.T) {
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

func TestUserManagerLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mock.NewAuthService(ctrl)
	userService := mock.NewUserService(ctrl)

	sampleUser := user.Model{}
	f := fuzz.New()
	f.Fuzz(&sampleUser)

	t.Run("login user success", func(t *testing.T) {
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

	t.Run("login with get user error", func(t *testing.T) {
		sampleUser.Verified = true
		userService.
			EXPECT().
			IdentifierTypeToConst(gomock.Eq("Username")).
			Return(user.Username)
		userService.
			EXPECT().
			GetUserByIdentifier(gomock.Eq(user.Username), gomock.Eq(sampleUser.Username)).
			Return(sampleUser, errors.New("fail get user"))

		um := manager.NewUser(userService, authService)
		assert.NotNil(t, um)

		token, err := um.Login("Username", sampleUser.Username, "secret")
		assert.NotNil(t, err)
		assert.Equal(t, token, "")
	})

	t.Run("error user name or password", func(t *testing.T) {
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
			Return(false)

		um := manager.NewUser(userService, authService)
		assert.NotNil(t, um)

		token, err := um.Login("Username", sampleUser.Username, "secret")
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New("please check again username or password"))
		assert.Equal(t, token, "")
	})

	t.Run("user not verified error", func(t *testing.T) {
		sampleUser.Verified = false
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

		um := manager.NewUser(userService, authService)
		assert.NotNil(t, um)

		token, err := um.Login("Username", sampleUser.Username, "secret")
		assert.NotNil(t, err)
		assert.Equal(t, token, "")
	})

	t.Run("create token error", func(t *testing.T) {
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
			Return("", errors.New("fail create token"))

		um := manager.NewUser(userService, authService)
		assert.NotNil(t, um)

		token, err := um.Login("Username", sampleUser.Username, "secret")
		assert.NotNil(t, err)
		assert.Equal(t, token, "")
	})
}

func TestUserManagerValidateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mock.NewAuthService(ctrl)
	userService := mock.NewUserService(ctrl)

	t.Run("validate success", func(t *testing.T) {
		authService.
			EXPECT().
			ValidateToken(gomock.Eq("token")).
			Return(jwt.MapClaims{}, nil)

		um := manager.NewUser(userService, authService)
		assert.NotNil(t, um)

		err := um.ValidateToken("token")
		assert.Nil(t, err)
	})

	t.Run("validate fail", func(t *testing.T) {
		authService.
			EXPECT().
			ValidateToken(gomock.Eq("token")).
			Return(jwt.MapClaims{}, errors.New("fail validate token"))

		um := manager.NewUser(userService, authService)
		assert.NotNil(t, um)

		err := um.ValidateToken("token")
		assert.NotNil(t, err)
	})
}

func TestUserManagerMe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mock.NewAuthService(ctrl)
	userService := mock.NewUserService(ctrl)

	claims := jwt.MapClaims{
		"sub": float64(123),
	}
	dataUser := user.Model{}
	f := fuzz.New()
	f.Fuzz(&dataUser)

	t.Run("fetch me success", func(t *testing.T) {
		authService.
			EXPECT().
			ValidateToken(gomock.Eq("token")).
			Return(claims, nil)

		userService.
			EXPECT().
			FindById(gomock.Eq(uint(claims["sub"].(float64))), gomock.Any()).
			Return(nil)

		um := manager.NewUser(userService, authService)
		assert.NotNil(t, um)

		_, err := um.Me("token")
		assert.Nil(t, err)
	})

	t.Run("fetch me fail token", func(t *testing.T) {
		authService.
			EXPECT().
			ValidateToken(gomock.Eq("token")).
			Return(claims, errors.New("fail validate token"))

		um := manager.NewUser(userService, authService)
		assert.NotNil(t, um)

		_, err := um.Me("token")
		assert.NotNil(t, err)
	})

	t.Run("fail find by id", func(t *testing.T) {
		authService.
			EXPECT().
			ValidateToken(gomock.Eq("token")).
			Return(claims, nil)

		userService.
			EXPECT().
			FindById(gomock.Eq(uint(claims["sub"].(float64))), gomock.Any()).
			Return(errors.New("not found"))

		um := manager.NewUser(userService, authService)
		assert.NotNil(t, um)

		_, err := um.Me("token")
		assert.NotNil(t, err)
	})
}

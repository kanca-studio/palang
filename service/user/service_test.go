package user_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kanca-studio/palang/mock"
	"github.com/kanca-studio/palang/service/user"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	baserepo := mock.NewBaseRepository(ctrl)
	t.Run("create user using username success", func(t *testing.T) {
		baserepo.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil)
		baserepo.EXPECT().Create(gomock.Any()).Return(nil)
		repo := user.NewRepository(user.WithInterface(baserepo))
		service := user.NewService(repo)
		_, err := service.CreateUser(user.Username, "ndaimen", "secrethash")
		assert.Nil(t, err)
	})

	t.Run("create user using email success", func(t *testing.T) {
		baserepo.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil)
		baserepo.EXPECT().Create(gomock.Any()).Return(nil)
		repo := user.NewRepository(user.WithInterface(baserepo))
		service := user.NewService(repo)
		_, err := service.CreateUser(user.Email, "ndaimen@bajindul.com", "secrethash")
		assert.Nil(t, err)
	})

	t.Run("create user using phone number success", func(t *testing.T) {
		baserepo.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil)
		baserepo.EXPECT().Create(gomock.Any()).Return(nil)
		repo := user.NewRepository(user.WithInterface(baserepo))
		service := user.NewService(repo)
		_, err := service.CreateUser(user.PhoneNumber, "081988888", "secrethash")
		assert.Nil(t, err)
	})

	t.Run("create user got error", func(t *testing.T) {
		baserepo.EXPECT().Find(gomock.Any(), gomock.Any()).Return(errors.New("aw snap"))
		repo := user.NewRepository(user.WithInterface(baserepo))
		service := user.NewService(repo)
		_, err := service.CreateUser(user.PhoneNumber, "081988888", "secrethash")
		assert.NotNil(t, err)
	})

	t.Run("get username identifier", func(t *testing.T) {
		repo := user.NewRepository(user.WithInterface(baserepo))
		service := user.NewService(repo)
		it := service.IdentifierTypeToConst("Username")
		assert.Equal(t, it, user.Username)
	})

	t.Run("get PhoneNumber identifier", func(t *testing.T) {
		repo := user.NewRepository(user.WithInterface(baserepo))
		service := user.NewService(repo)
		it := service.IdentifierTypeToConst("PhoneNumber")
		assert.Equal(t, it, user.PhoneNumber)
	})

	t.Run("get Email identifier", func(t *testing.T) {
		repo := user.NewRepository(user.WithInterface(baserepo))
		service := user.NewService(repo)
		it := service.IdentifierTypeToConst("Email")
		assert.Equal(t, it, user.Email)
	})
}

package user

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"kanca-studio/palang/service/base"
	"time"
)

var JWTSECRET = "secret-kanca"

func NewService(repo repository) service {
	s := service{}
	s.BaseService.Repository = &repo
	s.repo = repo

	return s
}

type service struct {
	base.BaseService
	repo repository
}

func (s *service) Register(email, password string) (Model, error) {
	hash, _ := s.hashPassword(password)
	data := Model{
		Email:    email,
		Password: hash,
		Verified: false,
	}
	err := s.Create(data)
	return data, err
}

func (s *service) Login(email, password string) (interface{}, error) {
	var err error
	result, err := s.Find(Model{Email: email})
	if err != nil {
		return err, nil
	}
	user := result.(Model)
	if !s.checkPasswordHash(password, user.Password) {
		return nil, errors.New("please check again username or password")
	}

	if user.Verified == false {
		return nil, errors.New("User not verified")
	}
	token, err := s.createToken(user)
	if err != nil {
		return err, nil
	}

	data := map[string]interface{}{
		"token": token,
	}
	return data, nil
}

func (s *service) Activated(token string) (interface{}, error) {
	data := map[string]interface{}{
		"message": "dummy",
	}
	return data, nil
}

func validateToken(tokenString string) error {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWTSECRET), nil
	});
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	} else {
		return errors.New("token not valid")
	}

}

func (s *service) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (s *service) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *service) createToken(user Model) (t string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err = token.SignedString([]byte(JWTSECRET))
	return
}

package user

import (
	"kanca-studio/palang/database"
	"testing"
)

var s Service

func init() {
	database.Init("localhost", 5432, "posgtres", "posgtres", "kanca", true)
	repo := NewRepository(database.GetInstance())
	s = NewService(repo)

}

func TestNewService(t *testing.T) {
}

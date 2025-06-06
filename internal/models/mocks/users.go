package mocks

import (
	"github.com/H-ADJI/letsgo/internal/models"
)

type UserModel struct{}

func (u UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@email.fail":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}
func (u UserModel) Authenticate(email, password string) (int, error) {
	if email == "alice@example.com" && password == "pa$$word" {
		return 1, nil
	}
	return 0, models.ErrInvalidCreds

}
func (u UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

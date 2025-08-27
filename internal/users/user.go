package users

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tsydim/otus-highload-architect-hw/internal/apperrs"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type UserID = string

type UserFields struct {
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	BirthDate  time.Time `json:"birthdate"`
	Biography  string    `json:"biography"`
	City       string    `json:"city"`
	Gender     Gender    `json:"gender"`
}
type User struct {
	UserFields
	ID       UserID   `json:"id"`
	Password Password `json:"-"`
}

func NewUser(userFields UserFields, password Password) (User, error) {
	var zero User
	if len(userFields.FirstName) == 0 || len(userFields.SecondName) == 0 {
		return zero, fmt.Errorf("firstname or secondname are empty: %w", apperrs.ErrConditionViolation)
	}
	if len(password) == 0 {
		return zero, fmt.Errorf("password is empty: %w", apperrs.ErrConditionViolation)
	}

	return User{
		ID:         uuid.New().String(),
		Password:   password,
		UserFields: userFields,
	}, nil
}

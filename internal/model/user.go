package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic,omitempty"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

type UserFilter struct {
	Name        string
	Surname     string
	Patronymic  string
	AgeFrom     int
	AgeTo       int
	Gender      string
	Nationality string
	Limit       int
}

type UpdateUser struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic,omitempty"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

func (u *UpdateUser) Copy(user *User) {
	if u.Name != "" {
		user.Name = u.Name
	}

	if u.Surname != "" {
		user.Surname = u.Surname
	}

	if u.Patronymic != "" {
		user.Patronymic = u.Patronymic
	}

	if u.Age != 0 {
		user.Age = u.Age
	}

	if u.Gender != "" {
		user.Gender = u.Gender
	}

	if u.Nationality != "" {
		user.Nationality = u.Nationality
	}
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Id, validation.Min(0), validation.Required, is.Digit),
		validation.Field(u.Name, validation.Required, is.Alpha),
		validation.Field(u.Surname, validation.Required, is.Alpha),
		validation.Field(u.Patronymic, is.Alpha),
		validation.Field(u.Age, validation.Required, validation.Min(0), is.Digit),
		validation.Field(u.Gender, validation.Required, is.Alpha),
		validation.Field(u.Nationality, validation.Required, is.Alpha),
	)
}

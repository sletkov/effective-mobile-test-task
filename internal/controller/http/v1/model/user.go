package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

type CreateUser struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

type UpdateUser struct {
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

func (u *CreateUser) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required, validation.Length(1, 255), is.Alpha),
		validation.Field(&u.Surname, validation.Required, validation.Length(1, 255), is.Alpha),
		validation.Field(&u.Patronymic, validation.Length(0, 255), is.Alpha),
	)
}

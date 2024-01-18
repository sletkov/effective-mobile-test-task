package model

import (
	"fmt"
	"strconv"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic,omitempty"`
	Age         uint8  `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

type UserFilter struct {
	Name        string
	Surname     string
	Patronymic  string
	AgeFrom     uint8
	AgeTo       uint8
	Gender      string
	Nationality string
	Limit       uint8
}

type UpdateUser struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic,omitempty"`
	Age         uint8  `json:"age"`
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

func (u *UserFilter) GetFilterRequest() string {
	fmt.Println(u)
	slice := make([]string, 0)

	if u.Name != "" {
		str := "name = " + "'" + u.Name + "'"
		slice = append(slice, str)
	}

	if u.Surname != "" {
		str := "surname = " + "'" + u.Surname + "'"
		slice = append(slice, str)
	}

	if u.Patronymic != "" {
		str := "patronymic = " + "'" + u.Patronymic + "'"
		slice = append(slice, str)
	}

	if u.AgeFrom != 0 {
		itoa := strconv.Itoa(int(u.AgeFrom))
		str := "age >= " + itoa
		slice = append(slice, str)
	}

	if u.AgeTo != 0 {
		itoa := strconv.Itoa(int(u.AgeTo))
		str := "age <= " + itoa
		slice = append(slice, str)
	}

	if u.Gender != "" {
		str := "gender = " + "'" + u.Gender + "'"
		slice = append(slice, str)
	}

	if u.Nationality != "" {
		str := "nationality = " + "'" + u.Nationality + "'"
		slice = append(slice, str)
	}

	result := strings.Join(slice, " AND ")
	fmt.Println(result)
	return result
}

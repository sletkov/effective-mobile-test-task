package model

import (
	"net/url"
	"strconv"

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

func (u *User) Validate() error {

	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required, validation.Length(0, 255), is.Alpha),
		validation.Field(&u.Surname, validation.Required, validation.Length(0, 255), is.Alpha),
		validation.Field(&u.Patronymic, validation.Length(0, 255), is.Alpha),
		validation.Field(&u.Age, validation.Required, validation.Min(1), validation.Max(100)),
		validation.Field(&u.Gender, validation.Required, validation.In("male", "female")),
		validation.Field(&u.Nationality, validation.Required, validation.Length(2, 2), is.Alpha),
	)
}

func (u *CreateUser) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required, validation.Length(1, 255), is.Alpha),
		validation.Field(&u.Surname, validation.Required, validation.Length(1, 255), is.Alpha),
		validation.Field(&u.Patronymic, validation.Length(0, 255), is.Alpha),
	)
}

func (u *UpdateUser) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Length(0, 255), is.Alpha),
		validation.Field(&u.Surname, validation.Length(0, 255), is.Alpha),
		validation.Field(&u.Patronymic, validation.Length(0, 255), is.Alpha),
		validation.Field(&u.Age, validation.Min(1), validation.Max(100)),
		validation.Field(&u.Gender, validation.In("male", "female")),
		validation.Field(&u.Nationality, validation.Length(2, 2), is.Alpha),
	)
}

func (u *UserFilter) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Length(0, 255), is.Alpha),
		validation.Field(&u.Surname, validation.Length(0, 255), is.Alpha),
		validation.Field(&u.Patronymic, validation.Length(0, 255), is.Alpha),
		validation.Field(&u.AgeFrom, validation.Min(1), validation.Max(100)),
		validation.Field(&u.AgeTo, validation.Min(1), validation.Max(100)),
		validation.Field(&u.Gender, validation.In("male", "female")),
		validation.Field(&u.Nationality, validation.Length(2, 2), is.Alpha),
		validation.Field(&u.Limit, validation.Min(1), validation.Max(50)),
	)
}

func (u *UserFilter) FillFilters(filters url.Values) error {
	defaultLimit := 10

	for k, v := range filters {
		if len(v) > 0 {
			switch k {
			case "name":
				if name := v[0]; name != "" {
					u.Name = v[0]
				}
			case "surname":
				if surname := v[0]; surname != "" {
					u.Surname = surname
				}
			case "patronymic":
				if patronymic := v[0]; patronymic != "" {
					u.Patronymic = patronymic
				}
			case "age_from":
				if value, err := strconv.Atoi(v[0]); err == nil {
					u.AgeFrom = value
				}
			case "age_to":
				if value, err := strconv.Atoi(v[0]); err == nil {
					u.AgeTo = value
				}
			case "gender":
				if gender := v[0]; gender != "" {
					u.Gender = gender
				}
			case "nationality":
				if nationality := v[0]; nationality != "" {
					u.Nationality = nationality
				}
			case "limit":
				if value, err := strconv.Atoi(v[0]); err == nil {
					u.Limit = value
				}
			}
		}
	}

	if u.Limit <= 0 {
		u.Limit = defaultLimit
	}

	return nil
}

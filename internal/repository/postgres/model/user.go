package model

import (
	"errors"
	"strconv"
	"strings"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Surname     string `db:"surname"`
	Patronymic  string `db:"patronymic"`
	Age         int    `db:"age"`
	Gender      string `db:"gender"`
	Nationality string `db:"nationality"`
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

func (u *UserFilter) GetFilterRequest() string {
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
		itoa := strconv.Itoa(u.AgeFrom)
		str := "age >= " + itoa
		slice = append(slice, str)
	}

	if u.AgeTo != 0 {
		itoa := strconv.Itoa(u.AgeTo)
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

	return result
}

package model

type User struct {
	Id          int
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Gender      string
	Nationality string
}

type UpdateUser struct {
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Gender      string
	Nationality string
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

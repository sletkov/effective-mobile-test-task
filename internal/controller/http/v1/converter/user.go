package converter

import (
	"github.com/sletkov/effective-mobile-test-task/internal/controller/http/v1/model"
	"github.com/sletkov/effective-mobile-test-task/internal/domain"
)

// refactor
func ToCreateUserFromController(user *model.CreateUser) *domain.User {
	return &domain.User{
		Name:       user.Name,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
	}
}

func ToUpdateUserFromController(user *model.UpdateUser) *domain.UpdateUser {
	return &domain.UpdateUser{
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  user.Patronymic,
		Age:         user.Age,
		Gender:      user.Gender,
		Nationality: user.Nationality,
	}
}

func ToUserFromController(user *model.User) *domain.User {
	return &domain.User{
		Id:          user.Id,
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  user.Patronymic,
		Age:         user.Age,
		Gender:      user.Gender,
		Nationality: user.Nationality,
	}
}

func ToUserFilterFromController(userFilter *model.UserFilter) *domain.UserFilter {
	return &domain.UserFilter{
		Name:        userFilter.Name,
		Surname:     userFilter.Surname,
		Patronymic:  userFilter.Patronymic,
		AgeFrom:     userFilter.AgeFrom,
		AgeTo:       userFilter.AgeTo,
		Gender:      userFilter.Gender,
		Nationality: userFilter.Nationality,
		Limit:       userFilter.Limit,
	}
}

func ToUserFromService(user *domain.User) *model.User {
	return &model.User{
		Id:          user.Id,
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  user.Patronymic,
		Age:         user.Age,
		Gender:      user.Gender,
		Nationality: user.Nationality,
	}
}

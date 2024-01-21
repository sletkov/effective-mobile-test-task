package converter

import (
	"github.com/sletkov/effective-mobile-test-task/internal/controller/http/v1/model"
	serviceModel "github.com/sletkov/effective-mobile-test-task/internal/service/model"
)

func ToCreateUserFromController(user *model.CreateUser) *serviceModel.User {
	return &serviceModel.User{
		Name:       user.Name,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
	}
}

func ToUpdateUserFromController(user *model.UpdateUser) *serviceModel.User {
	return &serviceModel.User{
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  user.Patronymic,
		Age:         user.Age,
		Gender:      user.Gender,
		Nationality: user.Nationality,
	}
}

func ToUserFromController(user *model.User) *serviceModel.User {
	return &serviceModel.User{
		Id:          user.Id,
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  user.Patronymic,
		Age:         user.Age,
		Gender:      user.Gender,
		Nationality: user.Nationality,
	}
}

func ToUserFilterFromController(userFilter *model.UserFilter) *serviceModel.UserFilter {
	return &serviceModel.UserFilter{
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

func ToUserFromService(user *serviceModel.User) *model.User {
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

package converter

import (
	repoModel "github.com/sletkov/effective-mobile-test-task/internal/repository/postgres/model"
	"github.com/sletkov/effective-mobile-test-task/internal/service/model"
)

func ToUserFromService(user *model.User) *repoModel.User {
	return &repoModel.User{
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  user.Patronymic,
		Age:         user.Age,
		Gender:      user.Gender,
		Nationality: user.Nationality,
	}
}

func ToUserFilterFromService(userFilter *model.UserFilter) *repoModel.UserFilter {
	return &repoModel.UserFilter{
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

func ToUserFromRepo(user *repoModel.User) *model.User {
	return &model.User{
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  user.Patronymic,
		Age:         user.Age,
		Gender:      user.Gender,
		Nationality: user.Nationality,
	}
}

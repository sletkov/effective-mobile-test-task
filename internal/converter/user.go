package converter

import (
	"github.com/sletkov/effective-mobile-test-task/internal/domain"
	repoModel "github.com/sletkov/effective-mobile-test-task/internal/repository/postgres/model"
)

func ToUserFromService(user *domain.User) *repoModel.User {
	return &repoModel.User{
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  user.Patronymic,
		Age:         user.Age,
		Gender:      user.Gender,
		Nationality: user.Nationality,
	}
}

func ToUserFilterFromService(userFilter *domain.UserFilter) *repoModel.UserFilter {
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

func ToUserFromRepo(user *repoModel.User) *domain.User {
	return &domain.User{
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  user.Patronymic,
		Age:         user.Age,
		Gender:      user.Gender,
		Nationality: user.Nationality,
	}
}

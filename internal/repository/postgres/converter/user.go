package converter

import (
	"github.com/sletkov/effective-mobile-test-task/internal/domain"
	repoModel "github.com/sletkov/effective-mobile-test-task/internal/repository/postgres/model"
)

// Convert from repository User to model User
func ToUserFromRepo(user *repoModel.User) *domain.User {
	return &domain.User{
		Id:          user.Id,
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  user.Patronymic,
		Age:         user.Age,
		Nationality: user.Nationality,
	}
}

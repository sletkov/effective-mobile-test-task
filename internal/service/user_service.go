package service

import (
	"context"
	"net/http"

	"github.com/sletkov/effective-mobile-test-task/internal/converter"
	"github.com/sletkov/effective-mobile-test-task/internal/domain"
	utils "github.com/sletkov/effective-mobile-test-task/internal/pkg"
	repoModel "github.com/sletkov/effective-mobile-test-task/internal/repository/postgres/model"
)

//go:generate mockgen -source=user_service.go -destination=../repository/postgres/mocks/mock.go

type UserRepository interface {
	Get(ctx context.Context, userFilter *repoModel.UserFilter) ([]repoModel.User, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, u *repoModel.User) error
	Create(ctx context.Context, u *repoModel.User) (int, error)
	GetUserById(ctx context.Context, id int) (*repoModel.User, error)
}

type Transport interface {
	Get(ctx context.Context, url string) (*http.Response, error)
}

type UserService struct {
	repository UserRepository
	transport  Transport
}

func New(repository UserRepository, transport Transport) *UserService {
	return &UserService{
		repository: repository,
		transport:  transport,
	}
}

// Get all users with filters and limit
func (s *UserService) Get(ctx context.Context, userFilter *domain.UserFilter) ([]domain.User, error) {
	users := make([]domain.User, 0)

	repoUsers, err := s.repository.Get(ctx, converter.ToUserFilterFromService(userFilter))

	if err != nil {
		return nil, err
	}

	for _, u := range repoUsers {
		users = append(users, *converter.ToUserFromRepo(&u))
	}

	return users, nil
}

// Delete user by id
func (s *UserService) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

// Update user
func (s *UserService) Update(ctx context.Context, id int, u *domain.User) error {
	err := s.repository.Update(ctx, id, converter.ToUserFromService(u))

	if err != nil {
		return err
	}

	return nil
}

// Create new user
func (s *UserService) Create(ctx context.Context, u *domain.User) error {

	// Get response from 3rd-party api
	ageResponse, err := s.transport.Get(ctx, "https://api.agify.io/?name="+u.Name)

	if err != nil {
		return err
	}

	// Add age to user
	if err := utils.Agify(ageResponse, u); err != nil {
		return err
	}

	// Get response from 3rd-party api
	genderResponse, err := s.transport.Get(ctx, "https://api.genderize.io/?name="+u.Name)

	if err != nil {
		return err
	}

	// Add gender to user
	if err := utils.Genderize(genderResponse, u); err != nil {
		return err
	}

	// Get response from 3rd-party api
	nationalityResponse, err := s.transport.Get(ctx, "https://api.nationalize.io/?name="+u.Name)

	if err != nil {
		return err
	}

	// Add nationality to user
	if err := utils.Nationalize(nationalityResponse, u); err != nil {
		return err
	}

	// Save user into db
	_, err = s.repository.Create(ctx, converter.ToUserFromService(u))

	if err != nil {
		return nil
	}

	return nil
}

func (s *UserService) GetById(ctx context.Context, id int) (*domain.User, error) {
	user, err := s.repository.GetUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(user), nil
}

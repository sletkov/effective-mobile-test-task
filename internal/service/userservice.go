package service

import (
	"context"

	"github.com/sletkov/effective-mobile-test-task/internal/model"
	"github.com/sletkov/effective-mobile-test-task/internal/repository"
	"github.com/sletkov/effective-mobile-test-task/internal/transport"
)

type UserService struct {
	repository repository.UserRepository
	transport  transport.Transport
}

func New(repository repository.UserRepository, transport transport.Transport) *UserService {
	return &UserService{
		repository: repository,
		transport:  transport,
	}
}

func (s *UserService) Get(ctx context.Context, userFilter *model.UserFilter) ([]model.User, error) {
	users, err := s.repository.Get(ctx, userFilter)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete user by id
func (s *UserService) Delete(ctx context.Context, id uint) error {
	err := s.repository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

// Update user
func (s *UserService) Update(ctx context.Context, id uint, u *model.User) error {
	err := s.repository.Update(ctx, id, u)

	if err != nil {
		return err
	}

	return nil
}

// Create user
func (s *UserService) Create(ctx context.Context, u *model.User) error {

	// Get response from 3rd-party api
	ageResponse, err := s.transport.Get(ctx, "https://api.agify.io/?name="+u.Name)

	if err != nil {
		return err
	}

	// Add age to user
	if err := agify(ageResponse, u); err != nil {
		return err
	}

	// Get response from 3rd-party api
	genderResponse, err := s.transport.Get(ctx, "https://api.genderize.io/?name="+u.Name)

	if err != nil {
		return err
	}

	// Add gender to user
	if err := genderize(genderResponse, u); err != nil {
		return err
	}

	// Get response from 3rd-party api
	nationalityResponse, err := s.transport.Get(ctx, "https://api.nationalize.io/?name="+u.Name)

	if err != nil {
		return err
	}

	// Add nationality to user
	if err := nationalize(nationalityResponse, u); err != nil {
		return err
	}

	// Save user into db
	_, err = s.repository.Create(ctx, u)

	if err != nil {
		return nil
	}

	return nil
}

func (s *UserService) GetById(ctx context.Context, id uint) (*model.User, error) {
	user, err := s.repository.GetUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

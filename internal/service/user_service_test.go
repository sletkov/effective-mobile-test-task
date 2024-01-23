package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sletkov/effective-mobile-test-task/internal/converter"
	"github.com/sletkov/effective-mobile-test-task/internal/domain"
	mock_postgres "github.com/sletkov/effective-mobile-test-task/internal/repository/postgres/mocks"
	repoModel "github.com/sletkov/effective-mobile-test-task/internal/repository/postgres/model"
	mock_httptransport "github.com/sletkov/effective-mobile-test-task/internal/transport/http/mocks"
	"github.com/stretchr/testify/assert"
)

func TestServiceGet(t *testing.T) {
	type mockRepoBehavior func(r *mock_postgres.MockUserRepository, ctx context.Context, userFilter *repoModel.UserFilter)

	expectedUsers := []domain.User{
		{
			Id:          0,
			Name:        "Ivan",
			Surname:     "Ivanov",
			Patronymic:  "Ivanovich",
			Age:         20,
			Gender:      "male",
			Nationality: "RU",
		},

		{
			Id:          0,
			Name:        "Galina",
			Surname:     "Petrova",
			Patronymic:  "Petrovna",
			Age:         40,
			Gender:      "female",
			Nationality: "US",
		},
	}

	repoUsers := []repoModel.User{
		{
			Id:          0,
			Name:        "Ivan",
			Surname:     "Ivanov",
			Patronymic:  "Ivanovich",
			Age:         20,
			Gender:      "male",
			Nationality: "RU",
		},

		{
			Id:          0,
			Name:        "Galina",
			Surname:     "Petrova",
			Patronymic:  "Petrovna",
			Age:         40,
			Gender:      "female",
			Nationality: "US",
		},
	}

	testCases := []struct {
		name string
		mockRepoBehavior
		userFilter    *domain.UserFilter
		expectedUsers []domain.User
	}{
		{
			name: "OK",
			mockRepoBehavior: func(r *mock_postgres.MockUserRepository, ctx context.Context, userFilter *repoModel.UserFilter) {
				r.EXPECT().Get(ctx, userFilter).Return(repoUsers, nil)
			},
			userFilter: &domain.UserFilter{
				Limit: 10,
			},
			expectedUsers: expectedUsers,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_postgres.NewMockUserRepository(c)
			tc.mockRepoBehavior(repo, context.Background(), converter.ToUserFilterFromService(tc.userFilter))

			transport := mock_httptransport.NewMockTransport(c)

			service := New(repo, transport)

			users, err := service.Get(context.Background(), tc.userFilter)

			fmt.Println(users)
			fmt.Println(tc.expectedUsers)
			assert.NoError(t, err)
			assert.True(t, reflect.DeepEqual(tc.expectedUsers, users))
		})
	}
}
func TestServiceDelete(t *testing.T) {
	type mockRepoBehavior func(r *mock_postgres.MockUserRepository, ctx context.Context, id int)

	testCases := []struct {
		name string
		mockRepoBehavior
		id int
	}{
		{
			name: "OK",
			mockRepoBehavior: func(r *mock_postgres.MockUserRepository, ctx context.Context, id int) {
				r.EXPECT().Delete(ctx, id).Return(nil)
			},
			id: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_postgres.NewMockUserRepository(c)
			tc.mockRepoBehavior(repo, context.Background(), tc.id)

			transport := mock_httptransport.NewMockTransport(c)

			service := New(repo, transport)

			err := service.Delete(context.Background(), tc.id)

			assert.NoError(t, err)
		})
	}
}
func TestServiceUpdate(t *testing.T) {
	type mockRepoBehavior func(r *mock_postgres.MockUserRepository, ctx context.Context, id int, user *repoModel.User)

	testCases := []struct {
		name string
		mockRepoBehavior
		id   int
		user *repoModel.User
	}{
		{
			name: "OK",
			mockRepoBehavior: func(r *mock_postgres.MockUserRepository, ctx context.Context, id int, user *repoModel.User) {
				r.EXPECT().Update(ctx, id, user).Return(nil)
			},
			id: 1,
			user: &repoModel.User{
				Name: "Ivan",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_postgres.NewMockUserRepository(c)
			tc.mockRepoBehavior(repo, context.Background(), tc.id, tc.user)

			transport := mock_httptransport.NewMockTransport(c)

			service := New(repo, transport)

			err := service.Update(context.Background(), tc.id, converter.ToUserFromRepo(tc.user))

			assert.NoError(t, err)
		})
	}
}

func TestServiceGetById(t *testing.T) {
	type mockRepoBehavior func(r *mock_postgres.MockUserRepository, ctx context.Context, id int)

	expectedUser := &domain.User{
		Id:          0,
		Name:        "Ivan",
		Surname:     "Ivanov",
		Patronymic:  "Ivanovich",
		Age:         30,
		Gender:      "male",
		Nationality: "RU",
	}
	testCases := []struct {
		name string
		mockRepoBehavior
		id           int
		expectedUser *domain.User
	}{
		{
			name: "OK",
			mockRepoBehavior: func(r *mock_postgres.MockUserRepository, ctx context.Context, id int) {
				r.EXPECT().GetUserById(ctx, id).Return(converter.ToUserFromService(expectedUser), nil)
			},
			id:           0,
			expectedUser: expectedUser,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_postgres.NewMockUserRepository(c)
			tc.mockRepoBehavior(repo, context.Background(), tc.id)

			transport := mock_httptransport.NewMockTransport(c)

			service := New(repo, transport)

			user, err := service.GetById(context.Background(), tc.id)

			assert.NoError(t, err)
			assert.True(t, reflect.DeepEqual(tc.expectedUser, user))
		})
	}

}

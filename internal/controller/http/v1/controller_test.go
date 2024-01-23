package v1

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/sletkov/effective-mobile-test-task/internal/domain"
	mock_service "github.com/sletkov/effective-mobile-test-task/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestControllerHandleGetUsers(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUserService, ctx context.Context, userFilter *domain.UserFilter)

	users := []domain.User{
		{
			Id:          1,
			Name:        "Ivan",
			Surname:     "Ivanov",
			Patronymic:  "Ivanovich",
			Age:         20,
			Gender:      "male",
			Nationality: "RU",
		},

		{
			Id:          2,
			Name:        "Galina",
			Surname:     "Petrova",
			Patronymic:  "Petrovna",
			Age:         40,
			Gender:      "female",
			Nationality: "US",
		},
	}

	testCases := []struct {
		name                 string
		url                  string
		userFilter           *domain.UserFilter
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{

		{
			name: "no filters",
			url:  "/api/v1/users",
			userFilter: &domain.UserFilter{
				Limit: 10,
			},
			mockBehavior: func(s *mock_service.MockUserService, ctx context.Context, userFilter *domain.UserFilter) {
				s.EXPECT().Get(ctx, userFilter).Return(users, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"id":1,"name":"Ivan","surname":"Ivanov","patronymic":"Ivanovich","age":20,"gender":"male","nationality":"RU"},{"id":2,"name":"Galina","surname":"Petrova","patronymic":"Petrovna","age":40,"gender":"female","nationality":"US"}]`,
		},

		{
			name: "filtering by name",
			url:  "/api/v1/users?name=Ivan",
			userFilter: &domain.UserFilter{
				Name:  "Ivan",
				Limit: 10,
			},
			mockBehavior: func(s *mock_service.MockUserService, ctx context.Context, userFilter *domain.UserFilter) {
				s.EXPECT().Get(ctx, userFilter).Return(users[0:1], nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"id":1,"name":"Ivan","surname":"Ivanov","patronymic":"Ivanovich","age":20,"gender":"male","nationality":"RU"}]`,
		},

		{
			name: "filtering by surname",
			url:  "/api/v1/users?surname=Petrova",
			userFilter: &domain.UserFilter{
				Surname: "Petrova",
				Limit:   10,
			},
			mockBehavior: func(s *mock_service.MockUserService, ctx context.Context, userFilter *domain.UserFilter) {
				s.EXPECT().Get(ctx, userFilter).Return(users[1:], nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"id":2,"name":"Galina","surname":"Petrova","patronymic":"Petrovna","age":40,"gender":"female","nationality":"US"}]`,
		},

		{
			name: "filtering by patronymic",
			url:  "/api/v1/users?patronymic=Ivanovich",
			userFilter: &domain.UserFilter{
				Patronymic: "Ivanovich",
				Limit:      10,
			},
			mockBehavior: func(s *mock_service.MockUserService, ctx context.Context, userFilter *domain.UserFilter) {
				s.EXPECT().Get(ctx, userFilter).Return(users[0:1], nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"id":1,"name":"Ivan","surname":"Ivanov","patronymic":"Ivanovich","age":20,"gender":"male","nationality":"RU"}]`,
		},

		{
			name: "filtering by age_from",
			url:  "/api/v1/users?age_from=30",
			userFilter: &domain.UserFilter{
				AgeFrom: 30,
				Limit:   10,
			},
			mockBehavior: func(s *mock_service.MockUserService, ctx context.Context, userFilter *domain.UserFilter) {
				s.EXPECT().Get(ctx, userFilter).Return(users[1:], nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"id":2,"name":"Galina","surname":"Petrova","patronymic":"Petrovna","age":40,"gender":"female","nationality":"US"}]`,
		},

		{
			name: "filtering by age_to",
			url:  "/api/v1/users?age_to=30",
			userFilter: &domain.UserFilter{
				AgeTo: 30,
				Limit: 10,
			},
			mockBehavior: func(s *mock_service.MockUserService, ctx context.Context, userFilter *domain.UserFilter) {
				s.EXPECT().Get(ctx, userFilter).Return(users[0:1], nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"id":1,"name":"Ivan","surname":"Ivanov","patronymic":"Ivanovich","age":20,"gender":"male","nationality":"RU"}]`,
		},

		{
			name: "filtering by gender",
			url:  "/api/v1/users?gender=male",
			userFilter: &domain.UserFilter{
				Gender: "male",
				Limit:  10,
			},
			mockBehavior: func(s *mock_service.MockUserService, ctx context.Context, userFilter *domain.UserFilter) {
				s.EXPECT().Get(ctx, userFilter).Return(users[0:1], nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"id":1,"name":"Ivan","surname":"Ivanov","patronymic":"Ivanovich","age":20,"gender":"male","nationality":"RU"}]`,
		},

		{
			name: "filtering by nationality",
			url:  "/api/v1/users?nationality=US",
			userFilter: &domain.UserFilter{
				Nationality: "US",
				Limit:       10,
			},
			mockBehavior: func(s *mock_service.MockUserService, ctx context.Context, userFilter *domain.UserFilter) {
				s.EXPECT().Get(ctx, userFilter).Return(users[1:], nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"id":2,"name":"Galina","surname":"Petrova","patronymic":"Petrovna","age":40,"gender":"female","nationality":"US"}]`,
		},

		{
			name: "filtering with limit",
			url:  "/api/v1/users?limit=1",
			userFilter: &domain.UserFilter{
				Limit: 1,
			},
			mockBehavior: func(s *mock_service.MockUserService, ctx context.Context, userFilter *domain.UserFilter) {
				s.EXPECT().Get(ctx, userFilter).Return(users[0:1], nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"id":1,"name":"Ivan","surname":"Ivanov","patronymic":"Ivanovich","age":20,"gender":"male","nationality":"RU"}]`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			userService := mock_service.NewMockUserService(c)
			tc.mockBehavior(userService, context.Background(), tc.userFilter)

			controller := New(userService)

			// Test router
			r := chi.NewRouter()
			r.Get("/api/v1/users", controller.handleGetUsers(context.Background(), r))

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tc.url, bytes.NewBufferString(""))

			// Perform request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, tc.expectedStatusCode, w.Code)
			assert.Equal(t, tc.expectedResponseBody, w.Body.String())
		})
	}
}

func TestControllerHandleDeleteUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUserService, ctx context.Context, id int)

	testCases := []struct {
		name               string
		id                 string
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{

		{
			name: "valid id",
			id:   "7",
			mockBehavior: func(s *mock_service.MockUserService, ctx context.Context, id int) {
				s.EXPECT().Delete(ctx, id)
			},
			expectedStatusCode: http.StatusOK,
		},

		{
			name:               "invalid id",
			id:                 "id",
			mockBehavior:       func(s *mock_service.MockUserService, ctx context.Context, id int) {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			userService := mock_service.NewMockUserService(c)

			id, _ := strconv.Atoi(tc.id)

			tc.mockBehavior(userService, context.Background(), id)

			controller := New(userService)

			// Test router
			r := chi.NewRouter()
			r.Delete("/api/v1/users/{id}", controller.handleDeleteUser(context.Background()))

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/api/v1/users/"+tc.id, bytes.NewBufferString(""))

			// Perform request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, tc.expectedStatusCode, w.Code)
		})
	}
}

func TestControllerHandleUpdateUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUserService, ctx context.Context, id int, user *domain.User)

	testCases := []struct {
		name               string
		id                 string
		requestBody        string
		user               *domain.User
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{

		{
			name:        "OK",
			id:          "7",
			requestBody: `{"name":"Ivan"}`,
			user: &domain.User{
				Name: "Ivan",
			},
			mockBehavior: func(s *mock_service.MockUserService, ctx context.Context, id int, user *domain.User) {
				s.EXPECT().GetById(ctx, id).Return(&domain.User{}, nil)
				s.EXPECT().Update(ctx, id, user)
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			userService := mock_service.NewMockUserService(c)

			id, _ := strconv.Atoi(tc.id)

			tc.mockBehavior(userService, context.Background(), id, tc.user)

			controller := New(userService)

			// Test router
			r := chi.NewRouter()
			r.Patch("/api/v1/users/{id}", controller.handleUpdateUser(context.Background()))

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPatch, "/api/v1/users/"+tc.id, bytes.NewBufferString(tc.requestBody))

			// Perform request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, tc.expectedStatusCode, w.Code)
		})
	}
}

func TestControllerHandleCreateUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUserService, ctx context.Context, user *domain.User)

	testCases := []struct {
		name               string
		requstBody         string
		user               *domain.User
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{

		{
			name:       "OK",
			requstBody: `{"name":"Ivan","surname":"Ivanov","patronymic":"Ivanovich"}`,
			user: &domain.User{
				Name:       "Ivan",
				Surname:    "Ivanov",
				Patronymic: "Ivanovich",
			},
			mockBehavior: func(s *mock_service.MockUserService, ctx context.Context, user *domain.User) {
				s.EXPECT().Create(ctx, user).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},

		{
			name:       "blank patronymic",
			requstBody: `{"name":"Ivan","surname":"Ivanov"}`,
			user: &domain.User{
				Name:    "Ivan",
				Surname: "Ivanov",
			},
			mockBehavior: func(s *mock_service.MockUserService, ctx context.Context, user *domain.User) {
				s.EXPECT().Create(ctx, user).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},

		{
			name:       "blank name",
			requstBody: `{"surname":"Ivanov"}`,
			user: &domain.User{
				Surname: "Ivanov",
			},
			mockBehavior:       func(s *mock_service.MockUserService, ctx context.Context, user *domain.User) {},
			expectedStatusCode: http.StatusBadRequest,
		},

		{
			name:       "blank surname",
			requstBody: `{"name":"Ivan"}`,
			user: &domain.User{
				Name: "Ivan",
			},
			mockBehavior:       func(s *mock_service.MockUserService, ctx context.Context, user *domain.User) {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			userService := mock_service.NewMockUserService(c)
			tc.mockBehavior(userService, context.Background(), tc.user)

			controller := New(userService)

			// Test router
			r := chi.NewRouter()
			r.Post("/api/v1/users", controller.handleCreateUser(context.Background()))

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBufferString(tc.requstBody))

			// Perform request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, tc.expectedStatusCode, w.Code)
		})
	}
}

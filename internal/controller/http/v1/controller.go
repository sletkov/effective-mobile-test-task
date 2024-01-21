package v1

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sletkov/effective-mobile-test-task/internal/controller/http/v1/converter"
	"github.com/sletkov/effective-mobile-test-task/internal/controller/http/v1/model"
	serviceModel "github.com/sletkov/effective-mobile-test-task/internal/service/model"
)

type UserService interface {
	Get(ctx context.Context, userFilter *serviceModel.UserFilter) ([]serviceModel.User, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, u *serviceModel.User) error
	Create(ctx context.Context, u *serviceModel.User) error
	GetById(ctx context.Context, id int) (*serviceModel.User, error)
}

type UserController struct {
	service UserService
	// logger  slog.Logger
}

func New(service UserService) *UserController {
	return &UserController{
		service: service,
		// logger:  logger,
	}
}

// Initialize routes and return router
func (c *UserController) InitRoutes(ctx context.Context) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {

				r.Get("/", c.handleGetUsers(ctx, r))
				r.Post("/", c.handleCreateUser(ctx))

				r.Route("/{id}", func(r chi.Router) {
					r.Delete("/", c.handleDeleteUser(ctx))
					r.Patch("/", c.handleUpdateUser(ctx))
				})
			})
		})
	})

	return r
}

// Get all users with filters and limit
func (c *UserController) handleGetUsers(ctx context.Context, r chi.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := make([]model.User, 0)

		defaultLimit := 10
		userFilter := &model.UserFilter{}

		values := r.URL.Query()

		for k, v := range values {
			if len(v) > 0 {
				switch k {
				case "name":
					userFilter.Name = v[0]
				case "surname":
					userFilter.Surname = v[0]
				case "patronymic":
					userFilter.Patronymic = v[0]
				case "age_from":
					value, _ := strconv.Atoi(v[0])
					userFilter.AgeFrom = value
				case "age_to":
					value, _ := strconv.Atoi(v[0])
					userFilter.AgeTo = value
				case "gender":
					userFilter.Gender = v[0]
				case "nationality":
					userFilter.Nationality = v[0]
				case "limit":
					value, _ := strconv.Atoi(v[0])
					userFilter.Limit = value
				}
			}
		}

		if userFilter.Limit <= 0 {
			userFilter.Limit = defaultLimit
		}

		serviceUsers, err := c.service.Get(ctx, converter.ToUserFilterFromController(userFilter))

		for _, u := range serviceUsers {
			users = append(users, *converter.ToUserFromService(&u))
		}

		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(users)

		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

// Delete user by id
func (c *UserController) handleDeleteUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := c.service.Delete(ctx, id); err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// Update user
func (c *UserController) handleUpdateUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var updateUser = model.UpdateUser{}

		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := io.ReadAll(r.Body)

		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(data, &updateUser)

		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		u, err := c.service.GetById(ctx, id)

		// Convert from service to controller
		user := converter.ToUserFromService(u)

		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		updateUser.Copy(user)

		if err := c.service.Update(ctx, id, converter.ToUserFromController(user)); err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// Create user
func (c *UserController) handleCreateUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.CreateUser

		data, err := io.ReadAll(r.Body)

		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(data, &user)

		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Validate struct
		err = user.Validate()

		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := c.service.Create(ctx, converter.ToCreateUserFromController(&user)); err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

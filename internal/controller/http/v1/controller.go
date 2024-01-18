package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sletkov/effective-mobile-test-task/internal/model"
	"github.com/sletkov/effective-mobile-test-task/internal/service"
)

type UserController struct {
	service service.UserService
	logger  slog.Logger
}

func New(service service.UserService, logger slog.Logger) *UserController {
	return &UserController{
		service: service,
		logger:  logger,
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
		defaultLimit := 10
		userFilter := &model.UserFilter{}

		values := r.URL.Query()

		for k, v := range values {
			switch k {
			case "name":
				userFilter.Name = v[0]
			case "surname":
				userFilter.Surname = v[0]
			case "patronymic":
				userFilter.Patronymic = v[0]
			case "age_from":
				value, _ := strconv.Atoi(v[0])
				userFilter.AgeFrom = uint8(value)
			case "age_to":
				value, _ := strconv.Atoi(v[0])
				userFilter.AgeTo = uint8(value)
			case "gender":
				userFilter.Gender = v[0]
			case "nationality":
				userFilter.Nationality = v[0]
			case "limit":
				value, _ := strconv.Atoi(v[0])
				userFilter.Limit = uint8(value)
			}
		}

		if userFilter.Limit <= 0 {
			userFilter.Limit = uint8(defaultLimit)
		}

		users, err := c.service.Get(ctx, userFilter)

		if err != nil {
			c.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(users)

		if err != nil {
			c.logger.Error(err.Error())
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
			c.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := c.service.Delete(ctx, uint(id)); err != nil {
			c.logger.Error(err.Error())
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
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		data, err := io.ReadAll(r.Body)

		if err != nil {
			c.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(data, &updateUser)
		fmt.Println(updateUser)

		if err != nil {
			c.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user, err := c.service.GetById(ctx, uint(id))

		if err != nil {
			c.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		updateUser.Copy(user)

		if err := c.service.Update(ctx, uint(id), user); err != nil {
			c.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// Create user
func (c *UserController) handleCreateUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User

		data, err := io.ReadAll(r.Body)

		if err != nil {
			c.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(data, &user)

		if err != nil {
			c.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := c.service.Create(ctx, &user); err != nil {
			c.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

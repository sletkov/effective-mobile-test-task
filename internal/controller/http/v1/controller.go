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
	"github.com/sletkov/effective-mobile-test-task/internal/controller/http/v1/converter"
	"github.com/sletkov/effective-mobile-test-task/internal/controller/http/v1/model"
	"github.com/sletkov/effective-mobile-test-task/internal/domain"
)

//go:generate mockgen -source=controller.go -destination=../../../service/mocks/mock.go

type UserService interface {
	Get(ctx context.Context, userFilter *domain.UserFilter) ([]domain.User, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, u *domain.User) error
	Create(ctx context.Context, u *domain.User) error
	GetById(ctx context.Context, id int) (*domain.User, error)
}

type UserController struct {
	service UserService
}

func New(service UserService) *UserController {
	return &UserController{
		service: service,
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

// @Summary GetUsers
// @Tags users
// @Description get all users with filters and limit
// @ID get-users
// @Produce json
// @Param name query string false "name filter"
// @Param surname query string false "surname filter"
// @Param patronymic query string false "patronymic filter"
// @Param age_from query integer false "min age filter"
// @Param age_to query integer false "max age filter"
// @Param gender query string false "gender filter"
// @Param nationality query string false "nationality filter"
// @Param limit query integer false "limit"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/users [get]
func (c *UserController) handleGetUsers(ctx context.Context, r chi.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := make([]model.User, 0)

		userFilter := &model.UserFilter{}

		err := userFilter.FillFilters(r.URL.Query())

		if err != nil {
			slog.Error(fmt.Sprintf("controller: %s", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Validate struct
		err = userFilter.Validate()

		if err != nil {
			slog.Error(fmt.Sprintf("controller: %s", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		serviceUsers, err := c.service.Get(ctx, converter.ToUserFilterFromController(userFilter))

		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for _, u := range serviceUsers {
			users = append(users, *converter.ToUserFromService(&u))
		}

		data, err := json.Marshal(users)

		if err != nil {
			slog.Error(fmt.Sprintf("controller: %s", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

// @Summary DeleteUser
// @Tags users
// @Description delete user by id
// @ID delete-user
// @Param id path integer true "user id"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/users/{id} [delete]
func (c *UserController) handleDeleteUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			slog.Error(fmt.Sprintf("controller: %s", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
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

// @Summary UpdateUser
// @Tags users
// @Description update user
// @ID update-user
// @Accept json
// @Param id path integer true "user id"
// @Param name body string false "user name"
// @Param surname body string false "user surname"
// @Param patronymic body string false "user patronymic"
// @Param age body integer false "user age"
// @Param gender body string false "user gender"
// @Param nationality body string false "user nationality"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/users/{id} [patch]
func (c *UserController) handleUpdateUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var updateUser = model.UpdateUser{}

		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			slog.Error(fmt.Sprintf("controller: %s", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		data, err := io.ReadAll(r.Body)

		if err != nil {
			slog.Error(fmt.Sprintf("controller: %s", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(data, &updateUser)

		slog.Debug(fmt.Sprintf("controller: got structure: %v", updateUser))

		if err != nil {
			slog.Error(fmt.Sprintf("controller: %s", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validate struct
		err = updateUser.Validate()

		if err != nil {
			slog.Error(fmt.Sprintf("controller: %s", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		u, err := c.service.GetById(ctx, id)

		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Convert from service to controller
		user := converter.ToUserFromService(u)

		if err != nil {
			slog.Error(fmt.Sprintf("controller: %s", err.Error()))
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

// @Summary CreateUser
// @Tags users
// @Description create user
// @ID create-user
// @Accept json
// @Param name body string true "user name"
// @Param surname body string true "user surname"
// @Param patronymic body string false "user patronymic"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/users [post]
func (c *UserController) handleCreateUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.CreateUser

		data, err := io.ReadAll(r.Body)

		if err != nil {
			slog.Error(fmt.Sprintf("controller: %s", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(data, &user)

		slog.Debug(fmt.Sprintf("controller: got structure: %v", user))

		if err != nil {
			slog.Error(fmt.Sprintf("controller: %s", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validate struct
		err = user.Validate()

		if err != nil {
			slog.Error(fmt.Sprintf("controller: %s", err.Error()))
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

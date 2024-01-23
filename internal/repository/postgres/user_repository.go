package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/sletkov/effective-mobile-test-task/internal/repository/postgres/model"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Get all users with filters and limit
func (r *UserRepository) Get(ctx context.Context, userFilter *model.UserFilter) ([]model.User, error) {
	slog.Info("postgres: getting users")

	var users []model.User

	user := model.User{}

	filters := userFilter.GetFilterRequest()

	query, _, err := sq.
		Select("id, name, surname, patronymic, age, gender, nationality").
		From("users").
		Where(filters).
		Limit(uint64(userFilter.Limit)).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("postgres: getting users: %w", err)
	}

	slog.Debug(fmt.Sprintf("postgres: making db query: %s", query))

	rows, err := r.db.QueryContext(
		ctx,
		query,
	)

	if err != nil {
		return nil, fmt.Errorf("postgres: getting users: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Surname,
			&user.Patronymic,
			&user.Age,
			&user.Gender,
			&user.Nationality,
		)

		if err != nil {
			return nil, fmt.Errorf("postgres: getting users: %w", err)
		}

		users = append(users, user)
	}

	slog.Info("postgres: user were got successfully")

	return users, nil
}

// Delete user by id
func (r *UserRepository) Delete(ctx context.Context, id int) error {
	slog.Info(fmt.Sprintf("postgres: deleting user %d", id))

	query, _, err := sq.
		Delete("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return fmt.Errorf("postgres: deleting user %d: %w", id, err)
	}

	slog.Debug(fmt.Sprintf("postgres: making db query: %s", query))

	if _, err := r.db.ExecContext(
		ctx,
		query,
		id,
	); err != nil {
		return fmt.Errorf("postgres: deleting user %d: %w", id, err)
	}

	slog.Info(fmt.Sprintf("postgres: user %d was deleted successfully", id))

	return nil
}

// Update user
func (r *UserRepository) Update(ctx context.Context, id int, u *model.User) error {
	slog.Info(fmt.Sprintf("postgres: updating user %d", id))

	query, _, err := sq.
		Update("users").
		Set("name", u.Name).
		Set("surname", u.Surname).
		Set("patronymic", u.Patronymic).
		Set("age", u.Age).
		Set("gender", u.Gender).
		Set("nationality", u.Nationality).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return fmt.Errorf("postgres: updating user %d: %w", id, err)
	}

	slog.Debug(fmt.Sprintf("postgres: making db query: %s", query))

	if _, err := r.db.ExecContext(
		ctx,
		query,
		u.Name,
		u.Surname,
		u.Patronymic,
		u.Age,
		u.Gender,
		u.Nationality,
		id,
	); err != nil {
		return fmt.Errorf("postgres: updating user %d: %w", id, err)
	}

	slog.Info(fmt.Sprintf("postgres: user %d was updated successfully", id))

	return nil
}

// Create new user
func (r *UserRepository) Create(ctx context.Context, u *model.User) (int, error) {
	slog.Info("postgres: creating user")

	var id int

	query, _, err := sq.
		Insert("users").
		Columns("name", "surname", "patronymic", "age", "gender", "nationality").
		PlaceholderFormat(sq.Dollar).
		Values(u.Name, u.Surname, u.Patronymic, u.Age, u.Gender, u.Nationality).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("postgres: creating user %d: %w", id, err)
	}

	slog.Debug(fmt.Sprintf("postgres: making db query: %s", query))

	if err := r.db.QueryRowContext(
		ctx,
		query,
		u.Name,
		u.Surname,
		u.Patronymic,
		u.Age,
		u.Gender,
		u.Nationality,
	).Scan(&id); err != nil {
		return 0, fmt.Errorf("postgres: creating user %d: %w", id, err)
	}

	slog.Info(fmt.Sprintf("postgres: user %d was created successfully", id))

	return id, nil
}

// Get user by id
func (r *UserRepository) GetUserById(ctx context.Context, id int) (*model.User, error) {
	slog.Info(fmt.Sprintf("postgres: getting user %d", id))

	user := &model.User{}

	query, _, err := sq.
		Select("id", "name", "surname", "patronymic", "age", "gender", "nationality").
		From("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("postgres: getting user %d: %w", id, err)
	}

	slog.Debug(fmt.Sprintf("postgres: making db query: %s", query))

	if err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Surname,
		&user.Patronymic,
		&user.Age,
		&user.Gender,
		&user.Nationality,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrUserNotFound
		}
		return nil, fmt.Errorf("postgres: getting user %d: %w", id, err)
	}

	slog.Debug(fmt.Sprintf("postgres: user %d was got successfully", id))

	return user, nil
}

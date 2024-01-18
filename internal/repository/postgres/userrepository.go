package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sletkov/effective-mobile-test-task/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Get(ctx context.Context, userFilter *model.UserFilter) ([]model.User, error) {
	var users []model.User
	user := model.User{}
	filters := userFilter.GetFilterRequest()

	query := fmt.Sprintf("SELECT * FROM users WHERE %s LIMIT %d", filters, userFilter.Limit)

	rows, err := r.db.QueryContext(
		ctx,
		query,
	)

	if err != nil {
		return nil, err
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
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	if _, err := r.db.ExecContext(
		ctx,
		"DELETE FROM users WHERE id = $1",
		id,
	); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(ctx context.Context, id uint, u *model.User) error {
	if _, err := r.db.ExecContext(
		ctx,
		"UPDATE users SET name = $1, surname = $2, patronymic = $3, age = $4, gender = $5, nationality = $6 WHERE id = $7",
		u.Name,
		u.Surname,
		u.Patronymic,
		u.Age,
		u.Gender,
		u.Nationality,
		id,
	); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Create(ctx context.Context, u *model.User) (uint, error) {
	var id uint
	if err := r.db.QueryRowContext(
		ctx,
		"INSERT INTO users (name, surname, patronymic, age, gender, nationality) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		u.Name,
		u.Surname,
		u.Patronymic,
		u.Age,
		u.Gender,
		u.Nationality,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id uint) (*model.User, error) {
	user := &model.User{}

	if err := r.db.QueryRowContext(
		ctx,
		"SELECT * FROM users WHERE id = $1",
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
		fmt.Println("scan error")
		return nil, err
	}

	return user, nil
}

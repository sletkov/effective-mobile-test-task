package controller

import "context"

type UserController interface {
	handleGetUsers(ctx context.Context)
	handleDeleteUser(ctx context.Context)
	handleUpdateUser(ctx context.Context)
	handleCreateUser(ctx context.Context)
}

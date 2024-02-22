package model

import (
	"context"
	customError "main/errorLogger"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateUserResponse struct {
	User *User `json:"user"`
}

type GetUserRequest struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	User *User `json:"user"`
}

type UpdateUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UpdateUserResponse struct {
	User *User `json:"user"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

type UpsertUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UpsertUserResponse struct {
	User *User `json:"user"`
}

type UserServer interface {
	Create(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, *customError.Error)
	Get(ctx context.Context, in *GetUserRequest) (*GetUserResponse, *customError.Error)
	Update(ctx context.Context, in *UpdateUserRequest) (*UpdateUserResponse, *customError.Error)
	Delete(ctx context.Context, in *DeleteUserRequest) *customError.Error
	Upsert(ctx context.Context, in *UpsertUserRequest) (*UpsertUserResponse, *customError.Error)
}

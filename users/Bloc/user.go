package Bloc

import (
	"context"
	"errors"
	customError "main/errorLogger"
	"main/users/model"
	"main/users/query"
	"net/http"
	"time"
)

type Bloc struct {
	query *query.QueryExecuter
}

func NewBloc(query *query.QueryExecuter) *Bloc {
	return &Bloc{query: query}
}

func (b *Bloc) Get(ctx context.Context, in *model.GetUserRequest) (*model.GetUserResponse, *customError.Error) {
	if in.ID == "" {
		return nil, customError.NewError(http.StatusBadRequest, "ID is required", "id is not provided in request", "users", "Get User")
	}

	res, err := b.query.GetUser(ctx, in.ID)
	if err != nil {
		return nil, customError.NewError(http.StatusInternalServerError, "Internal Server error", err.Error(), "users", "Get User")
	}

	user := &model.User{}
	if res.Next() {
		err = res.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		if err != nil {
			return nil, customError.NewError(
				http.StatusInternalServerError, "Internal Server error", err.Error(), "users", "Get User")
		}
	}

	if user.ID == "" {
		return nil, customError.NewError(http.StatusNotFound, "Not Found", "user not found", "users", "Get User")
	}

	return &model.GetUserResponse{
		User: user,
	}, nil
}

func (b *Bloc) Create(ctx context.Context, in *model.CreateUserRequest) (*model.CreateUserResponse, *customError.Error) {
	id := "user_" + time.Now().String()
	user := &model.User{
		ID:       id,
		Username: in.Username,
		Password: in.Password,
		Email:    in.Email,
	}

	if err := validateUser(user); err != nil {
		return nil, customError.NewError(http.StatusBadRequest, "Invalid request", "some of the details are not present in request", "users", "Create User")
	}

	res, err := b.query.CreateUser(ctx, user.ID, user.Username, user.Password, user.Email)
	if err != nil {
		return nil, customError.NewError(http.StatusInternalServerError, "Internal Server Error", err.Error(), "users", "Create User")
	}

	dbUser := &model.User{}
	if res.Next() {
		err = res.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password, &dbUser.Email)
		if err != nil {
			return nil, customError.NewError(
				http.StatusInternalServerError, "Internal Server error", err.Error(), "users", "Create User")
		}
	}
	return &model.CreateUserResponse{
		User: dbUser,
	}, nil
}

func (b *Bloc) Update(ctx context.Context, in *model.UpdateUserRequest) (*model.UpdateUserResponse, *customError.Error) {
	if in.ID == "" {
		return nil, customError.NewError(http.StatusBadRequest, "ID is required", "id is not provided in request", "users", "Update User")
	}

	_, er := b.Get(ctx, &model.GetUserRequest{ID: in.ID})
	if er != nil {
		return nil, er
	}

	user := &model.User{
		ID:       in.ID,
		Username: in.Username,
		Password: in.Password,
		Email:    in.Email,
	}

	if err := validateUser(user); err != nil {
		return nil, customError.NewError(http.StatusBadRequest, "Invalid request", "some of the details are not present in request", "users", "Update User")
	}

	res, err := b.query.UpdateUser(ctx, user.ID, user.Username, user.Password, user.Email)
	if err != nil {
		return nil, customError.NewError(http.StatusInternalServerError, "Internal Server error", err.Error(), "users", "Update User")
	}

	dbUser := &model.User{}
	if res.Next() {
		err = res.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password, &dbUser.Email)
		if err != nil {
			return nil, customError.NewError(
				http.StatusInternalServerError, "Internal Server error", err.Error(), "users", "Update User")
		}
	}

	return &model.UpdateUserResponse{
		User: dbUser,
	}, nil
}

func (b *Bloc) Delete(ctx context.Context, in *model.DeleteUserRequest) *customError.Error {
	if in.ID == "" {
		return customError.NewError(http.StatusBadRequest, "ID is required", "id is not provided in request", "users", "Delete User")
	}

	err := b.query.DeleteUser(ctx, in.ID)
	if err != nil {
		return customError.NewError(http.StatusInternalServerError, "Internal Server error", err.Error(), "users", "Delete User")
	}
	return nil
}

func (b *Bloc) Upsert(ctx context.Context, in *model.UpsertUserRequest) (*model.UpsertUserResponse, *customError.Error) {

	if in.ID == "" {
		return nil, customError.NewError(http.StatusBadRequest, "ID is required", "id is not provided in request", "users", "Upsert User")
	}

	user := &model.User{
		ID:       in.ID,
		Username: in.Username,
		Password: in.Password,
		Email:    in.Email,
	}

	if err := validateUser(user); err != nil {
		return nil, customError.NewError(http.StatusBadRequest, "Invalid request", "some of the details are not present in request", "users", "Upsert User")
	}

	if user.ID == "" {
		id := "user_" + time.Now().String()
		user.ID = id
		res, err := b.Create(ctx, &model.CreateUserRequest{
			Username: user.Username,
			Password: user.Password,
			Email:    user.Email,
		})
		if err != nil {
			err.Information.Event = "Upsert User"
			return nil, err
		}

		return &model.UpsertUserResponse{
			User: res.User,
		}, nil
	}

	res, err := b.Update(ctx, &model.UpdateUserRequest{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	})
	if err != nil {
		err.Information.Event = "Upsert User"
		return nil, err
	}

	return &model.UpsertUserResponse{
		User: res.User,
	}, nil
}

func validateUser(user *model.User) error {
	if user.Username == "" {
		return errors.New("username is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

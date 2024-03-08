package model

import (
	"context"
	customError "main/errorLogger"
)

type Post struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Parent  string `json:"parent"`
}

type CreatePostRequest struct {
	Content string `json:"content"`
	Parent  string `json:"parent"`
}

type CreatePostResponse struct {
	Post *Post `json:"post"`
}

type GetPostRequest struct {
	Id string `json:"id"`
}

type GetPostResponse struct {
	Post *Post `json:"post"`
}

type UpdatePostRequest struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Parent  string `json:"parent"`
}

type UpdatePostResponse struct {
	Post *Post `json:"post"`
}

type DeletePostRequest struct {
	Id string `json:"id"`
}

type DeleteAllPostOfUserRequest struct {
	UserId string `json:"userId"`
}

type PostServer interface {
	Create(ctx context.Context, in *CreatePostRequest) (*CreatePostResponse, *customError.Error)
	Get(ctx context.Context, in *GetPostRequest) (*GetPostResponse, *customError.Error)
	Update(ctx context.Context, in *UpdatePostRequest) (*UpdatePostResponse, *customError.Error)
	Delete(ctx context.Context, in *DeletePostRequest) *customError.Error
	DeleteAllPostOfUser(ctx context.Context, in *DeleteAllPostOfUserRequest) *customError.Error
}

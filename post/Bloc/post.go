package Bloc

import (
	"context"
	customError "main/errorLogger"
	"main/post/model"
	"main/post/query"
	users "main/users/Bloc"
	usersModel "main/users/model"
	"time"
)

type Bloc struct {
	query *query.QueryExecuter
	user  *users.Bloc
}

func NewBloc(query *query.QueryExecuter) *Bloc {
	return &Bloc{query: query}
}

func (b *Bloc) Create(ctx context.Context, in *model.CreatePostRequest) (*model.CreatePostResponse, *customError.Error) {
	if in.Content == "" || in.Parent == "" {
		return nil, customError.NewError(400, "Content is required", "content is not provided in request", "post", "Create Post")
	}

	_, er := b.user.Get(ctx, &usersModel.GetUserRequest{ID: in.Parent})
	if er != nil {
		return nil, customError.NewError(404, "Not Found", "user not found", "post", "Delete All Post Of User")
	}

	id := "post_" + time.Now().String()
	res, err := b.query.CreatePost(ctx, id, in.Content, in.Parent)
	if err != nil {
		return nil, customError.NewError(500, "Internal Server error", err.Error(), "post", "Create Post")
	}

	post := &model.Post{}
	if res.Next() {
		err = res.Scan(&post.Id, &post.Content, &post.Parent)
		if err != nil {
			return nil, customError.NewError(500, "Internal Server error", err.Error(), "post", "Create Post")
		}
	}

	if post.Id == "" {
		return nil, customError.NewError(500, "Internal Server error", "post not created", "post", "Create Post")
	}

	return &model.CreatePostResponse{
		Post: post,
	}, nil

}

func (b *Bloc) Get(ctx context.Context, in *model.GetPostRequest) (*model.GetPostResponse, *customError.Error) {
	if in.Id == "" {
		return nil, customError.NewError(400, "ID is required", "id is not provided in request", "post", "Get Post")
	}

	res, err := b.query.GetPost(ctx, in.Id)
	if err != nil {
		return nil, customError.NewError(500, "Internal Server error", err.Error(), "post", "Get Post")
	}

	post := &model.Post{}
	if res.Next() {
		err = res.Scan(&post.Id, &post.Content, &post.Parent)
		if err != nil {
			return nil, customError.NewError(500, "Internal Server error", err.Error(), "post", "Get Post")
		}
	}

	if post.Id == "" {
		return nil, customError.NewError(404, "Not Found", "post not found", "post", "Get Post")
	}

	return &model.GetPostResponse{
		Post: post,
	}, nil

}

func (b *Bloc) Update(ctx context.Context, in *model.UpdatePostRequest) (*model.UpdatePostResponse, *customError.Error) {
	if in.Id == "" {
		return nil, customError.NewError(400, "ID is required", "id is not provided in request", "post", "Update Post")
	}

	if in.Content == "" || in.Parent == "" {
		return nil, customError.NewError(400, "Content is required", "content is not provided in request", "post", "Update Post")
	}

	_, er := b.user.Get(ctx, &usersModel.GetUserRequest{ID: in.Parent})
	if er != nil {
		return nil, customError.NewError(404, "Not Found", "user not found", "post", "Delete All Post Of User")
	}

	res, err := b.query.UpdatePost(ctx, in.Id, in.Content, in.Parent)
	if err != nil {
		return nil, customError.NewError(500, "Internal Server error", err.Error(), "post", "Update Post")
	}

	post := &model.Post{}
	if res.Next() {
		err = res.Scan(&post.Id, &post.Content, &post.Parent)
		if err != nil {
			return nil, customError.NewError(500, "Internal Server error", err.Error(), "post", "Update Post")
		}
	}

	if post.Id == "" {
		return nil, customError.NewError(500, "Internal Server error", "post not updated", "post", "Update Post")
	}

	return &model.UpdatePostResponse{
		Post: post,
	}, nil
}

func (b *Bloc) Delete(ctx context.Context, in *model.DeletePostRequest) *customError.Error {
	if in.Id == "" {
		return customError.NewError(400, "ID is required", "id is not provided in request", "post", "Delete Post")
	}

	_, err := b.query.DeletePost(ctx, in.Id)
	if err != nil {
		return customError.NewError(500, "Internal Server error", err.Error(), "post", "Delete Post")
	}
	return nil
}

func (b *Bloc) DeleteAllPostOfUser(ctx context.Context, in *model.DeleteAllPostOfUserRequest) *customError.Error {
	if in.UserId == "" {
		return customError.NewError(400, "UserId is required", "userId is not provided in request", "post", "Delete All Post Of User")
	}

	_, err := b.user.Get(ctx, &usersModel.GetUserRequest{ID: in.UserId})
	if err != nil {
		return customError.NewError(404, "Not Found", "user not found", "post", "Delete All Post Of User")
	}

	_, er := b.query.DeleteAllPostOfUser(ctx, in.UserId)
	if er != nil {
		return customError.NewError(500, "Internal Server error", err.Error(), "post", "Delete All Post Of User")
	}
	return nil
}

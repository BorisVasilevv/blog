package repository

import (
	"context"
	"crud/internal/core/model"
)

type AuthRepository interface {
	GetUser(ctx context.Context, login, hashPassword string) (string, string, error)
	Register(ctx context.Context, login, hashPassword string, role string) (string, string, error)
	LogIn(ctx context.Context, login, hashPassword string) (string, error)
}

type PostRepository interface {
	CreatePost(ctx context.Context, post model.Post) (int, error)
	GetPost(ctx context.Context, postId int) (model.Post, error)
	ChangePost(ctx context.Context, postId int, post model.Post) (model.Post, error)
	GetPosts(ctx context.Context) ([]model.Post, error)
	LikePost(ctx context.Context, postId int) (model.Post, error)
}

type ComRepository interface {
	CreateComment(ctx context.Context, comment model.Comment) (int, error)
	GetComment(ctx context.Context, comId int) (model.Comment, error)
}

package service

import (
	"context"
	"crud/internal/core/model"
)

type AuthService interface {
	Register(ctx context.Context, login, password string, role string) (string, error)
	LogIn(ctx context.Context, login, password string) (string, error)
	GenerateToken(ctx context.Context, login, password string) (string, error)
}

type PostService interface {
	CreatePost(ctx context.Context, post model.Post) (int, error)
	GetPost(ctx context.Context, postId int) (model.Post, error)
	LikePost(ctx context.Context, postId int) (model.Post, error)
}

type ComService interface {
	CreateComment(ctx context.Context, comment model.Comment) (int, error)
	GetComment(ctx context.Context, comId int) (model.Comment, error)
}

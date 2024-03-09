package service

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"errors"
	"log/slog"
)

type _comService struct {
	repo repository.ComRepository
}

func NewComService(repo repository.ComRepository) service.ComService {
	return _comService{repo: repo}
}

func (comService _comService) CreateComment(ctx context.Context, post model.Comment) (int, error) {
	id, err := comService.repo.CreateComment(ctx, post)

	if err != nil {
		slog.Error(err.Error())
		return 0, errors.New("ошибка создания комментария")
	}

	return id, nil
}

func (comService _comService) GetComment(ctx context.Context, comId int) (model.Comment, error) {
	return comService.repo.GetComment(ctx, comId)
}

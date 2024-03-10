package postgres

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/model"
	"crud/internal/lib/db"
	"crud/internal/repository/dbModel"
	"fmt"
)

type _comRepository struct {
	db *db.Db
}

func NewComRepo(db *db.Db) repository.ComRepository {
	return _comRepository{db}
}

func (comRepository _comRepository) CreateComment(ctx context.Context, comment model.Comment) (int, error) {
	comDb := dbModel.Comment(comment)
	var id int

	err := comRepository.db.PgConn.QueryRow(ctx,
		`INSERT INTO public.comment(id_post, body, author) values ($1,$2,$3) RETURNING comment_id`,
		comDb.Id_post,
		comDb.Body,
		comDb.Author).Scan(&id)

	return id, err
}

func (comRepository _comRepository) GetComment(ctx context.Context, comId int) (model.Comment, error) {
	var comment dbModel.Comment

	err := comRepository.db.PgConn.QueryRow(ctx,
		`SELECT p.id_post, p.body, p.author FROM public.comment p WHERE p.id=$1`,
		comId).Scan(&comment.Id_post, &comment.Body, &comment.Author)

	if err != nil {
		return model.Comment{}, fmt.Errorf("ошибка получения комментария: %s", err.Error())
	}

	return model.Comment(comment), nil

}

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
		`SELECT c.id_post, c.body, c.author FROM public.comment c WHERE c.comment_id=$1`,
		comId).Scan(&comment.Id_post, &comment.Body, &comment.Author)

	if err != nil {
		return model.Comment{}, fmt.Errorf("ошибка получения комментария: %s", err.Error())
	}

	return model.Comment(comment), nil

}

func (comRepository _comRepository) GetCommentsByPost(ctx context.Context, postId int) ([]model.Comment, error) {
	var comments []dbModel.Comment

	rows, err := comRepository.db.PgConn.Query(ctx,
		`SELECT c.body, c.author, c.id_post FROM public.comment c WHERE c.id_post=$1`, postId)
	if err != nil {
		return _commentModelDBtoProgram(comments), fmt.Errorf("ошибка получения комментариев: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var comment dbModel.Comment

		if err := rows.Scan(&comment.Body, &comment.Author, &comment.Id_post); err != nil {
			return nil, fmt.Errorf("ошибка сканирования комментария: %s", err.Error())
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка итерации постов: %s", err.Error())
	}

	return _commentModelDBtoProgram(comments), nil

}

func _commentModelDBtoProgram(commentsDB []dbModel.Comment) []model.Comment {
	var result []model.Comment

	for _, commentDB := range commentsDB {
		result = append(result, model.Comment(commentDB))
	}
	return result
}

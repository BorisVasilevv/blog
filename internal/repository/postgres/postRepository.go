package postgres

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/model"
	"crud/internal/lib/db"
	"crud/internal/repository/dbModel"
	"fmt"
)

type _postRepository struct {
	db *db.Db
}

func NewPostRepo(db *db.Db) repository.PostRepository {
	return _postRepository{db}
}

func (postRepository _postRepository) CreatePost(ctx context.Context, post model.Post) (int, error) {
	postDb := dbModel.Post(post)
	var id int

	err := postRepository.db.PgConn.QueryRow(ctx,
		`INSERT INTO public.post(title, body, image, author, likes) values ($1,$2,$3,$4,$5) RETURNING post_id`,
		postDb.Title,
		postDb.Body,
		postDb.ImageURL,
		postDb.Author,
		postDb.Likes).Scan(&id)

	return id, err
}

func (postRepository _postRepository) GetPost(ctx context.Context, postId int) (model.Post, error) {
	var post dbModel.Post

	err := postRepository.db.PgConn.QueryRow(ctx,
		`SELECT p.title, p.body, p.image, p.author, p.likes FROM public.post p WHERE p.post_id=$1`,
		postId).Scan(&post.Title, &post.Body, &post.ImageURL, &post.Author, &post.Likes)

	if err != nil {
		return model.Post{}, fmt.Errorf("ошибка получения поста: %s", err.Error())
	}

	return model.Post(post), nil

}

func (postRepository _postRepository) GetPosts(ctx context.Context) ([]model.Post, error) {
	var posts []dbModel.Post

	rows, err := postRepository.db.PgConn.Query(ctx, `SELECT p.title, p.body, p.image, p.author, p.likes FROM public.post p`)
	if err != nil {
		return _postModelDBtoProgram(posts), fmt.Errorf("ошибка получения постов: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var post dbModel.Post

		if err := rows.Scan(&post.Title, &post.Body, &post.ImageURL, &post.Author, &post.Likes); err != nil {
			return nil, fmt.Errorf("ошибка сканирования поста: %s", err.Error())
		}
		posts = append(posts, post)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка итерации постов: %s", err.Error())
	}

	return _postModelDBtoProgram(posts), nil
}

func (postRepository _postRepository) LikePost(ctx context.Context, postId int) (model.Post, error) {

	post, err := postRepository.GetPost(ctx, postId)

	likes := post.Likes + 1
	print(post.Likes)
	_, err = postRepository.db.PgConn.Exec(ctx,
		`UPDATE public.post SET likes=$2 WHERE post_id=$1`,
		postId, likes)
	if err != nil {
		return model.Post{}, fmt.Errorf("ошибка изменения поста: %s", err.Error())
	}
	newPost, err := postRepository.GetPost(ctx, postId)

	if err != nil {
		return model.Post{}, fmt.Errorf("ошибка изменения поста: %s", err.Error())
	}

	return newPost, nil
}

func (postRepository _postRepository) ChangePost(ctx context.Context, postId int, newPost model.Post) (model.Post, error) {
	var post dbModel.Post

	err := postRepository.db.PgConn.QueryRow(ctx,
		`UPDATE public.post p SET p.title=$1, p.body=$2, p.image=$3, p.author=$4 WHERE p.post_id=$5`,
		newPost.Title, newPost.Body, newPost.ImageURL, newPost.Author, postId).Scan(&post.Title, &post.Body, &post.ImageURL, &post.Author, &post.Likes)

	if err != nil {
		return model.Post{}, fmt.Errorf("ошибка изменения поста: %s", err.Error())
	}

	return model.Post(post), nil

}

func _postModelDBtoProgram(postsDB []dbModel.Post) []model.Post {
	var result []model.Post

	for _, postDB := range postsDB {
		result = append(result, model.Post(postDB))
	}
	return result
}

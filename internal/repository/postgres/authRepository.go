package postgres

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/lib/db"
	"crud/internal/repository/dbModel"
	"fmt"
)

type _authRepo struct {
	*db.Db
}

func NewRepo(db *db.Db) repository.AuthRepository {
	return _authRepo{db}
}

func (repo _authRepo) GetUser(ctx context.Context, login, hashPassword string) (string, string, error) {
	var user dbModel.User

	row := repo.PgConn.QueryRow(ctx, `SELECT * FROM public.user WHERE login=$1 AND pas=$2`, login, hashPassword)

	if err := row.Scan(&user); err != nil {
		return "", "", fmt.Errorf("не смогли получить юзера: %x", err)
	}
	var role = user.Role
	return login, role, nil

}

func (repo _authRepo) Register(ctx context.Context, login, hashPassword string, role string) (string, string, error) {
	_, err := repo.PgConn.Exec(
		ctx,
		`INSERT INTO public.user(login, pass, role) values ($1, $2, $3)`,
		login, hashPassword, role,
	)

	if err != nil {
		return "", "", fmt.Errorf("не смогли создать: %x", err)
	}

	return login, role, nil
}

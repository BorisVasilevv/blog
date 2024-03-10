package dbModel

type User struct {
	Login    string `db:"login"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

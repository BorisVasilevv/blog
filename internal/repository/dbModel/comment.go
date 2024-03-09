package dbModel

type Comment struct {
	Id_post int    `db:"id_post"`
	Body    string `db:"body"`
	Author  string `db:"author"`
}

package comment

import (
	"database/sql"
	postgresDB "youtube/pkg/db"
)

type postgres struct {
	Conn *sql.DB
}

var Repo Repository

func init() {
	Repo = &postgres{
		Conn: postgresDB.GetPGConnection(),
	}
}

type Repository interface {
	Insert(comment Comment) (int64, string, error)
	Delete(commentID int64) (int64, string, error)
}

package playlist

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
	Insert(playList PlayList) (int64, string, error)
}

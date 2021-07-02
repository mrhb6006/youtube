package likeDislike

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
	ChechExist(userID int64, videoID int64) (bool, string, error)
	Insert(like Like) (string, error)
	UpdateAction(like Like) (string, error)
}

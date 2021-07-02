package likeDislikeVideo

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
	Insert(like LikeVideo) (string, error)
	UpdateAction(like LikeVideo) (string, error)
}

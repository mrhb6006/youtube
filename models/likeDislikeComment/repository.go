package likeDislikeComment

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
	ChechExist(userID int64, commentID int64) (bool, string, error)
	Insert(like Like) (string, error)
	UpdateAction(like Like) (string, error)
	GetLikesCount(likeOrDislike int64, commentID int64) (int64, string, error)
}

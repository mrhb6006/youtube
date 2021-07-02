package playListVideo

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
	Insert(videoPlayList VideoPlayList) (string, error)
	Exist(playListID, videoID int64) (bool, string, error)
}

package seenVideo

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
	Insert(seen Seen) (errStr string, err error)
	CountVideoSeens(VideoID int64) (seenCount int64, errStr string, err error)
}

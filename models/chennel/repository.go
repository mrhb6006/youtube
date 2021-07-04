package chennel

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
	GetByID(ID int64) (Channel, bool, string, error)
	GetByName(name string) (channel Channel, isExist bool, errStr string, err error)
	CreateChannel(channel Channel) (id int64, errStr string, err error)
	DeleteChannel(channelID int64) (errStr string, err error)
}

package channelUser

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
	Insert(channelUser ChannelUser) (string, error)
	ExistenceCheck(channelUser ChannelUser) (bool, string, error)
	Delete(channelUser ChannelUser) (string, error)
}

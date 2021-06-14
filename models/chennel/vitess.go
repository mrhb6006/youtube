package chennel

import (
	"database/sql"
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) GetByID(ID int64) (channel Channel, isExist bool, errStr string, err error) {
	channel = Channel{}
	err = pg.Conn.QueryRow("SELECT id,name,creation_date,description,avatar,creator_id FROM channel WHERE id=$1", ID).Scan(
		&channel.ID,
		&channel.Name,
		&channel.CreationDate,
		&channel.Description,
		&channel.Avatar,
		&channel.CreatorID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return channel, false, "", nil
		}
		zap.L().Error("get_channel_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return channel, false, "01", err
	}
	return channel, true, "", nil
}

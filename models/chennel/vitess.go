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

func (pg *postgres) GetByName(name string) (channel Channel, isExist bool, errStr string, err error) {
	channel = Channel{}
	err = pg.Conn.QueryRow("SELECT id,name,creation_date,description,avatar,creator_id FROM channel WHERE name=$1", name).Scan(
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

func (pg *postgres) CreateChannel(channel Channel) (id int64, errStr string, err error) {
	err = pg.Conn.QueryRow("insert into channel (name, creation_date, description, avatar, creator_id) values ($1,$2,$3,$4,$5) returning id", channel.Name, time.Now().Format("2006-01-02"), channel.Description, channel.Avatar, channel.CreatorID).Scan(&id)
	if err != nil {
		zap.L().Error("create_channel_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return 0, "01", err
	}
	return id, "", nil
}

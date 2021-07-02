package channelUser

import (
	"database/sql"
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) Insert(channelUser ChannelUser) (string, error) {
	_, err := pg.Conn.Exec("insert into user_channel (user_id, channel_id) VALUES ($1,$2)", channelUser.UserID, channelUser.ChannelID)
	if err != nil {
		zap.L().Error("insert_channelUser_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return "01", err
	}
	return "", nil
}

func (pg postgres) ExistenceCheck(channelUser ChannelUser) (exist bool, errStr string, err error) {
	result := ChannelUser{}
	err = pg.Conn.QueryRow("select user_id,channel_id from user_channel where user_id=$1 and channel_id=$2", channelUser.UserID, channelUser.ChannelID).Scan(
		&result.UserID,
		&result.ChannelID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "01", nil
		}
		zap.L().Error("ExistenceCheck_channelUser_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return false, "02", err
	}
	return true, "03", nil
}

func (pg *postgres) Delete(channelUser ChannelUser) (string, error) {
	_, err := pg.Conn.Exec("delete from user_channel where user_id=$1 and channel_id=$2", channelUser.UserID, channelUser.ChannelID)
	if err != nil {
		zap.L().Error("delete_channelUser_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return "01", err
	}
	return "", nil
}

func (pg *postgres) GetMembersCount(channelID int64) (int64, string, error) {
	var count int64
	err := pg.Conn.QueryRow("SELECT count(user_id) from user_channel where channel_id=$1", channelID).Scan(&count)
	if err != nil {
		zap.L().Error("update_like_video_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return -1, "01", err
	}
	return count, "", nil
}

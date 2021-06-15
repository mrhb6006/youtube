package channelVideo

import (
	"database/sql"
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) Insert(channelVideo ChannelVideo) (string, error) {
	_, err := pg.Conn.Exec("INSERT INTO channel_video (channel_id,video_id) VALUES ($1,$2)", channelVideo.ChannelID, channelVideo.VideoID)
	if err != nil {
		zap.L().Error("insert_videoChannel_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return "01", err
	}
	return "", nil
}

func (pg postgres) ExistenceCheck(channelVideo ChannelVideo) (exist bool, errStr string, err error) {
	result := ChannelVideo{}
	err = pg.Conn.QueryRow("select channel_id,video_id from channel_video where channel_id=$1 and video_id=$2", channelVideo.ChannelID, channelVideo.VideoID).Scan(
		&result.ChannelID,
		&result.VideoID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "", nil
		}
		zap.L().Error("ExistenceCheck_videoChannel_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return false, "01", err
	}
	return true, "02", nil
}

package channelVideo

import (
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) Insert(channelVideo ChannelVideo) (string, error) {
	_, err := pg.Conn.Exec("INSERT INTO video_channel (channel_id,video_id) VALUES ($1,$2)", channelVideo.ChannelID, channelVideo.VideoID)
	if err != nil {
		zap.L().Error("insert_videoChannel_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return "01", err
	}
	return "", nil
}

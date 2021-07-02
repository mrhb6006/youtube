package seenVideo

import (
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) Insert(seen Seen) (errStr string, err error) {
	_, err = pg.Conn.Exec("INSERT INTO seen (user_id, video_id) VALUES ($1,$2);", seen.UserID, seen.VideoID)
	if err != nil {
		zap.L().Error("create_channel_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return "01", err
	}
	return "", nil
}

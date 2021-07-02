package seenVideo

import (
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) Insert(seen Seen) (errStr string, err error) {
	_, err = pg.Conn.Exec("INSERT INTO seen (user_id, video_id) VALUES ($1,$2);", seen.UserID, seen.VideoID)
	if err != nil {
		zap.L().Error("insert_video_seen_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return "01", err
	}
	return "", nil
}

func (pg *postgres) CountVideoSeens(VideoID int64) (seenCount int64, errStr string, err error) {
	var count int64
	err = pg.Conn.QueryRow("select count(user_id) from seen where video_id=$1 group by video_id;", VideoID).Scan(&count)
	if err != nil {
		zap.L().Error("vount_video_seens_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return -1, "01", err
	}
	return count, "", nil
}

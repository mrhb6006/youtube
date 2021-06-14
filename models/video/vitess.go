package video

import (
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) Insert(video Video) (insertedID int64, errStr string, err error) {
	err = pg.Conn.QueryRow("INSERT INTO video (title,description,upload_date,duration,thumbnail,storage_id) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id", video.Title, video.Description, time.Now().Format("2006-01-02"), video.Duration, video.Thumbnail, video.StorageID).Scan(&insertedID)
	if err != nil {
		zap.L().Error("insert_video_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return 0, "01", err
	}
	return insertedID, "", nil
}

package video

import (
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) Insert(video Video) (insertedID int64, errStr string, err error) {
	result, err := pg.Conn.Exec("INSERT INTO video (title,description,upload_date,duration,thumbnail,storage_id) VALUE (?,?,?,?,?,?)", video.Title, video.Description, video.UploadDate, video.Duration, video.Thumbnail, video.StorageID)
	if err != nil {
		zap.L().Error("insert_video_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return 0, "01", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, "02", err
	}
	return id, "", nil
}

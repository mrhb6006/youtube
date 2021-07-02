package playListVideo

import (
	"database/sql"
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) Insert(list VideoPlayList) (errStr string, err error) {
	_, err = pg.Conn.Exec("INSERT INTO playlist_video (playlist_id,video_id) VALUES ($1,$2)", list.PlayListID, list.VideoID)
	if err != nil {
		zap.L().Error("insert_playlist_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return "01", err
	}
	return "", nil
}

func (pg *postgres) Exist(playListID, videoID int64) (bool, string, error) {
	temp := 0
	err := pg.Conn.QueryRow("SELECT playlist_id FROM playlist_video WHERE playlist_id=$1 AND video_id=$2", playListID, videoID).Scan(&temp)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "", nil
		}
		zap.L().Error("exist_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return false, "01", err
	}
	return true, "", nil
}

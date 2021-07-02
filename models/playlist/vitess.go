package playlist

import (
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) Insert(playList PlayList) (insertedID int64, errStr string, err error) {
	err = pg.Conn.QueryRow("INSERT INTO playlist (name, is_default, is_public, creator_id) VALUES ($1,$2,$3,$4); RETURNING id", playList.Name, playList.IsDefault, playList.IsPublic, playList.CreatorID).Scan(&insertedID)
	if err != nil {
		zap.L().Error("insert_playlist_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return 0, "01", err
	}
	return insertedID, "", nil
}

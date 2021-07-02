package playlist

import (
	"database/sql"
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

func (pg *postgres) GetByName(name string, creatorID int64) (playList PlayList, exist bool, errStr string, err error) {
	playList = PlayList{}
	err = pg.Conn.QueryRow("SELECT * FROM playlist WHERE name=$1 AND creatorID=$2", name, creatorID).Scan(
		&playList.ID,
		&playList.Name,
		&playList.IsDefault,
		&playList.IsDefault,
		&playList.CreatorID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return playList, false, "", nil
		}
		zap.L().Error("get_playList_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return playList, false, "01", err
	}
	return playList, true, "", nil
}

func (pg *postgres) GetByID(playListId int64) (playList PlayList, exist bool, errStr string, err error) {
	playList = PlayList{}
	err = pg.Conn.QueryRow("SELECT * FROM playlist WHERE id=$1", playListId).Scan(
		&playList.ID,
		&playList.Name,
		&playList.IsDefault,
		&playList.IsDefault,
		&playList.CreatorID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return playList, false, "", nil
		}
		zap.L().Error("get_playList_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return playList, false, "01", err
	}
	return playList, true, "", nil
}

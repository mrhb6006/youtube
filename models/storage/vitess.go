package storage

import (
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) Insert(storage Storage) (insertedID int64, errStr string, err error) {
	err = pg.Conn.QueryRow("INSERT INTO storage (path) VALUES ($1) RETURNING id", storage.Path).Scan(&insertedID)
	if err != nil {
		zap.L().Error("insert_storage_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return 0, "01", err
	}
	return insertedID, "", nil
}

package storage

import (
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) Insert(storage Storage) (int64, string, error) {
	result, err := pg.Conn.Exec("INSERT INTO storage (path) VALUE ($1)", storage.Path)
	if err != nil {
		zap.L().Error("insert_storage_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return 0, "01", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, "02", err
	}
	return id, "", nil
}

package globalSearch

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
)

func (pg *postgres) Search(word, table string) ([]Search, string, error) {
	matched := make([]Search, 0)
	likeCondition := "'%" + word + "%'"
	playlistCondition := ""
	colName := "name"
	if table == "playlist" {
		playlistCondition = "AND is_public=true"
	} else if table == "video" {
		colName = "title"
	}
	query := fmt.Sprintf("SELECT %s,id from %s where lower(%s) LIKE %s %s", colName, table, colName, likeCondition, playlistCondition)
	rows, err := pg.Conn.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return matched, "", nil
		}
		zap.L().Error("search err ", zap.Error(err))
		return nil, "01", err
	}
	defer rows.Close()
	for rows.Next() {
		var tempName string
		var id int64
		err := rows.Scan(
			&tempName,
			&id,
		)
		if err != nil {
			zap.L().Error("scan db err", zap.Error(err))
			return nil, "02", err
		}
		matched = append(matched, Search{
			ID:   id,
			Name: tempName,
		})
	}
	return matched, "", nil
}

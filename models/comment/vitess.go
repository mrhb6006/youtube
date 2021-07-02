package comment

import (
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) Insert(comment Comment) (insertedID int64, errStr string, err error) {
	err = pg.Conn.QueryRow("INSERT INTO commnet (text,reply_id,user_id,is_deleted,video_id) VALUES ($1,$2,$3,$4,$5) RETURNING id", comment.Text, comment.ReplyID, comment.UserID, comment.IsDeleted, comment.VideoID).Scan(&insertedID)
	if err != nil {
		zap.L().Error("insert_comment_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return 0, "01", err
	}
	return insertedID, "", nil
}

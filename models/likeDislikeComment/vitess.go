package likeDislikeComment

import (
	"database/sql"
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) ChechExist(userID int64, commentID int64) (bool, string, error) {
	like := Like{}
	err := pg.Conn.QueryRow("select * from like_comment where user_id=$1 and comment_id=$2;", userID, commentID).Scan(&like.CommentID, &like.Action, &like.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "", nil
		}
		zap.L().Error("check_exist_like_comment_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return false, "01", err
	}
	return true, "", nil
}

func (pg *postgres) Insert(like Like) (string, error) {
	_, err := pg.Conn.Exec("insert into like_comment (comment_id, action, user_id) VALUES ($1,$2,$3);", like.CommentID, like.Action, like.UserID)
	if err != nil {
		zap.L().Error("insert_like_comment_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return "01", err
	}
	return "", nil
}

func (pg *postgres) UpdateAction(like Like) (string, error) {
	_, err := pg.Conn.Exec("update like_comment set action=$1 where comment_id=$2 and user_id=$3;", like.Action, like.CommentID, like.UserID)
	if err != nil {
		zap.L().Error("update_like_comment_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return "01", err
	}
	return "", nil
}

func (pg *postgres) GetLikesCount(likeOrDislike int64, commentID int64) (int64, string, error) {
	var count int64
	err := pg.Conn.QueryRow("select count(user_id) from like_comment where comment_id=$1 and action=$2;", commentID, likeOrDislike).Scan(&count)
	if err != nil {
		zap.L().Error("update_like_video_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return -1, "01", err
	}
	return count, "", nil
}

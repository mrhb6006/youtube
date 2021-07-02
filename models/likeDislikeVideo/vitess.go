package likeDislikeVideo

import (
	"database/sql"
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) ChechExist(userID int64, videoID int64) (bool, string, error) {
	like := LikeVideo{}
	err := pg.Conn.QueryRow("select * from like_video where user_id=$1 and video_id=$2;", userID, videoID).Scan(&like.VideoID, &like.Action, &like.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "", nil
		}
		zap.L().Error("check_exist_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return false, "01", err
	}
	return true, "", nil
}

func (pg *postgres) Insert(like LikeVideo) (string, error) {
	_, err := pg.Conn.Exec("insert into like_video (video_id, action, user_id) VALUES ($1,$2,$3);", like.VideoID, like.Action, like.UserID)
	if err != nil {
		zap.L().Error("insert_like_video_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return "01", err
	}
	return "", nil
}

func (pg *postgres) UpdateAction(like LikeVideo) (string, error) {
	_, err := pg.Conn.Exec("update like_video set action=$1 where video_id=$2 and user_id=$3;", like.Action, like.VideoID, like.UserID)
	if err != nil {
		zap.L().Error("update_like_video_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return "01", err
	}
	return "", nil
}

func (pg *postgres) GetVideoLikesCount(likeOrDislike int64, videoID int64) (int64, string, error) {
	var count int64
	err := pg.Conn.QueryRow("select count(user_id) from like_video where video_id=$1 and action=$2;", videoID, likeOrDislike).Scan(&count)
	if err != nil {
		zap.L().Error("update_like_video_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return -1, "01", err
	}
	return count, "", nil
}

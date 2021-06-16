package user

import (
	"database/sql"
	"go.uber.org/zap"
	"time"
)

func (pg *postgres) Insert(user User) (insertedID int64, errStr string, err error) {
	err = pg.Conn.QueryRow("INSERT INTO user (username,password,email,enroll_date,avatar) VALUES ($1,$2,$3,$4,$5) RETURNING id", user.UserName, user.Password, user.Email, time.Now().Format("2006-01-02"), user.Avatar).Scan(&insertedID)
	if err != nil {
		zap.L().Error("insert_user_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return 0, "01", err
	}
	return insertedID, "", nil
}

func (pg *postgres) GetByUserName(userName string) (user User, found bool, errStr string, err error) {
	user = User{}
	err = pg.Conn.QueryRow("SELECT id,username,password,email,enroll_date,avatar FROM user WHERE username=$1", userName).Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.EnrollDate,
		&user.Avatar,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, false, "", nil
		}
		zap.L().Error("get_user_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return user, false, "01", err
	}
	return user, true, "", nil
}

func (pg *postgres) GetByEmail(email string) (user User, found bool, errStr string, err error) {
	user = User{}
	err = pg.Conn.QueryRow("SELECT id,username,password,email,enroll_date,avatar FROM user WHERE email=$1", email).Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.EnrollDate,
		&user.Avatar,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, false, "", nil
		}
		zap.L().Error("get_user_err", zap.Any("error:", err), zap.Any("time :", time.Now().UnixNano()))
		return user, false, "01", err
	}
	return user, true, "", nil
}

package user

import (
	"net/mail"
	"strings"
)

type User struct {
	ID         int64
	UserName   string
	Password   string
	Email      string
	Avatar     string
	EnrollDate string
}

const MinNameLength = 5
const MinPassLength = 8

func PrepareAndValidationUser(user User) (formattedUser User, isValid bool, completeErrStr string) {
	user.UserName = strings.TrimSpace(user.UserName)
	user.Password = strings.TrimSpace(user.Password)
	user.Email = strings.TrimSpace(user.Email)

	if len(user.UserName) < MinNameLength {
		return User{}, false, "100100"
	} else if len(user.Password) < MinPassLength {
		return User{}, false, "100101"
	}
	if user.Email != "" {
		if len(user.Email) < 5 {
			return user, false, "100102"
		}
		_, err := mail.ParseAddress(user.Email)
		if err != nil {
			return user, false, "100102"
		}
	}
	return user, true, ""
}

package user

import (
	"crypto/md5"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
	"sync"
	"time"
	"youtube/apiType/commonType"
	"youtube/apiType/userType"
	"youtube/models/user"
	"youtube/pkg/auth"
	"youtube/pkg/response"
	"youtube/pkg/storageHandler"
)

func Register(ctx *fiber.Ctx) error {
	baseErrCode := "100"
	request := userType.RegisterRequest{}
	registerResponse := userType.RegisterResponse{}
	res := commonType.Response{}

	errStr, code, err := request.DecodeValidate(ctx)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "01", errStr, code)
	}

	var wg sync.WaitGroup
	var mutex sync.RWMutex

	wg.Add(2)
	var errStrChan = make(chan string)
	var sectionNoChan = make(chan string)
	var done = make(chan bool)
	found := make([]bool, 2)

	go func() {
		defer wg.Done()
		_, exist, errStr, err := user.Repo.GetByUserName(request.Username)
		if err != nil {
			sectionNoChan <- "02"
			errStrChan <- errStr
		}
		if exist {
			mutex.Lock()
			found[0] = true
			mutex.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		_, exist, errStr, err := user.Repo.GetByEmail(request.Email)
		if err != nil {
			sectionNoChan <- "03"
			errStrChan <- errStr
		}
		if exist {
			mutex.Lock()
			found[1] = true
			mutex.Unlock()
		}
	}()

	go func() {
		defer func() { done <- true }()
		wg.Wait()
	}()

	select {
	case <-done:
		break
	case sectionNo := <-sectionNoChan:
		errStr = <-errStrChan
		return response.ErrorResponse(ctx, res, baseErrCode, sectionNo, errStr, 500)
	}

	if found[0] || found[1] {
		if found[0] {
			errStr = "01"
		} else {
			errStr = "02"
		}
		return response.ErrorResponse(ctx, res, baseErrCode, "04", errStr, 200)
	}

	byteArray := md5.Sum([]byte(request.Password))
	hashPassStr := fmt.Sprintf("%x", byteArray)

	avatarPath := ""
	if request.Avatar != "" {
		avatarPath, err = storageHandler.SaveImage(request.Avatar, "userAvatar"+strconv.FormatInt(time.Now().UnixNano(), 10))
		if err != nil {
			zap.L().Error("save image error", zap.Any("err:", err))
			return response.ErrorResponse(ctx, res, baseErrCode, "05", "01", 500)
		}
	}

	id, errStr, err := user.Repo.Insert(user.User{
		UserName: request.Username,
		Password: hashPassStr,
		Email:    request.Email,
		Avatar:   avatarPath,
	})
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "06", errStr, 500)
	}
	registerResponse.Token, err = auth.CreateToken(id)
	if err != nil {
		return response.ErrorResponse(ctx, res, baseErrCode, "07", errStr, 500)
	}
	res.Res = registerResponse
	return response.SuccessResponse(ctx, res)
}

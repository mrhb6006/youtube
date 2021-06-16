package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"youtube/pkg/errorHandler"
)

func Middleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := string(ctx.Request().Header.Peek("token"))
		if token == "" {
			zap.L().Debug("empty token")
			return errorHandler.CreateError(ctx, "00001", 403)
		}
		payload, err := Verify(token)
		if err != nil {
			zap.L().Debug("invalid token")
			return errorHandler.CreateError(ctx, "00002", 403)
		}
		ctx.Locals("UserID", payload.UserID)
		return ctx.Next()
	}
}

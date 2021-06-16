package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strings"
	"youtube/pkg/errorHandler"
)

func Middleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		path := string(ctx.Request().URI().RequestURI())
		if strings.Contains(path, "register") || strings.Contains(path, "login") {
			return ctx.Next()
		}
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

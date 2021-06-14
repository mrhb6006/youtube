package main

import (
	"github.com/gofiber/fiber/v2"
	. "youtube/services/video"
)

func setUpRoute(app *fiber.App) {
	youtube := app.Group("/youtube")
	video := youtube.Group("/video")
	video.Post("/upload", Upload)
}

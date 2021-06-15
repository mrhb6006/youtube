package main

import (
	"github.com/gofiber/fiber/v2"
	. "youtube/services/channel"
	. "youtube/services/channelVideo"
	. "youtube/services/video"
)

func setUpRoute(app *fiber.App) {
	youtube := app.Group("/youtube")
	video := youtube.Group("/video")
	video.Post("/upload", Upload)

	channel := youtube.Group("/channel")
	channel.Post("/create", Create)
	channel.Post("/addvideo", AddVideoToChannel)
}

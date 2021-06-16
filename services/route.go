package main

import (
	"github.com/gofiber/fiber/v2"
	"youtube/pkg/auth"
	. "youtube/services/channel"
	. "youtube/services/channelUser"
	. "youtube/services/channelVideo"
	. "youtube/services/video"
)

func setUpRoute(app *fiber.App) {
	app.Use(auth.Middleware())

	youtube := app.Group("/youtube")
	video := youtube.Group("/video")
	video.Post("/upload", Upload)

	channel := youtube.Group("/channel")
	channel.Post("/create", Create)
	channel.Post("/addvideo", AddVideoToChannel)
	channel.Delete("/deletevideo", DeleteVideoFromChannel)
	channel.Post("/join", JoinToChannel)
	channel.Delete("/leave", LeaveChannel)
}

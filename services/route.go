package main

import (
	"github.com/gofiber/fiber/v2"
	"youtube/pkg/auth"
	. "youtube/services/channel"
	. "youtube/services/channelUser"
	. "youtube/services/channelVideo"
	. "youtube/services/comment"
	. "youtube/services/playlist"
	. "youtube/services/search"
	. "youtube/services/likeDislike"
	. "youtube/services/seenVideo"
	. "youtube/services/user"
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

	user := youtube.Group("/user")
	user.Post("/register", Register)
	user.Post("/login", Login)

	video.Post("/seen", SeenVideo)
	video.Get("/seen/count", CountVideoSeen)
	video.Post("/like", LikeDislikeVideo)

	comment := youtube.Group("/comment")
	comment.Post("/write", Write)
	comment.Post("/like", LikeDislikeComment)
	comment.Delete("/delete", DeleteComment)

	youtube.Post("/search", Search)

	playlist := youtube.Group("playlist")
	playlist.Post("/create", CreatePlayList)
}

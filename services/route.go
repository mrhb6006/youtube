package main

import (
	"github.com/gofiber/fiber/v2"
	"youtube/pkg/auth"
	. "youtube/services/channel"
	. "youtube/services/channelUser"
	. "youtube/services/channelVideo"
	. "youtube/services/comment"
	. "youtube/services/likeDislike"
	. "youtube/services/playlist"
	. "youtube/services/search"
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
	channel.Post("/members/count", ChannelMembersCount)

	user := youtube.Group("/user")
	user.Post("/register", Register)
	user.Post("/login", Login)

	video.Post("/seen", SeenVideo)
	video.Post("/seen/count", CountVideoSeen)
	video.Post("/like", LikeDislikeVideo)
	video.Post("/like/count", VideoLikesCount)
	video.Delete("/delete", DeleteVideo)

	comment := youtube.Group("/comment")
	comment.Post("/write", Write)
	comment.Post("/like", LikeDislikeComment)
	comment.Delete("/delete", DeleteComment)
	comment.Post("/like/count", CommentLikesCount)

	youtube.Post("/search", Search)

	playlist := youtube.Group("playlist")
	playlist.Post("/create", CreatePlayList)
	playlist.Post("/addVideo", AddVideoToPlayList)
	playlist.Post("/makePublic", MakePublic)
	playlist.Delete("/delete", DeletePlayList)
}

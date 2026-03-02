package main

import (
	"instaclone/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK", "message": "InstaClone API is running"})
	})

	users := r.Group("/api/users")
	{
		users.GET("", handlers.GetUsers)
		users.GET("/:id", handlers.GetUserByID)
		users.POST("", handlers.CreateUser)
		users.GET("/:id/feed", handlers.GetUserFeed)
		users.GET("/:id/posts", handlers.GetUserPosts)
		users.GET("/:id/likes", handlers.GetMyLikes)
		users.POST("/:id/follow", handlers.FollowUser)
		users.DELETE("/:id/follow", handlers.UnfollowUser)
		users.GET("/:id/followers", handlers.GetFollowers)
		users.GET("/:id/following", handlers.GetFollowing)
	}

	posts := r.Group("/api/posts")
	{
		posts.GET("", handlers.GetPosts)
		posts.GET("/search", handlers.SearchPosts)
		posts.GET("/:id", handlers.GetPostByID)
		posts.POST("", handlers.CreatePost)
		posts.DELETE("/:id", handlers.DeletePost)
		posts.POST("/:id/like", handlers.LikePost)
		posts.DELETE("/:id/like", handlers.UnlikePost)
		posts.GET("/:id/likes", handlers.GetLikes)
		posts.GET("/:id/comments", handlers.GetComments)
		posts.POST("/:id/comments", handlers.CreateComment)
	}

	comments := r.Group("/api/comments")
	{
		comments.DELETE("/:id", handlers.DeleteComment)
	}

	r.Run(":8080")
}

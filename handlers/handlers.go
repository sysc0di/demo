package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"instaclone/data"
	"instaclone/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": data.Users})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, user := range data.Users {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{"user": user})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

type CreateUserInput struct {
	Username    string `json:"username" binding:"required"`
	DisplayName string `json:"displayName" binding:"required"`
	Bio         string `json:"bio"`
	AvatarURL   string `json:"avatarUrl"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		ID:          "user-" + uuid.New().String()[:8],
		Username:    input.Username,
		DisplayName: input.DisplayName,
		Bio:         input.Bio,
		AvatarURL:   input.AvatarURL,
		CreatedAt:   data.Users[0].CreatedAt,
	}

	data.Users = append(data.Users, user)
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func GetPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 10
	}

	start := (page - 1) * limit
	end := start + limit

	if start >= len(data.Posts) {
		c.JSON(http.StatusOK, gin.H{
			"posts":      []models.Post{},
			"page":       page,
			"limit":      limit,
			"total":      len(data.Posts),
			"totalPages": (len(data.Posts) + limit - 1) / limit,
		})
		return
	}

	if end > len(data.Posts) {
		end = len(data.Posts)
	}

	c.JSON(http.StatusOK, gin.H{
		"posts":      data.Posts[start:end],
		"page":       page,
		"limit":      limit,
		"total":      len(data.Posts),
		"totalPages": (len(data.Posts) + limit - 1) / limit,
	})
}

func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	for _, post := range data.Posts {
		if post.ID == id {
			c.JSON(http.StatusOK, gin.H{"post": post})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}

type CreatePostInput struct {
	UserID   string `json:"userId" binding:"required"`
	ImageURL string `json:"imageUrl" binding:"required"`
	Caption  string `json:"caption"`
}

func CreatePost(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userExists := false
	for _, user := range data.Users {
		if user.ID == input.UserID {
			userExists = true
			break
		}
	}

	if !userExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	post := models.Post{
		ID:           "post-" + uuid.New().String()[:8],
		UserID:       input.UserID,
		ImageURL:     input.ImageURL,
		Caption:      input.Caption,
		LikeCount:    0,
		CommentCount: 0,
		CreatedAt:    data.Posts[0].CreatedAt,
	}

	data.Posts = append([]models.Post{post}, data.Posts...)
	c.JSON(http.StatusCreated, gin.H{"post": post})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	for i, post := range data.Posts {
		if post.ID == id {
			data.Posts = append(data.Posts[:i], data.Posts[i+1:]...)
			for i := len(data.Comments) - 1; i >= 0; i-- {
				if data.Comments[i].PostID == id {
					data.Comments = append(data.Comments[:i], data.Comments[i+1:]...)
				}
			}
			for i := len(data.Likes) - 1; i >= 0; i-- {
				if data.Likes[i].PostID == id {
					data.Likes = append(data.Likes[:i], data.Likes[i+1:]...)
				}
			}
			c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}

type LikeInput struct {
	UserID string `json:"userId" binding:"required"`
}

func LikePost(c *gin.Context) {
	postID := c.Param("id")
	var input LikeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postIndex := -1
	for i, post := range data.Posts {
		if post.ID == postID {
			postIndex = i
			break
		}
	}

	if postIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	for _, like := range data.Likes {
		if like.UserID == input.UserID && like.PostID == postID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Already liked"})
			return
		}
	}

	data.Likes = append(data.Likes, models.Like{UserID: input.UserID, PostID: postID})
	data.Posts[postIndex].LikeCount++

	c.JSON(http.StatusOK, gin.H{"message": "Post liked"})
}

func UnlikePost(c *gin.Context) {
	postID := c.Param("id")
	userID := c.Query("userId")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}

	postIndex := -1
	for i, post := range data.Posts {
		if post.ID == postID {
			postIndex = i
			break
		}
	}

	if postIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	for i, like := range data.Likes {
		if like.UserID == userID && like.PostID == postID {
			data.Likes = append(data.Likes[:i], data.Likes[i+1:]...)
			data.Posts[postIndex].LikeCount--
			c.JSON(http.StatusOK, gin.H{"message": "Post unliked"})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Not liked"})
}

func GetComments(c *gin.Context) {
	postID := c.Param("id")
	var comments []models.Comment

	for _, comment := range data.Comments {
		if comment.PostID == postID {
			comments = append(comments, comment)
		}
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

type CreateCommentInput struct {
	UserID string `json:"userId" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

func CreateComment(c *gin.Context) {
	postID := c.Param("id")
	var input CreateCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postIndex := -1
	for i, post := range data.Posts {
		if post.ID == postID {
			postIndex = i
			break
		}
	}

	if postIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	comment := models.Comment{
		ID:        "comment-" + uuid.New().String()[:8],
		PostID:    postID,
		UserID:    input.UserID,
		Text:      input.Text,
		CreatedAt: data.Comments[0].CreatedAt,
	}

	data.Comments = append(data.Comments, comment)
	data.Posts[postIndex].CommentCount++

	c.JSON(http.StatusCreated, gin.H{"comment": comment})
}

func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	for i, comment := range data.Comments {
		if comment.ID == id {
			for j, post := range data.Posts {
				if post.ID == comment.PostID {
					data.Posts[j].CommentCount--
					break
				}
			}
			data.Comments = append(data.Comments[:i], data.Comments[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
}

type FollowInput struct {
	FollowerID string `json:"followerId" binding:"required"`
}

func FollowUser(c *gin.Context) {
	userID := c.Param("id")
	var input FollowInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userID == input.FollowerID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot follow yourself"})
		return
	}

	for _, follow := range data.Follows {
		if follow.FollowerID == input.FollowerID && follow.FollowingID == userID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Already following"})
			return
		}
	}

	data.Follows = append(data.Follows, models.Follow{
		FollowerID:  input.FollowerID,
		FollowingID: userID,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Now following user"})
}

func UnfollowUser(c *gin.Context) {
	userID := c.Param("id")
	followerID := c.Query("followerId")

	if followerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "followerId is required"})
		return
	}

	for i, follow := range data.Follows {
		if follow.FollowerID == followerID && follow.FollowingID == userID {
			data.Follows = append(data.Follows[:i], data.Follows[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Unfollowed user"})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Not following"})
}

func GetFollowers(c *gin.Context) {
	userID := c.Param("id")
	var followerIDs []string

	for _, follow := range data.Follows {
		if follow.FollowingID == userID {
			followerIDs = append(followerIDs, follow.FollowerID)
		}
	}

	var followers []models.User
	for _, id := range followerIDs {
		for _, user := range data.Users {
			if user.ID == id {
				followers = append(followers, user)
				break
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"followers": followers, "count": len(followers)})
}

func GetFollowing(c *gin.Context) {
	userID := c.Param("id")
	var followingIDs []string

	for _, follow := range data.Follows {
		if follow.FollowerID == userID {
			followingIDs = append(followingIDs, follow.FollowingID)
		}
	}

	var following []models.User
	for _, id := range followingIDs {
		for _, user := range data.Users {
			if user.ID == id {
				following = append(following, user)
				break
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"following": following, "count": len(following)})
}

func GetUserFeed(c *gin.Context) {
	userID := c.Param("id")

	var followingIDs []string
	for _, follow := range data.Follows {
		if follow.FollowerID == userID {
			followingIDs = append(followingIDs, follow.FollowingID)
		}
	}

	var posts []models.Post
	for _, post := range data.Posts {
		if post.UserID == userID || contains(followingIDs, post.UserID) {
			posts = append(posts, post)
		}
	}

	if len(posts) == 0 {
		posts = data.Posts[:5]
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts, "count": len(posts)})
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func GetLikes(c *gin.Context) {
	postID := c.Param("id")
	var userIDs []string

	for _, like := range data.Likes {
		if like.PostID == postID {
			userIDs = append(userIDs, like.UserID)
		}
	}

	var users []models.User
	for _, id := range userIDs {
		for _, user := range data.Users {
			if user.ID == id {
				users = append(users, user)
				break
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"users": users, "count": len(users)})
}

func GetMyLikes(c *gin.Context) {
	userID := c.Param("id")
	var postIDs []string

	for _, like := range data.Likes {
		if like.UserID == userID {
			postIDs = append(postIDs, like.PostID)
		}
	}

	var posts []models.Post
	for _, id := range postIDs {
		for _, post := range data.Posts {
			if post.ID == id {
				posts = append(posts, post)
				break
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts, "count": len(posts)})
}

func SearchPosts(c *gin.Context) {
	query := strings.ToLower(c.Query("q"))
	var results []models.Post

	for _, post := range data.Posts {
		if strings.Contains(strings.ToLower(post.Caption), query) {
			results = append(results, post)
		}
	}

	c.JSON(http.StatusOK, gin.H{"posts": results, "count": len(results)})
}

func GetUserPosts(c *gin.Context) {
	userID := c.Param("id")
	var posts []models.Post

	for _, post := range data.Posts {
		if post.UserID == userID {
			posts = append(posts, post)
		}
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts, "count": len(posts)})
}

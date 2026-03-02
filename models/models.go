package models

import "time"

type User struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"displayName"`
	Bio         string    `json:"bio"`
	AvatarURL   string    `json:"avatarUrl"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Post struct {
	ID           string    `json:"id"`
	UserID       string    `json:"userId"`
	ImageURL     string    `json:"imageUrl"`
	Caption      string    `json:"caption"`
	LikeCount    int       `json:"likeCount"`
	CommentCount int       `json:"commentCount"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Comment struct {
	ID        string    `json:"id"`
	PostID    string    `json:"postId"`
	UserID    string    `json:"userId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

type Like struct {
	UserID  string `json:"userId"`
	PostID  string `json:"postId"`
}

type Follow struct {
	FollowerID string `json:"followerId"`
	FollowingID string `json:"followingId"`
}

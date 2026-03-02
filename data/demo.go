package data

import (
	"instaclone/models"
	"time"
)

var Users = []models.User{
	{
		ID:          "user-1",
		Username:    "johndoe",
		DisplayName: "John Doe",
		Bio:         "Photography enthusiast | Travel lover",
		AvatarURL:   "https://i.pravatar.cc/150?u=user-1",
		CreatedAt:   time.Now().Add(-30 * 24 * time.Hour),
	},
	{
		ID:          "user-2",
		Username:    "janedoe",
		DisplayName: "Jane Doe",
		Bio:         "Foodie | Chef | Recipe creator",
		AvatarURL:   "https://i.pravatar.cc/150?u=user-2",
		CreatedAt:   time.Now().Add(-25 * 24 * time.Hour),
	},
	{
		ID:          "user-3",
		Username:    "alexsmith",
		DisplayName: "Alex Smith",
		Bio:         "Tech geek | Startup founder",
		AvatarURL:   "https://i.pravatar.cc/150?u=user-3",
		CreatedAt:   time.Now().Add(-20 * 24 * time.Hour),
	},
	{
		ID:          "user-4",
		Username:    "sarahlee",
		DisplayName: "Sarah Lee",
		Bio:         "Yoga instructor | Wellness coach",
		AvatarURL:   "https://i.pravatar.cc/150?u=user-4",
		CreatedAt:   time.Now().Add(-15 * 24 * time.Hour),
	},
	{
		ID:          "user-5",
		Username:    "mikebrown",
		DisplayName: "Mike Brown",
		Bio:         "Fitness trainer | Nutritionist",
		AvatarURL:   "https://i.pravatar.cc/150?u=user-5",
		CreatedAt:   time.Now().Add(-10 * 24 * time.Hour),
	},
}

var Posts = []models.Post{
	{
		ID:           "post-1",
		UserID:       "user-1",
		ImageURL:     "https://picsum.photos/seed/post1/800/800",
		Caption:      "Beautiful sunset at the beach! 🌅",
		LikeCount:    42,
		CommentCount: 5,
		CreatedAt:    time.Now().Add(-2 * time.Hour),
	},
	{
		ID:           "post-2",
		UserID:       "user-2",
		ImageURL:     "https://picsum.photos/seed/post2/800/800",
		Caption:      "Made this delicious pasta from scratch! 🍝",
		LikeCount:    89,
		CommentCount: 12,
		CreatedAt:    time.Now().Add(-5 * time.Hour),
	},
	{
		ID:           "post-3",
		UserID:       "user-3",
		ImageURL:     "https://picsum.photos/seed/post3/800/800",
		Caption:      "New office setup - finally productive! 💻",
		LikeCount:    156,
		CommentCount: 23,
		CreatedAt:    time.Now().Add(-8 * time.Hour),
	},
	{
		ID:           "post-4",
		UserID:       "user-4",
		ImageURL:     "https://picsum.photos/seed/post4/800/800",
		Caption:      "Morning yoga session 🧘‍♀️",
		LikeCount:    67,
		CommentCount: 8,
		CreatedAt:    time.Now().Add(-12 * time.Hour),
	},
	{
		ID:           "post-5",
		UserID:       "user-5",
		ImageURL:     "https://picsum.photos/seed/post5/800/800",
		Caption:      "Leg day! No pain no gain 💪",
		LikeCount:    234,
		CommentCount: 31,
		CreatedAt:    time.Now().Add(-1 * time.Hour),
	},
	{
		ID:           "post-6",
		UserID:       "user-1",
		ImageURL:     "https://picsum.photos/seed/post6/800/800",
		Caption:      "Mountain hiking adventure 🏔️",
		LikeCount:    178,
		CommentCount: 15,
		CreatedAt:    time.Now().Add(-24 * time.Hour),
	},
	{
		ID:           "post-7",
		UserID:       "user-2",
		ImageURL:     "https://picsum.photos/seed/post7/800/800",
		Caption:      "Homemade pizza night! 🍕",
		LikeCount:    95,
		CommentCount: 18,
		CreatedAt:    time.Now().Add(-36 * time.Hour),
	},
	{
		ID:           "post-8",
		UserID:       "user-3",
		ImageURL:     "https://picsum.photos/seed/post8/800/800",
		Caption:      "Attended an amazing tech conference!",
		LikeCount:    203,
		CommentCount: 27,
		CreatedAt:    time.Now().Add(-48 * time.Hour),
	},
	{
		ID:           "post-9",
		UserID:       "user-4",
		ImageURL:     "https://picsum.photos/seed/post9/800/800",
		Caption:      "Meditation for inner peace 🧘",
		LikeCount:    145,
		CommentCount: 19,
		CreatedAt:    time.Now().Add(-60 * time.Hour),
	},
	{
		ID:           "post-10",
		UserID:       "user-5",
		ImageURL:     "https://picsum.photos/seed/post10/800/800",
		Caption:      "Post-workout smoothie 🥤",
		LikeCount:    167,
		CommentCount: 22,
		CreatedAt:    time.Now().Add(-72 * time.Hour),
	},
}

var Comments = []models.Comment{
	{ID: "comment-1", PostID: "post-1", UserID: "user-2", Text: "Amazing shot!", CreatedAt: time.Now().Add(-1 * time.Hour)},
	{ID: "comment-2", PostID: "post-1", UserID: "user-3", Text: "Where is this?", CreatedAt: time.Now().Add(-1 * time.Hour)},
	{ID: "comment-3", PostID: "post-2", UserID: "user-1", Text: "Looks delicious!", CreatedAt: time.Now().Add(-4 * time.Hour)},
	{ID: "comment-4", PostID: "post-2", UserID: "user-3", Text: "Recipe please!", CreatedAt: time.Now().Add(-4 * time.Hour)},
	{ID: "comment-5", PostID: "post-2", UserID: "user-4", Text: "Can I taste?", CreatedAt: time.Now().Add(-3 * time.Hour)},
	{ID: "comment-6", PostID: "post-3", UserID: "user-1", Text: "Cool setup!", CreatedAt: time.Now().Add(-7 * time.Hour)},
	{ID: "comment-7", PostID: "post-3", UserID: "user-2", Text: "What monitor is that?", CreatedAt: time.Now().Add(-6 * time.Hour)},
	{ID: "comment-8", PostID: "post-4", UserID: "user-5", Text: "So peaceful!", CreatedAt: time.Now().Add(-10 * time.Hour)},
	{ID: "comment-9", PostID: "post-5", UserID: "user-1", Text: "Beast mode!", CreatedAt: time.Now().Add(-30 * time.Minute)},
	{ID: "comment-10", PostID: "post-5", UserID: "user-4", Text: "Keep it up!", CreatedAt: time.Now().Add(-20 * time.Minute)},
	{ID: "comment-11", PostID: "post-6", UserID: "user-2", Text: "Stunning view!", CreatedAt: time.Now().Add(-20 * time.Hour)},
	{ID: "comment-12", PostID: "post-7", UserID: "user-3", Text: "Homemade looks better!", CreatedAt: time.Now().Add(-30 * time.Hour)},
	{ID: "comment-13", PostID: "post-8", UserID: "user-4", Text: "Wish I was there!", CreatedAt: time.Now().Add(-40 * time.Hour)},
	{ID: "comment-14", PostID: "post-9", UserID: "user-5", Text: "So zen!", CreatedAt: time.Now().Add(-50 * time.Hour)},
	{ID: "comment-15", PostID: "post-10", UserID: "user-2", Text: "What's in it?", CreatedAt: time.Now().Add(-60 * time.Hour)},
}

var Likes = []models.Like{
	{UserID: "user-1", PostID: "post-2"},
	{UserID: "user-1", PostID: "post-3"},
	{UserID: "user-2", PostID: "post-1"},
	{UserID: "user-2", PostID: "post-3"},
	{UserID: "user-3", PostID: "post-1"},
	{UserID: "user-3", PostID: "post-2"},
	{UserID: "user-4", PostID: "post-5"},
	{UserID: "user-5", PostID: "post-4"},
}

var Follows = []models.Follow{
	{FollowerID: "user-1", FollowingID: "user-2"},
	{FollowerID: "user-1", FollowingID: "user-3"},
	{FollowerID: "user-2", FollowingID: "user-1"},
	{FollowerID: "user-2", FollowingID: "user-3"},
	{FollowerID: "user-3", FollowingID: "user-1"},
	{FollowerID: "user-4", FollowingID: "user-5"},
	{FollowerID: "user-5", FollowingID: "user-4"},
}

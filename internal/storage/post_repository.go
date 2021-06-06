package storage

import "basics_web/internal/model"

type PostRepository interface {
	CreatePost(*model.Post) *model.Post
	EditPost(*model.Post) (*model.Post, error)
	GetPost(int64) (*model.Post, error)
	GetPosts() []*model.Post
}

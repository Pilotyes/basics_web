package teststorage

import (
	"basics_web/internal/model"
	"errors"
)

var (
	errNotFound = errors.New("post not found")
)

type postRepository struct {
	posts map[int64]*model.Post
}

func (p *postRepository) CreatePost(post *model.Post) *model.Post {
	id := int64(len(p.posts) + 1)
	post.ID = id
	p.posts[id] = post

	return post
}

func (p *postRepository) EditPost(post *model.Post) (*model.Post, error) {
	if _, ok := p.posts[post.ID]; !ok {
		return nil, errNotFound
	}

	p.posts[post.ID] = post
	return post, nil
}

func (p *postRepository) GetPost(ID int64) (*model.Post, error) {
	if post, ok := p.posts[ID]; ok {
		return post, nil
	}

	return nil, errNotFound
}

func (p *postRepository) GetPosts() []*model.Post {
	posts := make([]*model.Post, len(p.posts))
	for i, post := range p.posts {
		posts[i-1] = post
	}

	return posts
}

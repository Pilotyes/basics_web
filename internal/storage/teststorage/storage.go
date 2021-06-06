package teststorage

import (
	"basics_web/internal/model"
	"basics_web/internal/storage"
)

type TestStorage struct {
	postRepository *postRepository
}

//New ...
func New() *TestStorage {
	return &TestStorage{}
}

//Posts ...
func (t *TestStorage) Posts() storage.PostRepository {
	if t.postRepository != nil {
		return t.postRepository
	}

	t.postRepository = &postRepository{
		posts: map[int64]*model.Post{
			1: {
				ID:     1,
				Title:  "Post 1",
				Text:   "Text 1",
				Author: "Author 1",
			},
			2: {
				ID:     2,
				Title:  "Post 2",
				Text:   "Text 2",
				Author: "Author 2",
			},
			3: {
				ID:     3,
				Title:  "Post 3",
				Text:   "Text 3",
				Author: "Author 1",
			},
		},
	}

	return t.postRepository
}

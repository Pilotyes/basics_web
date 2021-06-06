package teststorage

import (
	"basics_web/internal/model"
	"basics_web/internal/storage"

	"github.com/sirupsen/logrus"
)

type TestStorage struct {
	logger         *logrus.Entry
	postRepository *postRepository
}

//New ...
func New(logger *logrus.Logger) *TestStorage {
	l := logger.WithFields(logrus.Fields{
		"storage": "TestStorage",
	})

	l.Debugln("Created new storage")

	return &TestStorage{
		logger: l,
	}
}

//Posts ...
func (t *TestStorage) Posts() storage.PostRepository {
	if t.postRepository != nil {
		return t.postRepository
	}

	t.logger.Debugln("Created new repository")
	t.postRepository = &postRepository{
		logger: t.logger,
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

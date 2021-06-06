package views

import (
	"basics_web/internal/model"
	"net/http"
)

type Renderer interface {
	RenderPosts(http.ResponseWriter, []*model.Post)
}

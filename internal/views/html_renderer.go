package views

import (
	"basics_web/internal/model"
	"html/template"
	"net/http"
)

const (
	templatesDir   = "./internal/views/templates/"
	templateCreate = "create.tmpl"
	templateEdit   = "edit.tmpl"
	templatePost   = "post.tmpl"
	templatePosts  = "posts.tmpl"
)

var (
	templateFiles = []string{
		templatesDir + templateCreate,
		templatesDir + templateEdit,
		templatesDir + templatePost,
		templatesDir + templatePosts,
	}
)

//HTMLRenderer ...
type HTMLRenderer struct {
	template *template.Template
}

func New() (*HTMLRenderer, error) {
	t := template.Must(template.ParseFiles(templateFiles...))

	return &HTMLRenderer{
		template: t,
	}, nil
}

func (h *HTMLRenderer) RenderCreate(w http.ResponseWriter, post *model.Post) {
	h.template.ExecuteTemplate(w, templateCreate, map[string]interface{}{})
}

func (h *HTMLRenderer) RenderEditPost(w http.ResponseWriter, post *model.Post) {
	h.template.ExecuteTemplate(w, templateEdit, map[string]interface{}{
		"ID":     post.ID,
		"Title":  post.Title,
		"Text":   post.Text,
		"Author": post.Author,
	})
}

func (h *HTMLRenderer) RenderPost(w http.ResponseWriter, post *model.Post) {
	h.template.ExecuteTemplate(w, templatePost, map[string]interface{}{
		"ID":     post.ID,
		"Title":  post.Title,
		"Text":   post.Text,
		"Author": post.Author,
	})
}

func (h *HTMLRenderer) RenderPosts(w http.ResponseWriter, posts []*model.Post) {
	h.template.ExecuteTemplate(w, templatePosts, map[string][]*model.Post{
		"Posts": posts,
	})
}

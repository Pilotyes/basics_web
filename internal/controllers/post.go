package controllers

import (
	"basics_web/internal/model"
	"basics_web/internal/storage"
	"basics_web/internal/views"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//PostController ...
type PostController struct {
	logger  *logrus.Entry
	storage storage.Storage
}

//New ...
func New(logger *logrus.Logger, storage storage.Storage) *PostController {
	l := logger.WithFields(logrus.Fields{
		"controller": "PostController",
	})

	return &PostController{
		logger:  l,
		storage: storage,
	}
}

//CreatePost ...
func (pc *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	pc.logger.Debugln("Create post method", r.Method)
	switch r.Method {
	case http.MethodGet:
		renderer, err := views.New()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		renderer.RenderCreate(w, nil)
	case http.MethodPost:
		if r.Body == nil {
			fmt.Fprintln(w, "Empty body")
			return
		}
		defer r.Body.Close()

		post := &model.Post{
			Title:  r.FormValue("title"),
			Text:   r.FormValue("text"),
			Author: r.FormValue("author"),
		}

		pc.storage.Posts().CreatePost(post)
		fmt.Fprintln(w, post)
	}
}

//EditPost ...
func (pc *PostController) EditPost(w http.ResponseWriter, r *http.Request) {
	pc.logger.Debugln("Edit post method", r.Method)
	switch r.Method {
	case http.MethodGet:
		id := mux.Vars(r)["id"]
		intID, _ := strconv.ParseInt(id, 10, 64)

		post, err := pc.storage.Posts().GetPost(intID)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		renderer, err := views.New()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		renderer.RenderEditPost(w, post)
	case http.MethodPost:
		if r.Body == nil {
			fmt.Fprintln(w, "Empty body")
			return
		}
		defer r.Body.Close()

		postID, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
		if err != nil {
			fmt.Fprintln(w, "Empty body")
			return
		}
		post := &model.Post{
			ID:     postID,
			Title:  r.FormValue("title"),
			Text:   r.FormValue("text"),
			Author: r.FormValue("author"),
		}

		pc.storage.Posts().EditPost(post)
		fmt.Fprintln(w, post)
	}

}

//GetPost ...
func (pc *PostController) GetPost(w http.ResponseWriter, r *http.Request) {
	pc.logger.Debugln("Get post")
	id := mux.Vars(r)["id"]
	intID, _ := strconv.ParseInt(id, 10, 64)

	post, err := pc.storage.Posts().GetPost(intID)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	renderer, err := views.New()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	renderer.RenderPost(w, post)
}

//GetPosts ...
func (pc *PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	pc.logger.Debugln("Get posts")
	posts := pc.storage.Posts().GetPosts()

	renderer, err := views.New()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	renderer.RenderPosts(w, posts)
}

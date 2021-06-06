package apiserver

import (
	"basics_web/internal/controllers"
	"basics_web/internal/storage/teststorage"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	controller *controllers.PostController
	router     *mux.Router
}

func New() *server {
	storage := teststorage.New()
	controller := controllers.New(storage)

	router := mux.NewRouter()
	router.HandleFunc("/post/{id:[0-9]+}", controller.GetPost).Methods(http.MethodGet)
	router.HandleFunc("/posts", controller.GetPosts).Methods(http.MethodGet)
	router.HandleFunc("/create", controller.CreatePost).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/edit/{id:[0-9]+}", controller.EditPost).Methods(http.MethodGet, http.MethodPost)

	return &server{
		controller: controller,
		router:     router,
	}
}

func (s *server) Start() error {
	return http.ListenAndServe(":8080", s.router)
}

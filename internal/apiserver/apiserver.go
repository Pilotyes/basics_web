package apiserver

import (
	"basics_web/internal/config"
	"basics_web/internal/controllers"
	"basics_web/internal/storage/teststorage"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	config     *config.Config
	controller *controllers.PostController
	logger     *logrus.Logger
	router     *mux.Router
}

func New(config *config.Config) (*server, error) {
	logger := logrus.New()
	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return nil, err
	}

	logger.SetLevel(logLevel)

	storage := teststorage.New(logger)
	controller := controllers.New(logger, storage)

	router := mux.NewRouter()

	handlers := []struct {
		path    string
		fn      http.HandlerFunc
		methods []string
	}{
		{
			path: "/post/{id:[0-9]+}",
			fn:   controller.GetPost,
			methods: []string{
				http.MethodGet,
			},
		},

		{
			path: "/posts",
			fn:   controller.GetPosts,
			methods: []string{
				http.MethodGet,
			},
		},
		{
			path: "/create",
			fn:   controller.CreatePost,
			methods: []string{
				http.MethodGet,
				http.MethodPost,
			},
		},
		{
			path: "/edit/{id:[0-9]+}",
			fn:   controller.EditPost,
			methods: []string{
				http.MethodGet,
				http.MethodPost,
			},
		},
	}

	for _, handler := range handlers {
		logger.Debugln("Register new handler:", handler.path, "methods:", handler.methods)
		router.HandleFunc(handler.path, handler.fn).Methods(handler.methods...)
	}

	return &server{
		config:     config,
		controller: controller,
		logger:     logger,
		router:     router,
	}, nil
}

func (s *server) Start() error {
	s.logger.Infoln("Started server at", s.config.BindAddr)
	return http.ListenAndServe(":8080", s.router)
}

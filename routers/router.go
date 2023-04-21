package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thecaptainprice/dictionary-app/backend/handlers"
)

func NewRouter(wordHandler *handlers.WordHandler, userHandler *handlers.UserHandler) http.Handler {
	r := mux.NewRouter()

	// // Middleware
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)

	// Routes for words
	var routes = []Route{
		{
			Path:    "/words",
			Method:  http.MethodGet,
			Name:    "get list of words",
			Handler: DefaultHttpHandler(wordHandler.GetWordsHandlerGeneric),
		},
		{
			Path:    "/words/{id}",
			Method:  http.MethodGet,
			Name:    "GetWordByIDHandler",
			Handler: DefaultHttpHandler(wordHandler.GetWordByIDHandlerGeneric),
		},
		{
			Path:    "/words",
			Method:  http.MethodPost,
			Name:    "GetWordByIDHandler",
			Handler: DefaultHttpHandler(wordHandler.CreateWordHandlerGeneric),
		},
	}
	for _, route := range routes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	// r.HandleFunc("/words", wordHandler.GetWordsHandler).Methods("GET")
	// r.HandleFunc("/words/{id}", wordHandler.GetWordByIDHandler).Methods("GET")
	// r.HandleFunc("/words", wordHandler.CreateWordHandler).Methods("POST")
	// r.HandleFunc("/words/{id}", wordHandler.UpdateWordHandler).Methods("PUT")
	// r.HandleFunc("/words/{id}", wordHandler.DeleteWordHandler).Methods("DELETE")

	// // Routes for users
	// r.HandleFunc("/users/{id}", userHandler.GetUserByIDHandler).Methods("GET")
	// r.HandleFunc("/users", userHandler.CreateUserHandler).Methods("POST")
	// r.HandleFunc("/users/{id}", userHandler.UpdateUserHandler).Methods("PUT")
	// r.HandleFunc("/users/{id}", userHandler.GetUserByIDHandler).Methods("DELETE")

	return r
}

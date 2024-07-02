package router

import (
	"net/http"
	"pawpawchat/internal/grpcclient"
	"pawpawchat/internal/middleware"
	"pawpawchat/internal/router/routes"

	_ "pawpawchat/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Router struct {
	middlewares []middleware.Middleware
	router      *mux.Router
	user        routes.User
}

func New(client *grpcclient.Client) *Router {
	return &Router{
		router:      mux.NewRouter(),
		user:        routes.NewUserRoutes(client),
		middlewares: make([]middleware.Middleware, 0),
	}
}

func (r *Router) Configure() {
	r.router.PathPrefix("/swagger/").HandlerFunc(httpSwagger.WrapHandler)
	r.Use(middleware.CORS)
	r.Use(middleware.Logging)

	r.router.HandleFunc("/signup", r.user.SignUp).Methods("POST")
	r.router.HandleFunc("/signin", r.user.SignIn).Methods("POST")
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler = r.router

	for idx := range r.middlewares {
		handler = r.middlewares[idx](handler)
	}

	handler.ServeHTTP(w, req)
}

func (r *Router) Use(mw middleware.Middleware) {
	r.middlewares = append(r.middlewares, mw)
}

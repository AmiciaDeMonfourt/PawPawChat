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
	client      *grpcclient.Client
}

func New(client *grpcclient.Client) *Router {
	return &Router{
		router:      mux.NewRouter(),
		user:        routes.NewUserRoutes(client),
		middlewares: make([]middleware.Middleware, 0),
		client:      client,
	}
}

func (r *Router) Configure() {
	r.router.PathPrefix("/swagger/").HandlerFunc(httpSwagger.WrapHandler)
	r.configureMiddlewares()
	r.configureUserRoutes()
}

func (r *Router) configureUserRoutes() {
	r.router.HandleFunc("/signup", r.user.SignUp).Methods("POST")
	r.router.HandleFunc("/signin", r.user.SignIn).Methods("POST")
	r.router.HandleFunc("/{username}", middleware.Auth(r.client, r.user.Page)).Methods("GET")
}

func (r *Router) configureMiddlewares() {
	r.Use(middleware.CORS)
	r.Use(middleware.Logging)
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

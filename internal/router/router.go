package router

import (
	"net/http"
	"pawpawchat/internal/grpc"
	"pawpawchat/internal/middleware"
	"pawpawchat/internal/producer"
	"pawpawchat/internal/router/routes"
	"pawpawchat/internal/router/routes/user"

	_ "pawpawchat/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Router struct {
	router      *mux.Router
	gRPCClient  *grpc.Client
	user        routes.User
	middlewares []middleware.Middleware
	producer    *producer.Producer
}

func New(gRPCClient *grpc.Client, producer *producer.Producer) *Router {
	return &Router{
		router:      mux.NewRouter(),
		user:        user.NewUserRoutes(gRPCClient, producer),
		middlewares: make([]middleware.Middleware, 0),
		gRPCClient:  gRPCClient,
		producer:    producer,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler = r.router

	// run of middlewares
	for idx := range r.middlewares {
		handler = r.middlewares[idx](handler)
	}

	handler.ServeHTTP(w, req)
}

func (r *Router) Configure() {
	r.configureSwagger()
	r.configureMiddlewares()
	r.configureUserRoutes()

	r.producer.Logs()
}

func (r *Router) configureUserRoutes() {
	r.router.HandleFunc("/signup", r.user.SignUp).Methods("POST")
	r.router.HandleFunc("/signin", r.user.SignIn).Methods("POST")
	r.router.HandleFunc("/api/user", middleware.Auth(r.gRPCClient, r.user.GetInfo)).Methods("GET")
	r.router.HandleFunc("/{username}", middleware.Auth(r.gRPCClient, r.user.Profile)).Methods("GET")
}

func (r *Router) configureMiddlewares() {
	r.Use(middleware.CORS)
	r.Use(middleware.Logging)
}

func (r *Router) Use(mw middleware.Middleware) {
	r.middlewares = append(r.middlewares, mw)
}

func (r *Router) configureSwagger() {
	r.router.PathPrefix("/swagger/").HandlerFunc(httpSwagger.WrapHandler)

}

package httpserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"go-labs-game-platform/internal/bootstrap"
	"go-labs-game-platform/internal/config"
	"go-labs-game-platform/internal/httpserver/auth"
	"go-labs-game-platform/internal/httpserver/middleware"
	"go-labs-game-platform/internal/httpserver/rooms"
	"go-labs-game-platform/internal/httpserver/swagger"
	"go-labs-game-platform/internal/httpserver/user"
)

type Server struct {
	router *mux.Router
}

func New(deps *bootstrap.Dependencies) *Server {
	authHandlers := auth.New(deps.AuthSrv)
	userHandlers := user.New(deps.UserSrv)
	roomHandlers := rooms.New(deps.RoomSrv)

	router := mux.NewRouter()
	//router.Use(mux.CORSMethodMiddleware(router))
	router.Use(middleware.CORS)

	v1Router := router.PathPrefix(config.Get().HTTP.URLPathPrefix).Subrouter()

	publicChain := alice.New()

	// swagger documentation route
	router.PathPrefix("/swaggerui/").HandlerFunc(swagger.Handle)

	userChain := alice.New(middleware.Authorization(deps.AuthSrv))
	wsAuthChain := alice.New(middleware.SetTokenFromQuery, middleware.Authorization(deps.AuthSrv))

	v1Router.Handle("/register", publicChain.ThenFunc(authHandlers.Register)).Methods(http.MethodPost, http.MethodOptions)
	v1Router.Handle("/login", publicChain.ThenFunc(authHandlers.Login)).Methods(http.MethodPost, http.MethodOptions)

	v1Router.Handle("/user", userChain.ThenFunc(userHandlers.GetMe)).Methods(http.MethodGet, http.MethodOptions)

	v1Router.Handle("/room", userChain.ThenFunc(roomHandlers.GetList)).Methods(http.MethodGet, http.MethodOptions)
	v1Router.Handle("/room/create", userChain.ThenFunc(roomHandlers.Create)).Methods(http.MethodPost, http.MethodOptions)
	v1Router.Handle("/room/{room_id}", userChain.ThenFunc(roomHandlers.GetByID)).Methods(http.MethodGet, http.MethodOptions)
	v1Router.Handle("/room/{room_id}", userChain.ThenFunc(roomHandlers.Join)).Methods(http.MethodPost, http.MethodOptions)
	v1Router.Handle("/room/{room_id}", userChain.ThenFunc(roomHandlers.Leave)).Methods(http.MethodDelete, http.MethodOptions)

	v1Router.Handle("/game/{room_id}/ws", wsAuthChain.ThenFunc(roomHandlers.WebSocket)).Methods(http.MethodGet, http.MethodOptions)

	return &Server{
		router: router,
	}
}

func (s Server) Run() error {
	return http.ListenAndServe(s.Addr(), s.router)
}

func (s Server) Addr() string {
	return fmt.Sprintf("%s:%d", config.Get().HTTP.Host, config.Get().HTTP.Port)
}

func (s Server) Router() http.Handler {
	return s.router
}

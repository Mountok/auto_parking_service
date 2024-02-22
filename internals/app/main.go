package app

import (
	"context"
	"go_service_parking/example/api"
	"go_service_parking/example/api/middleware"
	db3 "go_service_parking/example/internals/app/db"
	"go_service_parking/example/internals/app/handlers"
	"go_service_parking/example/internals/app/processors"
	"go_service_parking/example/internals/cfg"
	"log"
	"net/http"
	"time"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	config cfg.Cfg
	ctx    context.Context
	srv    *http.Server
	db     *pgxpool.Pool
}

func NewServer(config cfg.Cfg, ctx context.Context) *Server {
	server := new(Server)
	server.ctx = ctx
	server.config = config
	return server
}
func (server *Server) Serve() {
	log.Println("Starting server")
	var err error
	server.db, err = pgxpool.Connect(server.ctx, server.config.GetDBString())
	if err != nil {
		log.Fatalln(err)
	}

	carsStorage := db3.NewCarStorage(server.db)
	usersStorage := db3.NewUsersStorage(server.db)

	carsProcessor := processors.NewCarsProcessor(carsStorage)
	usersProcessor := processors.NewUsersProcessor(usersStorage)

	carsHandler := handlers.NewCarsHandler(carsProcessor)
	userHandler := handlers.NewUsersHandler(usersProcessor)

	routes := api.CreateRoutes(userHandler, carsHandler)
	routes.Use(middleware.RequestLog)

	server.srv = &http.Server{
		Addr:    ":" + server.config.Port,
		Handler: routes,
	}

	log.Println("Server started")

	err = server.srv.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (server *Server) Shutdown() {
	log.Println("server stopped")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	server.db.Close()
	defer func() {
		cancel()
	}()
	var err error
	if err = server.srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalln("server shutdown failed err ",err.Error())
	}
	log.Println("Server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

}

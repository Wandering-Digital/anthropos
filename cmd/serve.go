package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	userhandler "github.com/Wandering-Digital/anthropos/domain/user/delivery/http"
	userrepository "github.com/Wandering-Digital/anthropos/domain/user/repository/postgres"
	userusecase "github.com/Wandering-Digital/anthropos/domain/user/usecase"
	"github.com/Wandering-Digital/anthropos/internal/cache"
	"github.com/Wandering-Digital/anthropos/internal/config"
	"github.com/Wandering-Digital/anthropos/internal/conn"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve run available servers such as: HTTP/JSON or gRPC",
		Long:  `Serve run available servers such as: HTTP/JSON or gRPC`,
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Println("Connecting database")
			if err := conn.ConnectDB(); err != nil {
				log.Fatalln(err)
			}
			log.Println("Database connected successfully!")

			log.Println("Connecting redis")
			if err := conn.ConnectRedis(); err != nil {
				log.Fatalln(err)
			}
			log.Println("Redis connected successfully!")

			log.Println("Initializing http-client")
			conn.InitClient()
			log.Println("Initialized http-client successfully!")
		},
		Run: serve,
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	// boot http server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// build and run http server
	srv := buildHTTP(cmd, args)

	go func(sr *http.Server) {
		if err := sr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}(srv)

	<-stop
	log.Println("Shutting down HTTP server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}

	log.Println("server shutdown successful!")
}

// buildHTTP register available handlers and return a http server
func buildHTTP(_ *cobra.Command, _ []string) *http.Server {
	r := chi.NewRouter()
	// middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/"))

	cfg := config.App()
	db := conn.GetDB()

	cacheCfg := config.Cache()

	cache := cache.NewRedis(conn.Redis(), cacheCfg.Prefix, cacheCfg.TTLDefault)
	_ = cache

	userRepo := userrepository.NewUser(db)

	userUseCase := userusecase.NewUser(userRepo)

	userhandler.NewUserHandler(r, userUseCase)

	log.Println("HTTP Listening on port: ", cfg.Port)
	log.Println("For system check use cURL request: ", "[curl localhost:"+fmt.Sprint(cfg.Port))

	return &http.Server{
		Addr:              fmt.Sprintf("%s:%d", config.App().Base, config.App().Port),
		Handler:           r,
		ReadHeaderTimeout: cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}
}

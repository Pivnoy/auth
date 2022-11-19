package app

import (
	"auth_reg/config"
	v1 "auth_reg/internal/controller/http/v1"
	"auth_reg/internal/usecase"
	"auth_reg/internal/usecase/repo"
	"auth_reg/pkg/httpserver"
	"auth_reg/pkg/postgres"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {

	pg, err := postgres.New(cfg)

	if err != nil {
		log.Fatal("Error in creating postgres instance")
	}

	usRepo := repo.NewUserRepo(pg)
	sqRepo := repo.NewSecretQuestionRepo(pg)
	sq := usecase.NewSecretQuestionUseCase(sqRepo)
	us := usecase.NewUserUseCase(usRepo)
	rg := usecase.NewRegisterUseCase(us)
	jc := usecase.NewJwtUseCase(us, cfg.JwtSecret)

	handler := gin.New()
	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	v1.NewRouter(handler,
		us,
		rg,
		jc,
		sq)

	serv := httpserver.New(handler, httpserver.Port(cfg.Port))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interruption:
		log.Printf("signal: " + s.String())
	case err = <-serv.Notify():
		log.Printf("Notify from http server")
	}

	err = serv.Shutdown()
	if err != nil {
		log.Printf("Http server shutdown")
	}
}

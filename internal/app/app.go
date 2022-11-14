package app

import (
	"auth_reg/config"
	v1 "auth_reg/internal/controller/http/v1"
	"auth_reg/internal/usecase"
	"auth_reg/internal/usecase/repo"
	"auth_reg/pkg/httpserver"
	"auth_reg/pkg/mongodb"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {

	mng, err := mongodb.NewMongo(cfg)
	if err != nil {
		log.Fatal("Cannot connect to Mongo")
	}

	usRepo := repo.NewUserRepo(mng, "accounts")
	us := usecase.NewUserUseCase(usRepo)
	rg := usecase.NewRegisterUseCase(us)
	jc := usecase.NewJwtUseCase(us, cfg.JwtSecret)

	handler := gin.New()

	v1.NewRouter(handler,
		us,
		rg,
		jc)

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

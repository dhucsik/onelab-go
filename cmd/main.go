package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"practice/config"
	_ "practice/docs"
	"practice/service"
	"practice/storage"
	"practice/transport/http"
	"practice/transport/http/handler"
)

// @title 	Book microservice
// @version	1.0
// @description Simple microservice for borrowing books with jwt auth
// @termsOfService http://swagger.io/terms/

// @contact.name Dias Galikhanov
// @contact.email diasgalikhanov@gmail.com

// @host localhost:8586
// @BasePath /api/v1

// User + Auth CRUD + Swagger
// Book CRUD + Swagger

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log.Fatalln(run())
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gracefullyShutdown(cancel)
	conf, err := config.New()
	if err != nil {
		return err
	}

	stg, err := storage.New(ctx, conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	svc, svcErr := service.NewManager(stg)
	if svcErr != nil {
		return svcErr
	}

	h := handler.NewManager(svc)
	HTTPServer := http.NewServer(conf, h)

	return HTTPServer.StartHTTPServer(ctx)
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}

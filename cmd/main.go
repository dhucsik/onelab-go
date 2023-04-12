package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"practice/config"
	"practice/service"
	"practice/storage"
	"practice/transport/http"
	"practice/transport/http/handler"
)

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

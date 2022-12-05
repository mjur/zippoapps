package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/braintree/manners"
	"github.com/julienschmidt/httprouter"
	"github.com/mjur/zippo/cmd/config"
	"github.com/mjur/zippo/pkg/cache"
	"github.com/mjur/zippo/pkg/configuration"
	handlers "github.com/mjur/zippo/pkg/configuration/http"
	httpMiddleware "github.com/mjur/zippo/pkg/configuration/http/middleware"
	"github.com/mjur/zippo/pkg/configuration/middleware"
	storeMiddleware "github.com/mjur/zippo/pkg/configuration/store/middleware"
)

func main() {
	c, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	cache := cache.New(time.Now, sync.Map{})
	store, err := config.NewStore(*c)
	if err != nil {
		log.Fatal("error connecting to db:" + err.Error())
	}
	store = storeMiddleware.NewCacheMiddleware(c, store, cache)

	randomNumberGenerator := configuration.NewRandomNumberGenerator(time.Now)

	service := configuration.New(c, store, randomNumberGenerator)
	service = middleware.NewLogMiddleware(c, service, c.Log)
	handler := handlers.New(c, service)
	handler = httpMiddleware.NewLogMiddleware(c, handler, c.Log)

	router := httprouter.New()

	router.GET("/configuration/:package", handler.GetMainSku)
	timeoutrouter := http.TimeoutHandler(router, (time.Second * time.Duration(c.Timeout)), "The request timed out")
	addr := fmt.Sprintf("%s:%s", c.Host, c.Port)

	fmt.Printf("Server listening on:%s", addr)

	server := manners.NewServer()
	server.Addr = addr
	server.Handler = timeoutrouter

	errChan := make(chan error)

	go func() {
		c.Log.Infof("listening on %s", addr)
		errChan <- server.ListenAndServe()
	}()

	err = <-errChan
	if err != nil {
		log.Fatal(err)
	}
}

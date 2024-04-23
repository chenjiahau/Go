package main

import (
	"flag"
	"fmt"
	"net/http"

	"ivanfun.com/mis/config"
	"ivanfun.com/mis/handler"
	"ivanfun.com/mis/router"
	"ivanfun.com/mis/util"
)

var (
	AppName string
	Version string
)

func main() {
	flag.Parse()
  AppName = flag.Arg(0)
	Version = flag.Arg(1)

	c := handler.NewConfig(AppName, Version)
	handler.NewHandler(c)
	RunServer(c)
}

func RunServer(c *handler.Config) {
	srv := &http.Server{
		Addr:    config.Server["Addr"],
		Handler: router.GetRoutes(),
	}

	logMessage := fmt.Sprintf("Server is running on %s\n", config.Server["Addr"])
	util.WriteInfoLog(logMessage)

	err := srv.ListenAndServe()
	if err != nil {
		util.WriteErrorLog(err.Error())
	}
}

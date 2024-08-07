package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"ivanfun.com/mis/db/driver"
	"ivanfun.com/mis/internal/config"
	"ivanfun.com/mis/internal/handler"
	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/router"
	"ivanfun.com/mis/internal/util"
)

var (
	appName	string
	version	string
)

func main() {
	flag.Parse()
	appName = flag.Arg(0)
	version = flag.Arg(1)

	pgConn, err := driver.ConnectSQL(driver.PostgreSQLDataSourceName)
	if err != nil {
		log.Fatal("cannot connect to database")
	}
	model.NewDbConfig(pgConn)
	defer pgConn.SQL.Close()

	c := handler.NewConfig(appName, version)
	handler.NewHandler(c)
	RunServer(c)
}

func RunServer(c *handler.Config) {
	srv := &http.Server{
		Addr:			config.Server["Addr"],
		Handler:	router.GetRoutes(),
	}

	logMessage := fmt.Sprintf("Server is running on %s\n", config.Server["Addr"])
	util.WriteInfoLog(logMessage)

	err := srv.ListenAndServe()
	if err != nil {
		util.WriteErrorLog(err.Error())
	}
}

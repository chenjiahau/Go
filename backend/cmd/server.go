package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"ivanfun.com/mis/config"
	"ivanfun.com/mis/db/driver"
	"ivanfun.com/mis/handler"
	"ivanfun.com/mis/router"
	"ivanfun.com/mis/util"
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
	dbConf := handler.DbConf{
		PgConn: pgConn,
	}
	defer pgConn.SQL.Close()

	c := handler.NewConfig(appName, version, &dbConf)
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

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
	appName			string
	appVersion	string
	dbHost			string
	dbUser			string
	dbPass			string
)

func main() {
	flag.Parse()
	appName = flag.Arg(0)
	appVersion = flag.Arg(1)
	dbHost = flag.Arg(2)
	dbUser = flag.Arg(3)
	dbPass = flag.Arg(4)

	var systemInfo string
	if appName == "" || appVersion == "" {
		systemInfo = driver.SystemInfo
	} else {
		systemInfo = fmt.Sprintf("%s Version %s", appName, appVersion)
	}
	util.WriteInfoLog(systemInfo)

	var dbConnect string
	if dbHost == "" || dbUser == "" || dbPass == "" {
		dbConnect = driver.PostgreSQLDataSourceName
	} else {
		dbConnect = fmt.Sprintf("postgres://%s:%s@%s:5432/mis?sslmode=disable", dbUser, dbPass, dbHost)
	}

	pgConn, err := driver.ConnectSQL(dbConnect)
	if err != nil {
		util.WriteErrorLog(err.Error())
		log.Fatal("cannot connect to database")
	}
	model.NewDbConfig(pgConn)
	defer pgConn.SQL.Close()

	c := handler.NewConfig(appName, appVersion)
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

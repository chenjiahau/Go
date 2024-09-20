package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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
	dbName			string
	dbUser			string
	dbPass			string
	emailHost		string
	emailPort		string
	emailFrom		string
	emailPass		string
)

func main() {
	flag.Parse()
	appName = flag.Arg(0)
	appVersion = flag.Arg(1)
	dbHost = flag.Arg(2)
	dbName = flag.Arg(3)
	dbUser = flag.Arg(4)
	dbPass = flag.Arg(5)
	emailHost = flag.Arg(6)
	emailPort = flag.Arg(7)
	emailFrom = flag.Arg(8)
	emailPass = flag.Arg(9)

	if appName == "" || appVersion == "" || dbHost == "" || dbUser == "" || dbPass == "" || emailHost == "" || emailPort == "" || emailFrom == "" || emailPass == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		appName = os.Getenv("APPLICATION_NAME")
		appVersion = os.Getenv("APPLICATION_VERSION")
		dbHost = os.Getenv("POSTGRES_HOST")
		dbName = os.Getenv("POSTGRES_DB")
		dbUser = os.Getenv("POSTGRES_USER")
		dbPass = os.Getenv("POSTGRES_PASSWORD")
		emailHost = os.Getenv("EMAIL_HOST")
		emailPort = os.Getenv("EMAIL_PORT")
		emailFrom = os.Getenv("EMAIL_USER")
		emailPass = os.Getenv("EMAIL_PASSWORD")
	}

	// Application info
	systemInfo := fmt.Sprintf("%s Version %s", appName, appVersion)
	util.WriteInfoLog(systemInfo)

	// Database configuration
	dbConnect := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", dbUser, dbPass, dbHost, dbName)

	pgConn, err := driver.ConnectSQL(dbConnect)
	if err != nil {
		util.WriteErrorLog(err.Error())
		log.Fatal("cannot connect to database")
	}
	model.NewDbConfig(pgConn)
	defer pgConn.SQL.Close()

	// Email configuration
	ePort, err := strconv.ParseInt(emailPort, 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		log.Fatal("cannot parse email port")
	}
	emailConf := handler.NewEmailConfig(emailHost, int(ePort), emailFrom, emailPass)

	// Server configuration
	c := handler.NewConfig(appName, appVersion, emailConf)
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

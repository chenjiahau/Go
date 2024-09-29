package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	portalUrl			string
	appName				string
	appVersion		string
	dbHost				string
	dbName				string
	dbUser				string
	dbPass				string
	emailHost			string
	emailPort			string
	emailFrom			string
	emailPass			string
	domain				string
	awsRegion			string
	awsAccessKey	string
	awsSecretKey	string
	awsBucketName	string
)

func main() {
	flag.Parse()
	portalUrl = flag.Arg(0)
	appName = flag.Arg(1)
	appVersion = flag.Arg(2)
	dbHost = flag.Arg(3)
	dbName = flag.Arg(4)
	dbUser = flag.Arg(5)
	dbPass = flag.Arg(6)
	emailHost = flag.Arg(7)
	emailPort = flag.Arg(8)
	emailFrom = flag.Arg(9)
	emailPass = flag.Arg(10)
	awsRegion = flag.Arg(11)
	awsAccessKey = flag.Arg(12)
	awsSecretKey = flag.Arg(13)
	awsBucketName = flag.Arg(14)

	if portalUrl == "" || appName == "" || appVersion == "" || dbHost == "" || dbUser == "" || dbPass == "" || emailHost == "" || emailPort == "" || emailFrom == "" || emailPass == "" {
		cwd, err := os.Getwd()
    if err != nil {
        log.Fatalf("Error getting current working directory: %v", err)
    }
		envPath := filepath.Join(cwd, ".env")

		err = godotenv.Load(envPath)
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		portalUrl = os.Getenv("PORTAL_URL")
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
		awsRegion = os.Getenv("AWS_REGION")
		awsAccessKey = os.Getenv("AWS_ACCESS_KEY")
		awsSecretKey = os.Getenv("AWS_SECRETE_KEY")
		awsBucketName = os.Getenv("AWS_BUCKET_NAME")
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

	// AWS configuration
	awsConf := handler.NewAWSConfig(awsRegion, awsAccessKey, awsSecretKey, awsBucketName)

	// Server configuration
	c := handler.NewConfig(portalUrl, appName, appVersion, emailConf, awsConf)
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

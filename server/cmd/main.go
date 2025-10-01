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
	portalUrl						string
	appName							string
	appVersion					string
	dbHost							string
	dbName							string
	dbUser							string
	dbPass							string
	emailHost						string
	emailPort						string
	emailFrom						string
	emailPass						string
	awsRegion						string
	awsAccessKey				string
	awsSecretKey				string
	awsBucketName				string
	recaptchaSecretKey	string
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
	smtpHost := flag.Arg(7)
	smtpPort := flag.Arg(8)
	smtpUsername := flag.Arg(9)
	smtpPassword := flag.Arg(10)
	smtpSender := flag.Arg(11)
	awsRegion = flag.Arg(12)
	awsAccessKey = flag.Arg(13)
	awsSecretKey = flag.Arg(14)
	awsBucketName = flag.Arg(15)
	recaptchaSecretKey = flag.Arg(16)

	if portalUrl == "" || appName == "" || appVersion == "" ||
	dbHost == "" || dbUser == "" || dbPass == "" ||
	smtpHost == "" || smtpPort == "" || smtpUsername == "" || smtpPassword == "" || smtpSender == "" ||
	awsRegion == "" || awsAccessKey == "" || awsSecretKey == "" || awsBucketName == "" ||
	recaptchaSecretKey == "" {
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
		smtpHost = os.Getenv("SMTP_HOST")
		smtpPort = os.Getenv("SMTP_PORT")
		smtpUsername = os.Getenv("SMTP_USERNAME")
		smtpPassword = os.Getenv("SMTP_PASSWORD")
		smtpSender = os.Getenv("SMTP_SENDER")
		awsRegion = os.Getenv("AWS_REGION")
		awsAccessKey = os.Getenv("AWS_ACCESS_KEY")
		awsSecretKey = os.Getenv("AWS_SECRETE_KEY")
		awsBucketName = os.Getenv("AWS_BUCKET_NAME")
		recaptchaSecretKey = os.Getenv("RECAPTCHA_SECRET_KEY")
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
	ePort, err := strconv.ParseInt(smtpPort, 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		log.Fatal("cannot parse email port")
	}
	smtpConf := handler.NewSmtpConfig(smtpHost, int(ePort), smtpUsername, smtpPassword, smtpSender)

	// AWS configuration
	awsConf := handler.NewAWSConfig(awsRegion, awsAccessKey, awsSecretKey, awsBucketName)

	// Recaptcha configuration
	recaptchaConf := handler.NewRecaptchaConfig(recaptchaSecretKey)

	// Server configuration
	c := handler.NewConfig(portalUrl, appName, appVersion, smtpConf, awsConf, recaptchaConf)
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

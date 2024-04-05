package main

import (
	"log"
	"net/http"

	"example.com/project/config"
	"example.com/project/db/driver"
	"example.com/project/page"
	"example.com/project/render"
	"example.com/project/router"
)

const portNumber = ":8080"

func main() {
	// Application wide config
	var appConfig config.AppConfig

	// Create template cache for application
	tmpCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	appConfig.TemplateCache = tmpCache
	appConfig.UseCache = false

	// DB config
	pgConn, err := driver.ConnectSQL(config.PostgreSQLDataSourceName)
	if err != nil {
		log.Fatal("cannot connect to database")
	}
	dbConfig := config.DbConfig{ PgConn: pgConn }

	// Set the appConfig and dbConfig for the page package
	repo := page.NewRepo(&appConfig, &dbConfig)
	page.NewHandler(repo)

	// Set the appConfig for the render package
	render.NewAppConfig(&appConfig)

	// Use Chi router
	log.Println("Starting application on port", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: router.GetRoutes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

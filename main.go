package main

import (
	"log"
	"net/http"

	"example.com/project/config"
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

	// appConfig for the page package
	page.NewAppConfig(&appConfig)

	// Set the appConfig for the render package
	render.NewTemplates(&appConfig)

	// Use Chi router
	log.Println("Starting application on port", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: router.GetRouter(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

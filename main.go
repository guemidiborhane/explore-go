package main

import (
	"github.com/guemidiborhane/explore-go/config"
	"github.com/guemidiborhane/explore-go/database"
	"github.com/guemidiborhane/explore-go/pkg"
	"github.com/guemidiborhane/explore-go/router"
	"github.com/guemidiborhane/explore-go/server"
)

func main() {
	config.Load()
	server.Setup()
	database.Setup()

	// Required to run first since it registers the group
	// all other packages are gonna register their routes on
	router.Setup()
	pkg.Setup()
	SetupStatic()
	server.Start()
}

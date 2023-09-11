package main

import (
	"github.com/guemidiborhane/explore-go/config"
	"github.com/guemidiborhane/explore-go/http"
	"github.com/guemidiborhane/explore-go/modules"
	"github.com/guemidiborhane/explore-go/routes"
)

func main() {
	config.Load()
	modules.Setup()
	routes.Setup()
	http.Start()
}

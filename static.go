package main

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"explore-go/server"
)

//go:embed all:build
var vite embed.FS

func SetupStatic() {
	server.App.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(vite),
		Browse:       true,
		PathPrefix:   "/build",
		Index:        "index.html",
		NotFoundFile: "/build/index.html",
	}))
}

package main

import (
	"embed"
	"explore-go/server"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed all:build
var vite embed.FS

func SetupStatic() {
	server.App.Use("/", server.CsrfMiddleware, filesystem.New(filesystem.Config{
		Root:         http.FS(vite),
		Browse:       true,
		PathPrefix:   "/build",
		Index:        "index.html",
		NotFoundFile: "/build/index.html",
	}))
}

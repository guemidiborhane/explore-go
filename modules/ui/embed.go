package ui

import (
	"core/server"
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed all:build
var vite embed.FS

func Setup() {
	app := server.Application

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(vite),
		Browse:       true,
		PathPrefix:   "/build",
		Index:        "index.html",
		NotFoundFile: "/build/index.html",
	}))
}

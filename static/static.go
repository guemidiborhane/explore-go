package static

import (
	"embed"
	"explore-go/router"
	"explore-go/server"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed all:build
var static embed.FS

func Setup() {
	server.App.Use("/", router.CsrfMiddleware, filesystem.New(filesystem.Config{
		Root:         http.FS(static),
		Browse:       true,
		PathPrefix:   "/build",
		Index:        "index.html",
		NotFoundFile: "/build/index.html",
	}))
}

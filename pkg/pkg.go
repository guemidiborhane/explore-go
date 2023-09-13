package pkg

import (
	"github.com/guemidiborhane/explore-go/database"
	"github.com/guemidiborhane/explore-go/database/session"
	"github.com/guemidiborhane/explore-go/pkg/links"
	"github.com/guemidiborhane/explore-go/pkg/setup"
	"github.com/guemidiborhane/explore-go/router"
	"github.com/guemidiborhane/explore-go/server"
)

func registerPkgs(args *setup.SetupArgs) {
	links.Setup(args)
}

func Setup() {
	registerPkgs(&setup.SetupArgs{
		Application:  server.App,
		Router:       &router.ApiRouter,
		Database:     database.DB,
		SessionStore: session.Session,
	})
}

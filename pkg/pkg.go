package pkg

import (
	"explore-go/database"
	"explore-go/database/session"
	"explore-go/pkg/auth"
	"explore-go/pkg/links"
	"explore-go/pkg/setup"
	"explore-go/router"
	"explore-go/server"
)

func registerPkgs(args *setup.SetupArgs) {
	links.Setup(args)
	auth.Setup(args)
}

func Setup() {
	registerPkgs(&setup.SetupArgs{
		Application: server.App,
		Router:      &router.ApiRouter,
		Database:    database.DB,
		Session:     session.Session,
	})
}

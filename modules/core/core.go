package core

import (
	"core/database"
	"core/router"
	"core/server"
)

func Setup() {
	database.Setup()
	server.Setup()
	router.Setup()
}

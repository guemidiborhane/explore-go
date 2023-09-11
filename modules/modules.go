package modules

import (
	"core"
	"links"
	"ui"
)

func Setup() {
	core.Setup()
	links.Setup()
	ui.Setup()
}

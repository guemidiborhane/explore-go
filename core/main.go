package core

import "core/config/initializers"

func Setup() {
	initializers.InitDB()
	initializers.InitServer()
}

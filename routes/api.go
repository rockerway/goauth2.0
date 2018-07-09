package routes

import (
	"oauth/packages/router"
)

// Open method to register router url
func Open() {
	router.RegisterByString("/activity/join", "ActivityController", "Join")
	router.Start()
}

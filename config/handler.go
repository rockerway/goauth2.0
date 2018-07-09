package config

import (
	"oauth/handlers"
)

// HandlerMap is the global class map
var HandlerMap map[string]*handlers.HandlerInterface = make(map[string]*handlers.HandlerInterface)

func registerClass(name string, handler handlers.HandlerInterface) {
	HandlerMap[name] = &handler
}

func init() {
	registerClass("Handler", &handlers.Handler{})
}

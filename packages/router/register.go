package router

import (
	"net/http"
	"oauth/config"
	"reflect"
)

var router = make(map[string]func(http.ResponseWriter, *http.Request))

// Start method is web router's entry point.
func Start() {
	for path, method := range router {
		http.HandleFunc(path, method)
	}
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

// RegisterByHandleFunc is the method to define route rule by HandleFunc
func RegisterByHandleFunc(path string, method func(res http.ResponseWriter, req *http.Request)) {
	router[path] = method
}

// RegisterByString is the method to define route rule by handler name and method name
func RegisterByString(path, handlerName, methodName string) {
	handler := config.HandlerMap[handlerName]
	handlerType := reflect.TypeOf(*handler)
	handlerValue := reflect.ValueOf(*handler)
	_, state := handlerType.MethodByName(methodName)
	if state {
		method := handlerValue.MethodByName(methodName)
		RegisterByHandleFunc(path, func(res http.ResponseWriter, req *http.Request) {
			method.Call([]reflect.Value{reflect.ValueOf(res), reflect.ValueOf(req)})
		})
	}
}

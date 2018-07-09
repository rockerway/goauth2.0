package handlers

import "net/http"

// HandlerInterface to define router sub element
type HandlerInterface interface {
	Get(req *http.Request, res http.ResponseWriter)
	Post(req *http.Request, res http.ResponseWriter)
}

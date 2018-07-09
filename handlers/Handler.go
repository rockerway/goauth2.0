package handlers

import (
	"net/http"
)

// Handler struct
type Handler struct {
}

// Get function
func (h *Handler) Get(req *http.Request, res http.ResponseWriter) {

}

// Post function
func (h *Handler) Post(req *http.Request, res http.ResponseWriter) {

}

func (h *Handler) clientCredential(req *http.Request, res http.ResponseWriter) {
	if req.Method == "POST" {

	} else if req.Method == "GET" {

	} else {

	}
}

package service

import (
	"github.com/julienschmidt/httprouter"
)

/*
creates a new router instance and iterate through all the Routes to get each’s
Route’s Method, Pattern and Handle and registers a new request handle.
*/
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Path, route.Handle)
	}
	return router
}

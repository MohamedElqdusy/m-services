package service

import (
	"app/authentication"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Method string            //HTTP method
	Path   string            //url endpoint
	Handle httprouter.Handle //Controller function which dispatches the right HTML page and/or data for each route
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		Index,
	},
	Route{
		"GET",
		"/patient/:id",
		FindPatient,
	},
	Route{
		"GET",
		"/patients",
		FindPatients,
	},
	Route{
		"GET",
		"/doctor/:id",
		FindDoctor,
	},
	Route{
		"GET",
		"/doctors",
		FindDoctors,
	},
	Route{
		"POST",
		"/register",
		Register,
	},
	Route{
		"POST",
		"/reserve",
		ReserveAppoinment,
	},
	Route{
		"GET",
		"/appoinment/doctor/:id",
		FindAllAppoinmentsForDoctor,
	},
	Route{
		"GET",
		"/jwt-token",
		authentication.GenerateJwtToken,
	},
	Route{
		"GET",
		"/authindex",
		authentication.AuthenticationMiddleware(Index),
	},
}

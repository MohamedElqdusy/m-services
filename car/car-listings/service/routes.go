package service

import (
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
		"/cars/make/:make/",
		SearchCarsByMake,
	},
	Route{
		"GET",
		"/cars/model/:model/",
		SearchCarsByModel,
	},
	Route{
		"GET",
		"/cars/year/:year/",
		SearchCarsByYear,
	},
	Route{
		"GET",
		"/cars/color/:color/",
		SearchCarsByColor,
	},
	Route{
		"GET",
		"/cars/search/",
		AllCars,
	},
	Route{
		"POST",
		"/upload_csv/:dealer_id/",
		saveCarByDealer,
	},
	Route{
		"POST",
		"/vehicle_listings/",
		SaveCars,
	},
}

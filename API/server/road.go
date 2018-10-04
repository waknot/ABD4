/*
 * File: road.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 9:46:55 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package server

import (
	"ABD4/API/context"
	"ABD4/API/handler"
	"ABD4/API/utils"

	"github.com/gorilla/mux"
)

const (
	GET     = "GET"
	POST    = "POST"
	OPTIONS = "OPTIONS"
	PUT     = "PUT"
	DELETE  = "DELETE"
)

// RoadGetter type for router definition
type RoadGetter func() []*Road

// Road define what a router entry must embed
type Road struct {
	Name            string
	Method          string
	Pattern         string
	StatusProtected bool
	HandlerFunc     context.CustomHandler
}

func (r *Road) appendTo(ctx *context.AppContext, preparedHandler *context.HandlerWrapper, router *mux.Router) {
	ctx.Log.Info.Printf("%s __ Load... __ %s %s%s", utils.Use().GetStack(r.appendTo), r.Method, ctx.Opts.GetAddress(), r.Name)
	router.StrictSlash(true).
		Methods(r.Method).
		Name(r.Name).
		Path(r.Pattern).
		Handler(preparedHandler)
}

// GetRoadKit must return a map wih a key for the
// mux.Router router.Pathprefix(key).Subrouter() method as key (see appendRoadKit in router.go)
// and a roadGetter function as value
func GetRoadKit() map[string]RoadGetter {
	return map[string]RoadGetter{
		// exemple for localhost:8000/user/* set of road:
		"/user":    getUserRouting,
		"/auth":    getAuthRouting,
		"/backup":  getBackupRouting,
		"/elastic": getElasticRouting,
	}
}

// getUserRouting return the /user routing
func getUserRouting() []*Road {
	return []*Road{
		{
			Name:            "/user",
			Method:          GET,
			Pattern:         "/",
			StatusProtected: false,
			HandlerFunc:     handler.GetUsers,
		},
		/*
			 * exemple for GET /user/{id}
			 * retrieve id value with vars := mux.Vars(*http.Request),
			 * id will be indexed in vars["id"]
			{
				Name: "GET /user/{id}",
				Method: GET,
				Pattern: "/{id}",
				StatusProtected: false,
				HandlerFunc: handler.?,
			},
			 *
		*/
		/*
			 * exemple for POST /user
			 * retrieve data in *http.Request.body
			 * see model.User.UnmarshalFromRequest method
			 * a POST on user exist in authentication.go/Register
			{
				Name: "POST /user",
				Method: POST,
				Pattern: "/",
				StatusProtected: false,
				HandlerFunc: handler.?,
			},
			 *
		*/
	}
}

func getAuthRouting() []*Road {
	return []*Road{
		{
			Name:            "/auth/login",
			Method:          POST,
			Pattern:         "/login",
			StatusProtected: false,
			HandlerFunc:     handler.Login,
		},
		{
			Name:            "/auth/register",
			Method:          POST,
			Pattern:         "/register",
			StatusProtected: false,
			HandlerFunc:     handler.Register,
		},
	}
}

func getBackupRouting() []*Road {
	return []*Road{
		{
			Name:            "/backup",
			Method:          GET,
			Pattern:         "/",
			StatusProtected: false,
			HandlerFunc:     handler.BackupBoltDatabase,
		},
	}
}

func getElasticRouting() []*Road {
	return []*Road{
		{
			Name:            "/elastic",
			Method:          GET,
			Pattern:         "/",
			StatusProtected: false,
			HandlerFunc:     handler.GetElasticHealth,
		},
	}
}

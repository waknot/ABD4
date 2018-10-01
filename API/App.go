/*
 * File: README.md
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 5:47:26 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package main

import (
	"ABD4/API/context"
	"ABD4/API/server"
	"ABD4/API/utils"
	"net/http"
)

type App struct {
	Ctx *context.AppContext
}

// Initialize launch the server, making the match between server/model/handler and context/database/logger
func (a *App) Initialize(opts *server.Option) error {
	a.Ctx = &context.AppContext{}
	a.Ctx.Instanciate(opts)
	a.Ctx.Log.Info.Printf("%s API starting...", utils.Use().GetStack(a.Initialize))
	router := server.Routing(a.Ctx)
	http.Handle("/", router)
	return http.ListenAndServe(opts.GetAddress(), nil)
}

// InitializeAndWaitForSignal launch the server, making the match between server/model/handler and context/database/logger
// this function need to be used with a channel assuring that Ctx is instanciated (see main_test.go)
func (a *App) InitializeAndWaitForSignal(opts *server.Option, signal chan bool) error {
	a.Ctx = &context.AppContext{}
	a.Ctx.Instanciate(opts)
	a.Ctx.Log.Info.Printf("%s API starting...", utils.Use().GetStack(a.InitializeAndWaitForSignal))
	router := server.Routing(a.Ctx)
	http.Handle("/", router)
	signal <- true
	return http.ListenAndServe(opts.GetAddress(), nil)
}

/*
 * File: elasticHandler.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 2:36:54 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package handler

import (
	"ABD4/API/context"
	"net/http"
)

// GetElasticHealth handler for Get /elastic
func GetElasticHealth(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	response := ctx.Rw.NewResponse(200, string("usersJSON"), "", "")
	response.SendItSelf(ctx, w)
	return
}

/*
 * File: transaction.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 15th October 2018 7:42:41 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 16th October 2018 5:05:27 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package handler

import (
	"ABD4/API/context"
	"ABD4/API/model"
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTransaction(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	ctx.Log.Info.Printf("%s %s ", utils.Use().GetStack(GetTransaction), "Add Transaction process")
	tx, err := ctx.TransactionManager.FindAll()
	if err != nil {
		ctx.Log.Error.Printf("%s %s", utils.Use().GetStack(GetTransaction), err.Error())
		return
	}
	txJSON, err := json.Marshal(tx)
	if err != nil {
		msg := fmt.Sprintf("%s marshal failed: %s", utils.Use().GetStack(GetUsers), err.Error())
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "encoding users failed", msg)
		return
	}
	response := ctx.Rw.NewResponse(200, string(txJSON), "", "")
	response.SendItSelf(ctx, w)
	return
}

func AddTransaction(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	transaction := &model.Transaction{}
	err := transaction.UnmarshalFromRequest(r)
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "Decode request data failed", err.Error())
		return
	}
	transaction, err = ctx.TransactionManager.Create(transaction)
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusBadRequest, "Insert transaction in mongo failed", err.Error())
		return
	}
	ctx.Rw.Send(ctx, w, http.StatusCreated, transaction, "", "")
	return
}

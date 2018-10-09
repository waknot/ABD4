/*
 * File: authentication.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 5:50:33 pm
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

// Login handler, allow to get a token if authentication goes well
// Must receive: {"email": string, "password": string}
func Login(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	isTrusted := false
	err := user.UnmarshalFromRequest(r)
	if nil != err {
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "can't interpret users data", fmt.Sprintf("%s %s", utils.Use().GetStack(Login), err.Error()))
		return
	}
	// Replace this local fake implementation by a database implementation
	users, err := ctx.UserManager.Seek()
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "retrieving users failed", fmt.Sprintf("%s failed to seek users %s", utils.Use().GetStack(Login), err.Error()))
		return
	}
	for _, u := range users {
		if u.Email == user.Email && u.Password == user.Password {
			isTrusted = true
			user = u
			ctx.Log.Info.Printf("%s user identified: %v", utils.Use().GetStack(Login), u)
		}
	}
	if user.ID != "" && isTrusted {
		claim := &model.Claim{}
		if user.Claim != "" {
			// L'utilisateur est connu des services et a des claims
			err = json.Unmarshal([]byte(user.Claim), claim)
			if err != nil {
				ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "decoding existing informations failed", fmt.Sprintf("%s %s", utils.Use().GetStack(Login), err.Error()))
				return
			}
		} else {
			if err := claim.NewFromUser(user); err != nil {
				ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "generate claim failed", fmt.Sprintf("%s %s", utils.Use().GetStack(Login), err.Error()))
				return
			}
		}
		token, err := claim.GetToken()
		if err != nil {
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "generate token failed", fmt.Sprintf("%s %s", utils.Use().GetStack(Login), err.Error()))
			return
		}
		ctx.SessionUser = user
		ctx.Log.Info.Printf("%s login process ok, token: %s", utils.Use().GetStack(Login), token)
		response := ctx.Rw.NewResponse(200, token, "", "")
		response.SendItSelf(ctx, w)
		return
	}
	msg := fmt.Sprintf("%s invalid user in request", utils.Use().GetStack(Login))
	ctx.Rw.SendError(ctx, w, 401, msg, "")
	return
}

// Register manage the inscription of a user
// Must receive: {"email": string, "password": string}
func Register(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	err := user.UnmarshalFromRequest(r)
	if nil != err {
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "can't interpret user data", fmt.Sprintf("%s %s", utils.Use().GetStack(Register), err.Error()))
		return
	}
	user, err = ctx.UserManager.NewUser(user.Name, user.Email, user.Password, user.Permission, user.Claim)
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "error saving user", fmt.Sprintf("%s %s", utils.Use().GetStack(Register), err.Error()))
		return
	}
	ctx.Rw.Send(ctx, w, 200, user, "", "")
	return
}

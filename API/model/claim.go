/*
 * File: claim.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 4:14:40 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package model

import (
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	appName = "abd4"
	// SECRET used in jwt signature
	SECRET = "==+ABD4@VDMEscapeGame13683"
)

// Claim is the type used to parse
// a unique Json Web Token, we can stock what we want,
// here the user id to retrieve it during authentication process
type Claim struct {
	User string `json:"user,string"`
	jwt.StandardClaims
}

func (c *Claim) toString() (string, error) {
	json, err := json.Marshal(c)
	if err != nil {
		return "", fmt.Errorf("<<<< <<<< %s %s", utils.Use().GetStack(c.toString), err.Error())
	}
	return string(json), nil
}

// NewFromUser fullfill the Claim object with
// the user information, set the token validity to one week
func (c *Claim) NewFromUser(user *User) error {
	var err error

	exp := time.Now()
	c.User = user.ToString()
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(c.NewFromUser), err.Error())
	}
	exp = exp.Add(time.Minute * 60 * 24 * 7) // token is set for a week
	c.ExpiresAt = exp.Unix()
	c.Issuer = fmt.Sprintf("%s for %s", appName, user.Email)
	user.Claim, err = c.toString()
	return err
}

// GetToken generate a new token based on relying Claim
func (c *Claim) GetToken() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenStr, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", fmt.Errorf("<<<< %s Error generating jwt: %s", utils.Use().GetStack(c.GetToken), err.Error())
	}
	return tokenStr, nil
}

// Marshal to implement ISerial interface
func (c *Claim) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

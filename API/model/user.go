/*
 * File: user.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 8:19:21 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package model

import (
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// User is composed:
// Claim are used for jwt token signature
// Permission
type User struct {
	ID         string    `json:"id,omitempty"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Permission string    `json:"permission,omitempty"`
	Claim      string    `json:"claim,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}

// ToString return string conversion of marshal user
// absorb error...
func (u *User) ToString() string {
	ret, _ := u.Marshal()
	return string(ret)
}

// UnmarshalFromRequest take a request object supposed to contain
// a user data object, extract it, convert to User and send back
// the string representation of the content
func (u *User) UnmarshalFromRequest(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(u.UnmarshalFromRequest), err.Error())
	}
	err = json.Unmarshal(body, u)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(u.UnmarshalFromRequest), err.Error())
	}
	return nil
}

// GetId implement ISessionUser
func (u *User) GetID() string {
	return u.ID
}

// Marshal implement ISerial
func (u *User) Marshal() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) OrderById(results []*User, i, j int) bool {
	return strings.Compare(results[i].GetID(), results[j].GetID()) == 1
}

func (u *User) OrderByEmail(results []*User, i, j int) bool {
	return strings.Compare(results[i].Email, results[j].Email) == 1
}

/*
 * File: acheteur.go
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
)

// Acheteur is composed:
type Acheteur struct {
	Spectateur
	Email 		string    	`json:"Email"`
}

// ToString return string conversion of marshal user
// absorb error...
func (a *Acheteur) ToString() string {
	ret, _ := a.Marshal()
	return string(ret)
}

// UnmarshalFromRequest take a request object supposed to contain
// a user data object, extract it, convert to User and send back
// the string representation of the content
func (a *Acheteur) UnmarshalFromRequest(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(a.UnmarshalFromRequest), err.Error())
	}
	err = json.Unmarshal(body, a)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(a.UnmarshalFromRequest), err.Error())
	}
	return nil
}

// Marshal implement ISerial
func (a *Acheteur) Marshal() ([]byte, error) {
	return json.Marshal(a)
}

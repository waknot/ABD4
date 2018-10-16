/*
 * File: reservation.go
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

// Reservation is composed:
type Reservation struct {
	Spectateur 	Spectateur			`json:"Spectateur"`
	Tarif		string				`json:"Tarif"`
}

// ToString return string conversion of marshal user
// absorb error...
func (reserv *Reservation) ToString() string {
	ret, _ := reserv.Marshal()
	return string(ret)
}

// UnmarshalFromRequest take a request object supposed to contain
// a user data object, extract it, convert to User and send back
// the string representation of the content
func (reserv Reservation) UnmarshalFromRequest(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(reserv.UnmarshalFromRequest), err.Error())
	}
	err = json.Unmarshal(body, reserv)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(reserv.UnmarshalFromRequest), err.Error())
	}
	return nil
}

// Marshal implement ISerial
func (reserv *Reservation) Marshal() ([]byte, error) {
	return json.Marshal(reserv)
}

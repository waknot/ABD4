/*
 * File: README.md
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 5:18:37 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package utils

import (
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

// Utils regroup all development utilitaries
type Utils struct{}

var u = &Utils{}

// Use return &Utils{}, struct to wrap utilitaries functions
func Use() *Utils {
	return u
}

// GetFunctionName return the package.name of the passed function
func (u *Utils) GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// GetStack return the name of the function between brackets []
func (u *Utils) GetStack(i interface{}) string {
	return "[" + u.GetFunctionName(i) + "] "
}

//Rand/Generate Ids
var letterRunes = []rune("bcdefghijklmnopquvwxyzABDEFHIJKLNOPQRTUVWXYZ0123456789")

// InitRand launch the random seed for the program
func (u *Utils) InitRand() {
	rand.Seed(time.Now().UnixNano())
}

// RandStringRunes return an aleatory string
func (u *Utils) RandStringRunes(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

/*
 * File: mock.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 8:24:27 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package test

import (
	"ABD4/API/context"
	"ABD4/API/model"
	"log"
	"sort"
	"strconv"
)

// Mock embed model for tests
// Stock it in App global variable in test
// Retrieve and stock all persisted data in
// temporary database.
type Mock struct {
	// exemple:
	Ctx      *context.AppContext
	Users    []string
	PostUser *model.User
}

func (m *Mock) MockUsers(n int) {
	var it int
	manager := m.Ctx.UserManager
	users := []*model.User{}
	for it = 0; it < n; it++ {
		name := "test_" + strconv.Itoa(it)
		email := name + "@etna-alternance.net"
		password := "test"
		u, err := manager.NewUser(name, email, password, "permission", "")
		if err != nil {
			log.Fatalf("mocking users failed: %s", err.Error())
		}
		users = append(users, u)
		// m.Users = append([]string{u.ToString()}, m.Users...)
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].OrderByEmail(users, i, j)
	})
	for _, u := range users {
		m.Users = append(m.Users, u.ToString())
	}
}

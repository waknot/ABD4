/*
 * File: idatamanager.go
 * Project: ABD4/VMD Escape Game
 * File Created: Thursday, 11th October 2018 4:41:33 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 16th October 2018 12:58:45 am
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

import (
	"ABD4/API/model"
)

type IDataManager interface {
	Init(map[string]string) error
	GetDB() interface{}
	SetDB(dbObject interface{}) error
	GetEntity() string
	SetEntity(entity string)
	GetDBName() string
	SetDBName(dbName string)
}

type IUserManager interface {
	IDataManager
	FindAll() ([]*model.User, error)
	FindBy(map[string]string) ([]*model.User, error)
	FindOneBy(map[string]string) (*model.User, error)
	RemoveBy(map[string]string) (int, error)
	RemoveAll() (int, error)
	Create(user *model.User) (*model.User, error)
	// For dev
	FakePersist(n int) error
}

type ITransactionManager interface {
	IDataManager
	FindAll() ([]*model.Transaction, error)
	FindBy(map[string]string) ([]*model.Transaction, error)
	FindOneBy(map[string]string) (*model.Transaction, error)
	RemoveBy(map[string]string) (int, error)
	RemoveAll() (int, error)
	Create(tx *model.Transaction) (*model.Transaction, error)
	// For dev
	FakePersist(n int) error
}

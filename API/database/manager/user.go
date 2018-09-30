/*
 * File: user.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 9:34:29 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package manager

import (
	"ABD4/API/database"
	"ABD4/API/model"
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

type UserManager struct {
	dbm        *database.DBManager
	bucket     string
	userPrefix string
}

func NewUserManager(name, fullpath, secret string) (*UserManager, error) {
	dbm, err := database.NewDBManager(name, fullpath, secret, false, []string{"al-user"})
	if err != nil {
		return nil, fmt.Errorf("<<<< %s %s", utils.Use().GetStack(NewUserManager), err.Error())
	}
	var um *UserManager
	if dbm != nil {
		um = &UserManager{
			dbm:        dbm,
			bucket:     "al-user",
			userPrefix: "a-",
		}
	}

	utils.Use().InitRand()
	return um, nil
}

func (um *UserManager) GetDBM() *database.DBManager {
	return um.dbm
}

func (um *UserManager) NewUser(name string, email string, pass string, permission string, claim string) (*model.User, error) {
	user := &model.User{
		Name:       name,
		Password:   pass,
		Email:      email,
		Permission: permission,
		Claim:      claim,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := um.Save(user)
	if err != nil {
		return nil, fmt.Errorf("<<<< %s %s", utils.Use().GetStack(um.NewUser), err.Error())
	}
	return user, nil
}

func (um *UserManager) Seek() (results []*model.User, err error) {
	if um == nil || um.dbm == nil {
		return
	}
	results = make([]*model.User, 0)

	var byteResults [][]byte

	if byteResults, err = um.dbm.GetByPrefix(um.bucket, um.userPrefix); err != nil {
		return nil, fmt.Errorf("<<<< %s : %s", utils.Use().GetStack(um.Seek), err.Error())
	}

	for _, iter := range byteResults {
		resNew := new(model.User)
		if err = json.Unmarshal([]byte(iter), resNew); err != nil {
			err = fmt.Errorf("<<<< %s : %s", utils.Use().GetStack(um.Seek), err.Error())
			return nil, err
		} else {
			results = append(results, resNew)
		}
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].OrderByEmail(results, i, j)
	})
	return
}

func (um *UserManager) GetByID(id string) (result *model.User, err error) {
	if id == "" {
		return nil, fmt.Errorf("<<<< %s: id is empty", utils.Use().GetStack(um.GetByID))
	}

	var key string
	key = fmt.Sprintf("%s%s", um.userPrefix, id)

	var byt []byte
	byt, err = um.dbm.GetOne(um.bucket, key)

	result = new(model.User)
	if err = json.Unmarshal([]byte(byt), result); err != nil {
		return nil, fmt.Errorf("<<<< %s %s", utils.Use().GetStack(um.GetByID))
	}
	return
}

func (um *UserManager) Save(record *model.User) error {
	if record == nil {
		return fmt.Errorf("<<<< %s record is nil", utils.Use().GetStack(um.Save))
	}
	if record.ID == "" || record.ID == "0" {
		record.ID = utils.Use().RandStringRunes(12)
	}

	record.UpdatedAt = time.Now()

	key := fmt.Sprintf("%s%s", um.userPrefix, record.ID)
	err := um.dbm.Save(um.bucket, key, record)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(um.Save), err.Error())
	}
	return nil
}

func (um *UserManager) Update(record *model.User) error {
	if record == nil || record.ID == "0" {
		return fmt.Errorf("<<<< %s record is nil", utils.Use().GetStack(um.Update))
	}

	var oldRecord *model.User
	var err error

	if oldRecord, err = um.GetByID(record.ID); err != nil {
		return fmt.Errorf("<<<< %s record is nil", utils.Use().GetStack(um.Update))
	}
	//update the time values from the old record
	record.CreatedAt = oldRecord.CreatedAt
	record.UpdatedAt = time.Now()
	err = um.Save(record)
	if err != nil {
		return fmt.Errorf("<<<< %s record is nil", utils.Use().GetStack(um.Update))
	}
	return nil
}

func (um *UserManager) Delete(id string) error {
	if id == "" {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(um.Delete), database.ErrKeyInvalid.Error())
	}
	key := fmt.Sprintf("%s%s", um.userPrefix, id)
	err := um.dbm.Delete(um.bucket, key)
	if err != nil {
		return fmt.Errorf("<<<< %s record is nil", utils.Use().GetStack(um.Update))
	}
	return nil
}

func (um *UserManager) TestSaveUser() {
	var err error

	user, err := um.NewUser("test-name", "test-email", "test-pass", "test-permission", "test-claim")
	if err != nil {
		fmt.Printf("NewUser return err: %s", err)
	}

	if user.ID == "" {
		fmt.Print("NewUser Id is nil")
	}

	retRecord, err := um.GetByID(user.ID)
	if err != nil {
		fmt.Printf("GetByID return err: %s", err)
	} else {
		fmt.Printf("%+v", retRecord)
	}

	if user.Name != retRecord.Name {
		fmt.Print("GetByID returned values are not equal")
	}

	return
}

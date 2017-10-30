package datastore

import (
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store/datastore/sql"
	"github.com/russross/meddler"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/model"
)

func (db *datastore) CreateUser(user *model.User) error {
	return meddler.Insert(db, "users", user)
}

func (db *datastore) GetUser(id int64) (*model.User, error) {
	usr := new(model.User)
	err := meddler.Load(db, "users", usr, id)
	return usr, err
}

func (db *datastore) GetUserByLogin(login string) (*model.User, error) {
	stmt := sql.Lookup(db.driver, "user-find-login")
	data := new(model.User)
	err := meddler.QueryRow(db, data, stmt, login)
	return data, err
}

func (db *datastore) UpdateUser(user *model.User) error {
	return meddler.Update(db, "users", user)
}

func (db *datastore) DeleteUser(user *model.User) error {
	stmt := sql.Lookup(db.driver, "user-delete")
	_, err := db.Exec(stmt, user.ID)
	return err
}

func (db *datastore) GetUserList() ([]*model.User, error) {
	stmt := sql.Lookup(db.driver, "user-find")
	data := []*model.User{}
	err := meddler.QueryAll(db, &data, stmt)
	return data, err
}

func (db *datastore) GetUserCount() (count int, err error) {
	err = db.QueryRow(sql.Lookup(db.driver, "count-users")).Scan(&count)
	return
}
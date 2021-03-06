package datastore

import (
	"github.com/iH8c0ff33/cosmicbox-api-server/store/datastore/sql"
	"github.com/russross/meddler"

	"github.com/iH8c0ff33/cosmicbox-api-server/model"
)

func (db *Datastore) CreateUser(user *model.User) error {
	return meddler.Insert(db, "users", user)
}

func (db *Datastore) GetUser(id int64) (*model.User, error) {
	usr := new(model.User)
	err := meddler.Load(db, "users", usr, id)
	return usr, err
}

func (db *Datastore) GetUserByLogin(login string) (*model.User, error) {
	stmt := sql.Lookup(db.driver, "user-find-login")
	data := new(model.User)
	err := meddler.QueryRow(db, data, stmt, login)
	return data, err
}

func (db *Datastore) UpdateUser(user *model.User) error {
	return meddler.Update(db, "users", user)
}

func (db *Datastore) DeleteUser(user *model.User) error {
	stmt := sql.Lookup(db.driver, "user-delete")
	_, err := db.Exec(stmt, user.ID)
	return err
}

func (db *Datastore) GetAllUsers() ([]*model.User, error) {
	stmt := sql.Lookup(db.driver, "user-find")
	data := []*model.User{}
	err := meddler.QueryAll(db, &data, stmt)
	return data, err
}

func (db *Datastore) GetUsersCount() (count int, err error) {
	err = db.QueryRow(sql.Lookup(db.driver, "count-users")).Scan(&count)
	return
}

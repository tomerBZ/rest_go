package user

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/tomerBZ/web/pkg/interfaces"
)

type userModel struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Username string
	Password string
}

func userModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func newUserModel(u *interfaces.User) *userModel {
	return &userModel{
		Username: u.Username,
		Password: u.Password}
}

func (u *userModel) toRootUser() *interfaces.User {
	return &interfaces.User{
		Id:       u.Id.Hex(),
		Username: u.Username,
		Password: u.Password}
}

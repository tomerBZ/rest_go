package user

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/tomerBZ/web/pkg/crypto"
	"github.com/tomerBZ/web/pkg/interfaces"
	"github.com/tomerBZ/web/pkg/logs"
	"github.com/tomerBZ/web/pkg/mongo"
)

type UserService struct {
	collection *mgo.Collection
}

func NewUserService(session *mongo.Session, dbName string) *UserService {
	collection := session.GetCollection(dbName, "users")
	_ = collection.EnsureIndex(userModelIndex())
	return &UserService{collection}
}

func (userService *UserService) CreateUser(u *interfaces.User) error {
	user := newUserModel(u)

	hashedPassword, err := crypto.Generate(user.Password)
	if err != nil {
		logs.Error.Println(err)
		return err
	}
	user.Password = hashedPassword
	return userService.collection.Insert(&user)
}

func (userService *UserService) UpdateUser(id string, u *interfaces.User) error {
	hashedPassword, err := crypto.Generate(u.Password)
	if err != nil {
		logs.Error.Println(err)
		return err
	}
	u.Password = hashedPassword
	logs.Info.Println(u.Username)
	return userService.collection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, newUserModel(u))
}

func (userService *UserService) GetByUsername(username string) (*interfaces.User, error) {
	model := userModel{}
	err := userService.collection.Find(bson.M{"username": username}).One(&model)
	return model.toRootUser(), err
}

func (userService *UserService) GetUsers() ([]*interfaces.User, error) {
	var users []*interfaces.User
	err := userService.collection.Find(nil).All(&users)
	return users, err
}

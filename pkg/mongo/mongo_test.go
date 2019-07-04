package mongo

import (
	"github.com/tomerBZ/web/pkg/interfaces"
	"github.com/tomerBZ/web/pkg/mock"
	user2 "github.com/tomerBZ/web/pkg/mongo/user"
	"log"
	"testing"
)

const (
	dbName             = "test_db"
	userCollectionName = "user"
)

func Test_UserService(t *testing.T) {
	t.Run("CreateUser", createuserShouldInsertUserIntoMongo)
}

func createuserShouldInsertUserIntoMongo(t *testing.T) {
	//Arrange
	session, err := NewSession()
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer session.Close()
	mockHash := mock.Hash{}
	userService := user2.NewUserService(session.Copy(), dbName, userCollectionName, &mockHash)

	testUsername := "integration_test_user"
	testPassword := "integration_test_password"
	user := interfaces.User{
		Username: testUsername,
		Password: testPassword}

	//Act
	err = userService.CreateUser(&user)

	//Assert
	if err != nil {
		t.Error("Unable to create user: %s", err)
	}
	var results []interfaces.User
	_ = session.GetCollection(dbName, userCollectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Error("Incorrect number of results. Expected `1`, got: `%i`", count)
	}
	if results[0].Username != user.Username {
		t.Error("Incorrect Username. Expected `%s`, Got: `%s`", testUsername, results[0].Username)
	}
}

package mongo

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/joho/godotenv"
	"github.com/tomerBZ/web/pkg/logs"
	"os"
	"time"
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		logs.Error.Fatal("Error loading .env file")
	}
	db := os.Getenv("MONGO_DB")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")

	logs.Info.Println("connecting to mongo at:", host+":"+port)

	logs.Info.Println("enter main - connecting to mongo")

	// tried doing this - doesn't work as intended
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Detected panic")
			var ok bool
			err, ok := r.(error)
			if !ok {
				logs.Info.Printf("pkg:  %v,  error: %s", r, err)
			}
		}
	}()

	maxWait := time.Duration(5 * time.Second)
	session, sessionErr := mgo.DialWithTimeout("localhost:27017", maxWait)
	if sessionErr == nil {
		session.SetMode(mgo.Monotonic, true)
		database := session.DB(db)
		if database == nil {
			return
		}
		logs.Info.Println("Got a collection object")
		_, err := database.CollectionNames()
		if err != nil {
			logs.Error.Printf("No collections in database %v", database)
			return
		}
	} else { // never gets here
		logs.Error.Fatal("Unable to connect to local mongo instance!")
	}
}

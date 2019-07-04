package app

import (
	"github.com/joho/godotenv"
	"github.com/tomerBZ/web/pkg/logs"
	"github.com/tomerBZ/web/pkg/mongo"
	"github.com/tomerBZ/web/pkg/routes"
	"github.com/tomerBZ/web/pkg/utils"
	"net/http"
)

type App struct {
	session *mongo.Session
}

func (a *App) Run() {
	//init logs
	logs.Init()
	// init env file
	err := godotenv.Load()
	if err != nil {
		logs.Error.Fatal("Error loading .env file")
	}
	db := utils.GetEnv("MONGO_DB", "users")
	host := utils.GetEnv("APP_HOST", "localhost")
	port := utils.GetEnv("APP_PORT", "3000")
	// Init routes
	logs.Info.Println("Server is listening on:", port)
	// Connect to mongo db
	//mongo.Connect()
	a.session, err = mongo.NewSession()
	if err != nil {
		logs.Error.Fatal("unable to connect to mongodb")
	}
	// Serve
	router := routes.Router(a.session.Copy(), db)
	err = http.ListenAndServe(host+":"+port, router)
	if err != nil {
		logs.Error.Fatal(err)
		return
	}
}

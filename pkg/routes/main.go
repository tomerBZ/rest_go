package routes

import (
	"github.com/go-chi/chi"
	"github.com/tomerBZ/web/pkg/mongo"
	"github.com/tomerBZ/web/pkg/mongo/user"
)

func Router(a *mongo.Session, dbName string) chi.Router {
	u := user.NewUserService(a.Copy(), dbName)
	router := chi.NewRouter()
	router.Mount("/books", booksRouter())
	router.Mount("/users", userRouting(u))

	return router
}

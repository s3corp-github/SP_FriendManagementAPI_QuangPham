package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/quangpham789/golang-assessment/handler/relation"
	"github.com/quangpham789/golang-assessment/handler/user"
	"github.com/quangpham789/golang-assessment/utils"
	"github.com/quangpham789/golang-assessment/utils/db"
	"log"
	"net/http"
)

func main() {
	//load config
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// Connection database
	dbConn, err := db.ConnectDB(config.DBDriver, config.DBSource)
	if err != nil {
		log.Println("Connection database fail with error: ", err)
	}

	// Init router with gochi
	router := initRouter(dbConn)

	// Start server with port 8080
	log.Println("Server start at port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println("Error start server with port 8080", err)
	}
}

func initRouter(dbConn *sql.DB) *chi.Mux {
	// init user handler
	userHandler := user.NewUserHandler(dbConn)
	relationHandler := relation.NewRelationsHandler(dbConn)
	router := chi.NewRouter()
	router.Use(logRequest)

	// create: main router
	router.Route("/", func(r chi.Router) {
		r.Mount("/users", userRouter(userHandler))
		r.Mount("/relations", relationRouter(relationHandler))
	})
	return router
}

// TODO: Implement getUsers() API
func userRouter(userHandler user.UserHandler) http.Handler {
	router := chi.NewRouter()
	router.Post("/", userHandler.CreateUser)
	return router
}

func relationRouter(relationHandler relation.RelationsHandler) http.Handler {
	router := chi.NewRouter()
	router.Post("/createfriendrelation", relationHandler.CreateFriendsRelation)
	router.Post("/createsubscriptionrelation", relationHandler.CreateSubscriptionRelation)
	router.Post("/createblockrelation", relationHandler.CreateBlockRelation)
	router.Post("/getfriendfromemail", relationHandler.GetAllFriendOfUser)
	router.Post("/getcommonfriends", relationHandler.GetCommonFriend)
	router.Post("/getemailreceive", relationHandler.GetEmailReceive)

	return router
}

// Log middleware
func logRequest(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request path: %s", r.URL.Path)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/handler/rest"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/config"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/db"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbConn, err := db.ConnectDB(config.DBSource)
	if err != nil {
		log.Println("Connection database fail with error: ", err)
	}

	router := initRouter(dbConn)

	// Start server with port 8080
	log.Println("Server start at port 8080")
	if err := http.ListenAndServe(config.ServerAddress, router); err != nil {
		log.Println("Error start server with port 8080", err)
	}
}

func initRouter(dbConn *sql.DB) *chi.Mux {
	userHandler := rest.NewUserHandler(dbConn)
	relationHandler := rest.NewRelationsHandler(dbConn)
	router := chi.NewRouter()
	router.Use(utils.LogRequest)

	router.Route("/", func(r chi.Router) {
		r.Mount("/users", userRouter(userHandler))
		r.Mount("/relations", relationRouter(relationHandler))
	})

	return router
}

func userRouter(userHandler rest.UsersHandler) http.Handler {
	router := chi.NewRouter()
	router.Post("/", userHandler.CreateUser)
	router.Get("/", userHandler.GetListUser)
	return router
}

func relationRouter(relationHandler rest.RelationsHandler) http.Handler {
	router := chi.NewRouter()
	router.Post("/friend", relationHandler.CreateFriendsRelation)
	router.Post("/subscription", relationHandler.CreateSubscriptionRelation)
	router.Post("/block", relationHandler.CreateBlockRelation)
	router.Post("/friends", relationHandler.GetAllFriendOfUser)
	router.Post("/commonfriends", relationHandler.GetCommonFriend)
	router.Post("/emailreceive", relationHandler.GetEmailReceive)

	return router
}

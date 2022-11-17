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
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println("Error start server with port 8080", err)
	}
}

func initRouter(dbConn *sql.DB) *chi.Mux {
	handler := rest.NewHandler(dbConn)
	router := chi.NewRouter()
	router.Use(utils.LogRequest)

	router.Route("/", func(r chi.Router) {
		r.Mount("/users", usersRouter(handler))
		r.Mount("/friends", friendsRouter(handler))
	})

	return router
}

func usersRouter(userHandler rest.Handler) http.Handler {
	router := chi.NewRouter()
	router.Post("/", userHandler.CreateUser)
	router.Get("/", userHandler.GetListUser)
	return router
}

func friendsRouter(friendsHandler rest.Handler) http.Handler {
	router := chi.NewRouter()
	router.Post("/friend", friendsHandler.CreateFriends)
	router.Post("/subscription", friendsHandler.CreateSubscription)
	router.Post("/block", friendsHandler.CreateBlock)
	router.Post("/friends", friendsHandler.GetFriends)
	router.Post("/commonfriends", friendsHandler.GetCommonFriends)
	router.Post("/emailreceive", friendsHandler.GetEmailReceive)

	return router
}

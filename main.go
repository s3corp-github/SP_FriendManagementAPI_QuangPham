package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/quangpham789/golang-assessment/handler"
	"github.com/quangpham789/golang-assessment/utils/db"
	"log"
	"net/http"
)

func main() {
	// Connection database
	dbURL := "postgresql://root:secret@localhost:5432/friends_management?sslmode=disable"
	dbConn, err := db.ConnectDB(dbURL)
	if err != nil {
		log.Println("Connection database fail with error: ", err)
	}

	// Init router with gochi
	router := initRouter(dbConn)

	// Start server with port 5000
	log.Println("Server start at port 3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Println("Error start server with port 3000", err)
	}
}

func initRouter(dbConn *sql.DB) *chi.Mux {
	// init user handler
	userHandler := handler.NewUserHandler(dbConn)
	relationHandler := handler.NewRelationsHandler(dbConn)
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
func userRouter(userHandler handler.UserHandler) http.Handler {
	router := chi.NewRouter()
	router.Post("/", userHandler.CreateUser)
	return router
}

func relationRouter(relationHandler handler.RelationsHandler) http.Handler {
	router := chi.NewRouter()
	router.Post("/getfriendfromemail", relationHandler.GetAllFriendOfUser)
	router.Post("/createrelation", relationHandler.CreateRelation)
	router.Post("/getcommonfriends", relationHandler.GetCommonFriend)

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

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"kanca-studio/palang/database"
	"kanca-studio/palang/delivery/http/routing"
	"kanca-studio/palang/manager"
	"kanca-studio/palang/service/auth"
	"kanca-studio/palang/service/user"
	"net/http"
	"time"
)

var userManager manager.UserManager

func init() {
	database.Init("localhost", 5432, "postgres", "postgres", "palang", true)

	//auto migrate
	database.GetInstance().AutoMigrate(user.Model{})
	var (
		userRepo = user.NewRepository(database.GetInstance())
	)

	var (
		userService = user.NewService(userRepo)
		authService = auth.NewService()
	)
	userManager = manager.NewUserManager(userService, authService)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET", "POST")
	routing.UserRouter(router, userManager)

	http.ListenAndServe(":8000", router)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"title": "palang",
		"time":  time.Now().UnixNano() / 1000000,
	})
}

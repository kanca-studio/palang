package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"kanca-studio/palang/database"
	"kanca-studio/palang/manager"
	"kanca-studio/palang/service/auth"
	"kanca-studio/palang/service/user"
	"net/http"
	"time"
)

var userManager manager.UserManager

func init() {
	database.Init("localhost", 5432, "posgtres", "posgtres", "kanca", true)

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
	router.HandleFunc("/", Index).Methods("GET")
	http.ListenAndServe(":8000", router)
}


func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"title": "palang",
		"time":  time.Now().UnixNano(),
	})
}

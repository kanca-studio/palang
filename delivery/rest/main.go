package main

import (
	"encoding/json"
	"github.com/gorilla/context"
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
	router.HandleFunc("/", Index).Methods("GET", "POST")
	router.HandleFunc("/register", Register).Methods("POST")
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/user/me", checkAuth(Me)).Methods("GET", "POST")
	router.HandleFunc("/validate", ValidateToken).Methods("GET", "POST")
	http.ListenAndServe(":8000", router)
}

func checkAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if err := userManager.ValidateToken(token); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
		context.Set(r, "token", token)
		next(w, r)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"title": "palang",
		"time":  time.Now().UnixNano() / 1000000,
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"title": "palang",
		"time":  time.Now().UnixNano() / 1000000,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"title": "palang",
		"time":  time.Now().UnixNano() / 1000000,
	})
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"title": "palang",
		"time":  time.Now().UnixNano() / 1000000,
	})
}

func Me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"title": "palang",
		"time":  time.Now().UnixNano() / 1000000,
	})
}

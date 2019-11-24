package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/gorilla/mux"
	"github.com/kanca-studio/palang/database"
	"github.com/kanca-studio/palang/delivery/http/routing"
	"github.com/kanca-studio/palang/manager"
	"github.com/kanca-studio/palang/service/auth"
	"github.com/kanca-studio/palang/service/user"

	"github.com/bukalapak/ottoman/x/env"
	"github.com/subosito/gotenv"
)

var userManager manager.User

func init() {
	if err := gotenv.Load(); err != nil {
		panic(err)
	}
	database.Init(
		env.String("DATABASE_HOST"),
		env.Int("DATABASE_PORT"),
		env.String("DATABASE_USERNAME"),
		env.String("DATABASE_PASSWORD"),
		env.String("DATABASE_NAME"),
		true,
	)

	//auto migrate
	database.GetInstance().AutoMigrate(user.Model{})
	var (
		userRepo = user.NewRepository(user.WithGorm(database.GetInstance()))
	)

	var (
		userService = user.NewService(userRepo)
		authService = auth.NewService()
	)
	userManager = manager.NewUser(userService, authService)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET", "POST")
	routing.UserRouter(router, userManager)

	fmt.Println("Running on 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"title": "palang",
		"time":  time.Now().UnixNano() / 1000000,
	})
}

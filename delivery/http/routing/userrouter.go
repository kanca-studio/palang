package routing

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"kanca-studio/palang/delivery/http/schema"
	"kanca-studio/palang/manager"
	"net/http"
	"time"
)

var validate *validator.Validate
var userManager manager.UserManager

func UserRouter(router *mux.Router, _userManager manager.UserManager) {
	validate = validator.New()
	userManager = _userManager

	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	//router.HandleFunc("/user/me", checkAuth(Me)).Methods("GET", "POST")
	router.HandleFunc("/validate", validateToken).Methods("GET", "POST")
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var body schema.ReqRegister
	if err := decoder.Decode(&body); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	if err := validate.Struct(&body); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	if err := userManager.Register(body.IdentifierType, body.Identifier, body.Password); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "success register",
	})
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var body schema.ReqRegister
	if err := decoder.Decode(&body); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	if err := validate.Struct(&body); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	token, err := userManager.Login(body.IdentifierType, body.Identifier, body.Password);
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "success login",
		"token":   token,
	})
}

func validateToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"title": "palang",
		"time":  time.Now().UnixNano() / 1000000,
	})
}

func me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"title": "palang",
		"time":  time.Now().UnixNano() / 1000000,
	})
}

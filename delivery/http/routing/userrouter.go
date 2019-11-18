package routing

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/kanca-studio/palang/delivery/http/schema"
	"github.com/kanca-studio/palang/manager"
)

var validate *validator.Validate
var userManager manager.User

func UserRouter(router *mux.Router, _userManager manager.User) {
	validate = validator.New()
	userManager = _userManager

	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/user/me", checkAuth(me)).Methods("GET", "POST")
	router.HandleFunc("/validate-token", validateToken).Methods("GET", "POST")
}

func checkAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if err := userManager.ValidateToken(token); err != nil {
			respondWithError(w, http.StatusUnauthorized, "Restricted Page", errors.New("Unauthorized"))
			return
		}
		context.Set(r, "token", token)
		next(w, r)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var body schema.ReqRegister
	if err := decoder.Decode(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "parsing body failed", err)
		return
	}

	if err := validate.Struct(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "validation body failed", err)
		return
	}

	if err := userManager.Register(body.IdentifierType, body.Identifier, body.Password); err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed Register ", err)
		return
	}

	var response schema.ResponseRegister
	response.Message = "success Register"

	respondWithSuccess(w, response)
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var body schema.ReqRegister
	if err := decoder.Decode(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "parsing body failed", err)
		return
	}

	if err := validate.Struct(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "validation body failed", err)
		return
	}

	token, err := userManager.Login(body.IdentifierType, body.Identifier, body.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed Login ", err)
		return
	}
	var response schema.ResponseLogin
	response.Message = "success login"
	response.Token = token
	respondWithSuccess(w, response)
}

func validateToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if err := userManager.ValidateToken(token); err != nil {
		json.NewEncoder(w).Encode(map[string]bool{
			"valid": false,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{
		"valid": true,
	})
}

func me(w http.ResponseWriter, r *http.Request) {
	token := context.Get(r, "token")
	data, err := userManager.Me(token.(string))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed Get Profile ", err)
		return
	}

	var response schema.ResponseMe
	response.Message = "success get profile"
	response.Data.Name = data.Name
	response.Data.Email = data.Email
	response.Data.Username = data.Username
	response.Data.PhoneNumber = data.PhoneNumber
	respondWithSuccess(w, response)

}

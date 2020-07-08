package session

import (
	"encoding/json"
	"net/http"

	// . "api/middleware"

	. "github.com/booscaaa/jwt-auth"
)

//SetOrigins .
func SetOrigins(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE, PUT")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	return w
}

//Create .
func Create(w http.ResponseWriter, r *http.Request) {
	w = SetOrigins(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		var access Access
		if err := json.NewDecoder(r.Body).Decode(&access); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
		} else {
			defer r.Body.Close()
			SessionCreate(access, w)
		}
	}
}

//Refresh .
func Refresh(w http.ResponseWriter, r *http.Request) {
	w = SetOrigins(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		bearToken := r.Header.Get("Authorization")
		SessionRefresh(bearToken, w)

		// if isAuth, _ := VerifyToken(w, r); isAuth {
		// var colaborador Colaborador
		// if err := json.NewDecoder(r.Body).Decode(&colaborador); err != nil {
		// 	log.Println("Nenhum parametro enviado")
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte("500 - Something bad happened!"))
		// } else {
		// 	defer r.Body.Close()
		// 	response := colaborador.Get()
		// 	w.Write(response)
		// }

		// w.Write([]byte("asdasd"))
		// }
	}
}

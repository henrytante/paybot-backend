package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"zapys-api/auth"
	"zapys-api/utils/db"
	"zapys-api/utils/send"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		fmt.Println(err)
		send.ERRO(w, http.StatusInternalServerError, err)
		return
	}
	defer r.Body.Close()

	if credentials.Username == "" || credentials.Password == "" {
		send.ERRO(w, http.StatusBadRequest, errors.New("As credenciais não podem estar em branco!"))
		return
	}

	userID, err := db.LoginUser(credentials.Username, credentials.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			send.ERRO(w, http.StatusUnauthorized, errors.New("Credenciais inválidas"))
			return
		}
		send.ERRO(w, http.StatusUnauthorized, err)
		return
	}
	token, err := auth.CriarToken(userID) // Convertendo para int, se necessário
	if err != nil {
		send.ERRO(w, http.StatusInternalServerError, err)
		return
	}

	response := map[string]string{"token": token}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		send.ERRO(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResponse)
}

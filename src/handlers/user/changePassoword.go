package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"zapys-api/auth"
	"zapys-api/utils/db"
	"zapys-api/utils/send"
)

type NewPassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	ID, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		send.ERRO(w, http.StatusUnauthorized, errors.New("Erro ao extrair id"))
		return
	}

	var data NewPassword
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		send.ERRO(w, http.StatusInternalServerError, err)
		return
	}
	if data.OldPassword == "" || data.NewPassword == "" {
		send.ERRO(w, http.StatusBadRequest, errors.New("Senha em branco"))
		return
	}

	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	var pass string
	err = database.QueryRow("SELECT password FROM users WHERE id = ?", ID).Scan(&pass)
	if err != nil {
		if err == sql.ErrNoRows {
			send.ERRO(w, http.StatusNotFound, errors.New("Usuário não encontrado"))
			return
		}
		send.ERRO(w, http.StatusInternalServerError, err)
		return
	}

	if data.OldPassword != pass {
		send.ERRO(w, http.StatusUnauthorized, errors.New("Senha antiga está incorreta"))
		return
	}

	_, err = database.Exec("UPDATE users SET password = ? WHERE id = ?", data.NewPassword, ID)
	if err != nil {
		send.ERRO(w, http.StatusInternalServerError, err)
		return
	}
}

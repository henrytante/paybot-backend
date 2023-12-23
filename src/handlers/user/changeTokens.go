package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"zapys-api/auth"
	"zapys-api/utils/db"
	"zapys-api/utils/send"
)

type Data struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}

func ChangeTokens(w http.ResponseWriter, r *http.Request) {
	id, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		send.ERRO(w, http.StatusUnauthorized, errors.New("erro no id"))
	}
	var data Data
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		send.ERRO(w, http.StatusInternalServerError, err)
		return
	}
	defer r.Body.Close()
	if data.Token == "" || data.Type == "" {
		send.ERRO(w, http.StatusBadRequest, errors.New("Os dados não podem estar em branco"))
		return
	}
	err = auth.ValidarToken(r)
	if err != nil {
		send.ERRO(w, http.StatusUnauthorized, errors.New("Erro ao validar token"))
		return
	}
	err = db.UpdateToken(data.Token, id, data.Type)
	if err != nil {
		send.ERRO(w, http.StatusInternalServerError, err)
		return
	}
	send.JSON(w, http.StatusOK, "Token alterado")
}

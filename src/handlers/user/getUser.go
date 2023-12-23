package user

import (
	"net/http"
	"zapys-api/auth"
	"zapys-api/utils/db"
	"zapys-api/utils/send"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	userIDToken, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		send.ERRO(w, http.StatusUnauthorized, err)
		return
	}

	getUser, err := db.GetUser(int(userIDToken))
	if err != nil {
		send.ERRO(w, http.StatusInternalServerError, err)
		return
	}
	send.JSON(w, http.StatusOK, getUser)

}

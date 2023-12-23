package user

import (
	"net/http"
	"zapys-api/auth"
	"zapys-api/utils/send"
)


func VerifyToken(w http.ResponseWriter, r *http.Request)  {
	if err := auth.ValidarToken(r); err != nil{
		send.ERRO(w, http.StatusUnauthorized, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
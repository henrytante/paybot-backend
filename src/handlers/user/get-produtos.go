package user

import (
	"database/sql"
	"errors"
	"net/http"
	"zapys-api/auth"
	"zapys-api/utils/db"
	"zapys-api/utils/send"
)



func GetProdutosByID(w http.ResponseWriter, r *http.Request)  {
	id, err := auth.ExtrairUsuarioID(r)
	if err != nil{
		send.ERRO(w, http.StatusUnprocessableEntity, err)
		return
	}
	produtos, err := db.GetProdutosByID(int(id))
	if err != nil{
		if err == sql.ErrNoRows{
			send.ERRO(w, http.StatusNotFound, errors.New("Este usuario não tem produtos"))
			return
		}
		send.ERRO(w, http.StatusInternalServerError, err)
		return
	}
	send.JSON(w, http.StatusOK, produtos)
}
package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"zapys-api/auth"
	"zapys-api/utils/db"
	"zapys-api/utils/send"
)

type jsonData struct {
	ProdutoID int `json:"pid"`
}

func DeletarProduto(w http.ResponseWriter, r *http.Request)  {
	jwtID, err := auth.ExtrairUsuarioID(r)
	if err != nil{
		send.ERRO(w, http.StatusUnauthorized, err)
		return
	}
	var jsonData jsonData
	err = json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil{
		send.ERRO(w, http.StatusBadRequest, err)
		return
	}
	if err = db.DeletarProduto(int(jwtID), jsonData.ProdutoID); err != nil{
		if err == sql.ErrNoRows{
			send.ERRO(w, http.StatusNotFound, errors.New("Produto não encontrado"))
			return
		}
		send.ERRO(w, http.StatusInternalServerError, err)
		return
	}
	send.JSON(w, http.StatusOK, "Produto deletado")
}
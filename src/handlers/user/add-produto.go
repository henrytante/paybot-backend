package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"zapys-api/auth"
	"zapys-api/utils/db"
	"zapys-api/utils/send"
)

type Produto struct{
	Name string `json:"name"`
	Content string `json:"content"`
	Price int `json:"price"`
}

func AddProduto(w http.ResponseWriter, r *http.Request)  {
	owner, err := auth.ExtrairUsuarioID(r)
	if err != nil{
		send.ERRO(w, http.StatusUnauthorized, errors.New("Erro ao extrair id"))
		return
	}
	var produto Produto
	err = json.NewDecoder(r.Body).Decode(&produto)
	if err != nil{
		send.ERRO(w, http.StatusUnprocessableEntity, errors.New("Erro ao ler json"))
		return
	}
	defer r.Body.Close()
	if produto.Name == "" || produto.Content == "" || produto.Price == 0{
		send.ERRO(w, http.StatusBadRequest, errors.New("Dados em branco"))
		return
	}
	if err := db.AddProduto(produto.Name, produto.Content, produto.Price, int(owner)); err != nil{
		send.ERRO(w, http.StatusInternalServerError, err)
		return
	}
	send.JSON(w, http.StatusOK, "Produto adicionado")
}
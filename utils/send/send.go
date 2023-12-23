package send

import (
	"encoding/json"
	"log"
	"net/http"
)


func JSON(w http.ResponseWriter, status int, dados interface{})  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if dados != nil{
		if err := json.NewEncoder(w).Encode(dados); err != nil{
			log.Fatal(err)
		}
	}
}

func ERRO(w http.ResponseWriter,  status int, err error)  {
	JSON(w, status, struct {
		Erro string `json:"erro"`
	}{
		Erro: err.Error(),
	})
} 
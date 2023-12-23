package middleware

import (
	"net/http"
	"zapys-api/auth"
	"zapys-api/utils/send"
)

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidarToken(r); err != nil {
			send.ERRO(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}

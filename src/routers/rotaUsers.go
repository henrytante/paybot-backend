package routers

import (
	"net/http"
	
	"zapys-api/src/handlers/user"
)

var rotasUsuarios = []Rota{
	{
		URI:    "/login",
		Metodo: http.MethodPost,
		Funcao: user.LoginUser,
		Auth:   false,
	},
	{
		URI:    "/change-token",
		Metodo: http.MethodPost,
		Funcao: user.ChangeTokens,
		Auth:   true,
	},
	{
		URI:    "/user",
		Metodo: http.MethodGet,
		Funcao: user.GetUser,
		Auth:   true,
	},
	{
		URI: "/change-password",
		Metodo: http.MethodPost,
		Funcao: user.ChangePassword,
		Auth: true,
	},
	{
		URI: "/verify",
		Metodo: http.MethodGet,
		Funcao: user.VerifyToken,
		Auth: false,
	},
	{
		URI: "/message",
		Metodo: http.MethodGet,
		Funcao: user.SendMessage,
		Auth: true,
	},
	{
		URI: "/produtos",
		Metodo: http.MethodPost,
		Funcao: user.AddProduto,
		Auth: true,
	},
	{
		URI: "/produtos",
		Metodo: http.MethodGet,
		Funcao: user.GetProdutosByID,
		Auth: true,
	},
	{
		URI: "/add-produto",
		Metodo: http.MethodPost,
		Funcao: user.AddProduto,
		Auth: true,
	},
	{
		URI: "/deletar-produto",
		Metodo: http.MethodPost,
		Funcao: user.DeletarProduto,
		Auth: true,
	},
}

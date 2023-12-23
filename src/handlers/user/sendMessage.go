package user

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"
	"zapys-api/auth"
	"zapys-api/utils/db"
	"zapys-api/utils/send"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	id, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		send.ERRO(w, http.StatusUnauthorized, errors.New("Erro ao coletar seu id"))
		return
	}
	message := r.URL.Query().Get("msg")
	if message == "" {
		send.ERRO(w, http.StatusBadRequest, errors.New("'msg' em branco"))
		return
	}
	message = strings.Replace(message, "\\n", "\n", -1) // Substitui "\\n" por "\n" para quebras de linha
	token, err := db.GetToken(id)
	if err != nil {
		if err == sql.ErrNoRows {
			send.ERRO(w, http.StatusNotFound, errors.New("Token não encontrado"))
			return
		}
		send.ERRO(w, http.StatusInternalServerError, err)
		return
	}
	if token == "empty" {
		send.ERRO(w, http.StatusServiceUnavailable, errors.New("Token não configurado"))
		return
	}
	chatid, err := db.GetChatID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			send.ERRO(w, http.StatusNotFound, errors.New("ChatID não encontrado"))
			return
		} else {
			send.ERRO(w, http.StatusInternalServerError, err)
			return
		}
	}
	if chatid == 0 {
		send.ERRO(w, http.StatusServiceUnavailable, errors.New("ChatID não configurado"))
		return
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		send.ERRO(w, http.StatusInternalServerError, errors.New("Erro ao iniciar o bot"))
		return
	}
	msg := tgbotapi.NewMessage(int64(chatid), message) // Use a variável 'message' diretamente
	msg.ParseMode = "Markdown"
	_, err = bot.Send(msg)
	if err != nil {
		send.ERRO(w, http.StatusInternalServerError, errors.New("Erro ao enviar mensagem"))
		return
	}
	send.JSON(w, http.StatusOK, err)
}

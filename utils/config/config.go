package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectString = ""
	Porta         = 0
	JWTKEY        = ""
)

func Carregar() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Porta, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Porta = 8080
	}
	JWTKEY = os.Getenv("JWTKEY")
	ConnectString = fmt.Sprintf("freedb_h1000admin:5xeYDt#@3KbqBum@tcp(sql.freedb.tech:3306)/freedb_paybotDB")
}

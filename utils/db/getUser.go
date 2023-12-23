package db

type User struct{
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	ChatID int `json:"chat_id"`
	Vendas int `json:"vendas"`
	ValorTotalVendas float32 `json:"valor_total_vendas"`
	TokenTelegram string `json:"token_telegram"`
	TokenBanco string `json:"token_banco"`
}

func GetUser(id int) (User, error) {
	db, err := ConnectDB()
	if err != nil{
		return User{}, err
	}
	defer db.Close()
	getUser := db.QueryRow("SELECT id, username, password, chatid, vendas, valor_total_vendas, COALESCE(token_telegram, 'Não configurado'), COALESCE(token_banco, 'Não configurado') from users WHERE id = ?", id)
	if err != nil{
		return User{}, err
	}
	var user User
	if err = getUser.Scan(&user.ID, &user.Username, &user.Password, &user.ChatID, &user.Vendas, &user.ValorTotalVendas, &user.TokenTelegram, &user.TokenBanco); err != nil{
		return user, err
	}
	return user, nil
}
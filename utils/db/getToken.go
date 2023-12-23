package db



func GetToken(id uint64) (string, error) {
	db, _ := ConnectDB()
	defer db.Close()
	var token string
	if err := db.QueryRow("SELECT COALESCE(token_telegram, 'empty') FROM users WHERE id = ?", id).Scan(&token); err != nil{
		return "", err
	}
	return token, nil
}
package db


func GetChatID(id uint64) (int, error) {
	db, _ := ConnectDB()
	defer db.Close()
	var chat_id int
	if err := db.QueryRow("SELECT chatid FROM users WHERE id = ?", id).Scan(&chat_id); err != nil{
		return 0, err
	}
	return chat_id, nil
}
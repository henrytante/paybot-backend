package db

import "log"

func LoginUser(username, password string) (uint64, error) {
	var id uint64
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?", username, password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

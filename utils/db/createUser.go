package db

func CreateUser(username, password string) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}

	defer db.Close()
	if _, err := db.Exec("INSERT INTO users (username, password) VALUES (?,?)", username, password); err != nil {
		return err
	}
	return nil
}

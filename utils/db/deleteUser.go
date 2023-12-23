package db

import (
	"errors"
)

func userExists(id int) (bool, error) {
	db, err := ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func DeleteUser(id int) error {
	exists, err := userExists(id)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Usuário não encontrado")
	}

	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec("DELETE FROM users WHERE id = ?", id); err != nil {
		return err
	}

	return nil
}

package db

import "errors"

func UpdateToken(token string, id uint64, op string) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	var query string
	switch op {
	case "telegram":
		query = "UPDATE users SET token_telegram = ? WHERE id = ?"
	case "banco":
		query = "UPDATE users SET token_banco = ? WHERE id = ?"
	default:
		return errors.New("Operação não suportada")
	}
	_, err = db.Exec(query, token, id)
	if err != nil {
		return err
	}
	return nil
}

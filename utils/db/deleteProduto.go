package db

func DeletarProduto(ownerID, pID int) error {
	db, err := ConnectDB()
	if err != nil{
		return err
	}
	defer db.Close()
	if _, err = db.Exec("DELETE FROM produtos WHERE id = ? AND owner = ?", pID, ownerID); err != nil{
		return err
	}
	return nil
}
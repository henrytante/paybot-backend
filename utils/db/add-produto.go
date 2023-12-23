package db

func AddProduto(name, content string, price, owner int) error {
	db, _ := ConnectDB()
	defer db.Close()
	if _, err := db.Exec("INSERT INTO produtos (name, content , price, owner) VALUES (?,?,?,?)", name, content, price, owner); err != nil{
		return err
	}
	return nil
}
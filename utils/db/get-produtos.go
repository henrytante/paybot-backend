package db

type Produto struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Price   int    `json:"price"`
}

func GetProdutosByID(ownerID int) ([]Produto, error) {
	db, err := ConnectDB()
	if err != nil{
		return []Produto{}, err
	}
	defer db.Close()
	getProdutos, err := db.Query("SELECT id,name,content,price FROM produtos WHERE owner = ?", ownerID)
	if err != nil{
		return []Produto{}, err
	}
	var produtos []Produto
	for getProdutos.Next(){
		var produto Produto
		err = getProdutos.Scan(&produto.ID, &produto.Name, &produto.Content, &produto.Price)
		if err != nil{
			return []Produto{}, err
		}
		produtos = append(produtos, produto)
	}
	return produtos, nil
}

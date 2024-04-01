package data

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Color    string `json:"color"`
	Size     string `json:"size"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
}

var products []Product = make([]Product, 0)

func GetProduct(id int) *Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

func GetProducts() []Product {
	return products
}

func CreateProduct(product Product) {
	products = append(products, product)
}

func UpdateProduct(product Product) {
	for i, p := range products {
		if p.ID == product.ID {
			products[i] = product
			return
		}
	}
}

func DeleteProduct(id int) {
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			return
		}
	}
}

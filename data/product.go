package data

type Product struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Price    int            `json:"price"`
	Quantity int            `json:"quantity"`
	Color    string         `json:"color"`
	Size     string         `json:"size"`
	Images   []ProductImage `json:"images"`
	Brand    string         `json:"brand"`
	Model    string         `json:"model"`
}

type ProductImage struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func GetProduct() *Product {
	return &Product{
		ID:       1,
		Name:     "Product 1",
		Price:    10000,
		Quantity: 10,
		Color:    "red",
		Size:     "large",
		Images: []ProductImage{
			{
				ID:  1,
				URL: "https://example.com/image1.jpg",
			},
			{
				ID:  2,
				URL: "https://example.com/image2.jpg",
			},
		},
		Brand: "brand",
		Model: "12345670",
	}
}

func GetProducts() []Product {
	return []Product{
		{
			ID:       1,
			Name:     "Product 1",
			Price:    10000,
			Quantity: 10,
			Color:    "red",
			Size:     "large",
			Images: []ProductImage{
				{
					ID:  1,
					URL: "https://example.com/image1.jpg",
				},
				{
					ID:  2,
					URL: "https://example.com/image2.jpg",
				},
			},
			Brand: "brand",
			Model: "1234567",
		},
		{
			ID:       1,
			Name:     "Product 1",
			Price:    10000,
			Quantity: 10,
			Color:    "red",
			Size:     "large",
			Images: []ProductImage{
				{
					ID:  1,
					URL: "https://example.com/image1.jpg",
				},
				{
					ID:  2,
					URL: "https://example.com/image2.jpg",
				},
			},
			Brand: "brand",
			Model: "12345678",
		},
		{
			ID:       1,
			Name:     "Product 1",
			Price:    10000,
			Quantity: 10,
			Color:    "red",
			Size:     "large",
			Images: []ProductImage{
				{
					ID:  1,
					URL: "https://example.com/image1.jpg",
				},
				{
					ID:  2,
					URL: "https://example.com/image2.jpg",
				},
			},
			Brand: "brand",
			Model: "123456780",
		},
		{
			ID:       1,
			Name:     "Product 1",
			Price:    10000,
			Quantity: 10,
			Color:    "red",
			Size:     "large",
			Images: []ProductImage{
				{
					ID:  1,
					URL: "https://example.com/image1.jpg",
				},
				{
					ID:  2,
					URL: "https://example.com/image2.jpg",
				},
			},
			Brand: "brand",
			Model: "12345670",
		},
	}
}

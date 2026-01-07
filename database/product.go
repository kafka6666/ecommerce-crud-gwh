package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

var ProductList []*Product

func init() {
	prd1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red. And I love oranges",
		Price:       100.0,
		ImgUrl:      "https://www.allrecipes.com/recipes/1295/fruits-and-vegetables/fruits/citrus/oranges/", // Orange recipes page with images :contentReference[oaicite:2]{index=2}
	}

	prd2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "Fresh apples, crisp and sweet",
		Price:       120.0,
		ImgUrl:      "https://www.allrecipes.com/recipe/14354/sunday-best-fruit-salad/", // Fruit salad including apples :contentReference[oaicite:3]{index=3}
	}

	prd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Ripe yellow bananas full of energy",
		Price:       60.0,
		ImgUrl:      "https://www.allrecipes.com/recipe/214947/perfect-summer-fruit-salad/", // Summer fruit salad with bananas :contentReference[oaicite:4]{index=4}
	}

	prd4 := &Product{
		ID:          4,
		Title:       "Mango",
		Description: "Sweet and juicy tropical mangoes",
		Price:       150.0,
		ImgUrl:      "https://www.allrecipes.com/recipe/228681/fresh-fruit-salad/", // Fresh fruit salad with mango :contentReference[oaicite:5]{index=5}
	}

	prd5 := &Product{
		ID:          5,
		Title:       "Strawberry",
		Description: "Fresh strawberries, bright red and delicious",
		Price:       200.0,
		ImgUrl:      "https://www.allrecipes.com/recipe/214947/perfect-summer-fruit-salad/", // Fruit salad including strawberries :contentReference[oaicite:6]{index=6}
	}

	prd6 := &Product{
		ID:          6,
		Title:       "Pineapple",
		Description: "Tropical pineapple with a sweet-tangy taste",
		Price:       180.0,
		ImgUrl:      "https://www.allrecipes.com/recipe/214947/perfect-summer-fruit-salad/", // Fruit salad with pineapple :contentReference[oaicite:7]{index=7}
	}

	prd7 := &Product{
		ID:          7,
		Title:       "Grapes",
		Description: "Fresh grapes, seedless and juicy",
		Price:       140.0,
		ImgUrl:      "https://www.allrecipes.com/recipe/214947/perfect-summer-fruit-salad/", // Fruit salad with grapes :contentReference[oaicite:8]{index=8}
	}

	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)
	ProductList = append(ProductList, prd4)
	ProductList = append(ProductList, prd5)
	ProductList = append(ProductList, prd6)
	ProductList = append(ProductList, prd7)

}

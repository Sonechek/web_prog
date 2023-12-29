package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type productType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageURL    string `json:"img_url"`
}

// var products []productType = []productType{
// 	{
// 		Name:        "Товар 1",
// 		Description: "Описание товара 1",
// 		Price:       100,
// 		ImageURL:    "/assets/img/t1.jpeg",
// 	},
// 	{
// 		Name:        "Товар 2",
// 		Description: "Описание товара 2",
// 		Price:       200,
// 		ImageURL:    "/assets/img/t2.jpeg",
// 	},
// 	{
// 		Name:        "Товар 3",
// 		Description: "Описание товара 3",
// 		Price:       300,
// 		ImageURL:    "/assets/img/t3.jpeg",
// 	},
// 	{
// 		Name:        "Товар 4",
// 		Description: "Описание товара 4",
// 		Price:       400,
// 		ImageURL:    "/assets/img/t4.jpeg",
// 	},
// 	{
// 		Name:        "Товар 5",
// 		Description: "Описание товара 5",
// 		Price:       500,
// 		ImageURL:    "/assets/img/t5.jpeg",
// 	},
// 	{
// 		Name:        "Товар 6",
// 		Description: "Описание товара 6",
// 		Price:       600,
// 		ImageURL:    "/assets/img/t6.jpeg",
// 	},
// 	{
// 		Name:        "Товар 7",
// 		Description: "Описание товара 7",
// 		Price:       700,
// 		ImageURL:    "/assets/img/t7.jpeg",
// 	},
// 	{
// 		Name:        "Товар 8",
// 		Description: "Описание товара 8",
// 		Price:       800,
// 		ImageURL:    "/assets/img/t8.jpeg",
// 	},
// 	{
// 		Name:        "Товар 9",
// 		Description: "Описание товара 9",
// 		Price:       900,
// 		ImageURL:    "/assets/img/t9.jpeg",
// 	},
// }

var products []productType

const productsPerPage = 3
const jsonFileName = "./products.json"

func saveProducts() error {
	data, err := json.MarshalIndent(products, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(jsonFileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func loadProducts() error {
	data, err := os.ReadFile(jsonFileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &products)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	e := echo.New()

	// Регистрируем маршруты для статичных файлов
	e.Static("/public", "public")
	e.Static("/assets", "public/assets")

	err := loadProducts()
	if err != nil {
		panic(err)
	}

	// err := saveProducts()
	// if err != nil {
	// 	panic(err)
	// }

	e.POST("/getproducts", func(c echo.Context) error {
		type requestData struct {
			Page int `json:"page"`
		}
		var request requestData
		err := c.Bind(&request)
		if err != nil {
			return err
		}

		startIndex := (request.Page - 1) * productsPerPage
		endIndex := startIndex + productsPerPage
		if endIndex > len(products) {
			endIndex = len(products)
		}

		return c.JSON(http.StatusOK, products[startIndex:endIndex])
	})
	e.POST("/getproductcount", func(c echo.Context) error {
		type responseResult struct {
			Count           int `json:"count"`
			ProductsPerPage int `json:"products_per_page"`
		}

		result := responseResult{
			Count:           len(products),
			ProductsPerPage: productsPerPage,
		}
		return c.JSON(http.StatusOK, result)
	})
	// Пример обработчика GET с получением параметров
	// e.GET("/testget", func(c echo.Context) error {
	// 	name := c.QueryParam("name")

	// 	return c.String(http.StatusOK, "Добрый день, "+name)
	// })

	// // Пример обработчика запроса POST с получением параметров
	// e.POST("/testpost", func(c echo.Context) error {
	// 	json_map := make(map[string]interface{})
	// 	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	fmt.Println(json_map)
	// 	name := json_map["name"].(string)
	// 	v := map[string]interface{}{
	// 		"response": "Добрый день, " + name,
	// 	}

	// 	return c.JSON(http.StatusOK, v)
	// })

	// Основной обработчик GET / - отдает файл index.html
	e.GET("*", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

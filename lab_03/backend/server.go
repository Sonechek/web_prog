package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Регистрируем маршруты для статичных файлов
	e.Static("/public", "public")
	e.Static("/assets", "public/assets")

	// Пример обработчика GET с получением параметров
	e.GET("/testget", func(c echo.Context) error {
		name := c.QueryParam("name")

		return c.String(http.StatusOK, "Добрый день, "+name)
	})

	// Пример обработчика запроса POST с получением параметров
	e.POST("/testpost", func(c echo.Context) error {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return err
		}

		fmt.Println(json_map)
		name := json_map["name"].(string)
		v := map[string]interface{}{
			"response": "Добрый день, " + name,
		}

		return c.JSON(http.StatusOK, v)
	})

	// Основной обработчик GET / - отдает файл index.html
	e.GET("*", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

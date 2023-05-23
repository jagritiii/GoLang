package products

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

var e = echo.New()
var v = validator.New()

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("Unable to read configuration")
	}
}

func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside custom serverMessage")
		return next(c)
	}
}

func Start() {
	//we can use middleware for all endpoints at once
	e.Use(serverMessage)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hellooo")
	})
	e.GET("/products", GetProducts) //we can use middleware as third argument
	e.GET("/product/:id", GetProduct)
	e.POST("/products", CreateProduct)
	e.PUT("/product/:id", UpdateProduct)
	e.DELETE("/product/:id", DeleteProduct)

	e.Logger.Print("leistening at port %s", cfg.Port)
	e.Logger.Fatal(e.Start(":" + cfg.Port))
	//e.Logger.Fatal(e.Start(fmt.Sprint("localhost:%s", cfg.Port)))

}

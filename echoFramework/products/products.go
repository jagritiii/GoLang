package products

import (
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

// echo validator for product
type ProductValidator struct {
	validator *validator.Validate
}

// Validate validates reqBody for product
func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var products = []map[int]string{{1: "mobile"}, {2: "tv"}, {3: "laptop"}}

func GetProducts(c echo.Context) error {
	fmt.Printf("inside products")
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	var product map[int]string
	for _, p := range products {
		for k := range p {
			pID, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return err
			}
			if pID == k {
				product = p
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {
	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var reqBody body
	e.Validator = &ProductValidator{validator: v}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	//if err := v.Struct(reqBody); err != nil {
	//	return err
	//}

	if err := c.Validate(reqBody); err != nil {
		return err
	}
	product := map[int]string{
		len(products) + 1: reqBody.Name,
	}
	products = append(products, product)
	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for _, p := range products {
		for k, _ := range p {
			if k == pID {
				product = p
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var reqBody body
	e.Validator = &ProductValidator{validator: v}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	product[pID] = reqBody.Name
	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	var product map[int]string
	var index int
	for i, p := range products {
		for k := range p {
			pID, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return err
			}
			if pID == k {
				product = p
				index = i
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}

	splice := func(s []map[int]string, i int) []map[int]string {
		return append(s[:i], s[i+1:]...)
	}
	products = splice(products, index)
	return c.JSON(http.StatusOK, product)
}

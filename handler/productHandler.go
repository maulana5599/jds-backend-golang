package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ProductData struct {
	Id         string `json:"id"`
	Price      string `json:"price"`
	Product    string `json:"product"`
	CreatedAt  string `json:"createdAt"`
	Department string `json:"department"`
}

func GetProduct(ctx *fiber.Ctx) error {

	req, err := http.NewRequest("GET", "https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	jsonBytes, errRead := io.ReadAll(resp.Body)

	if errRead != nil {
		log.Fatal(errRead)
	}

	var dataProduct []ProductData
	er := json.Unmarshal(jsonBytes, &dataProduct)

	if er != nil {
		log.Fatal(er)
	}

	return ctx.JSON(fiber.Map{
		"message": "Get data successfully",
		"data":    dataProduct,
	})

}

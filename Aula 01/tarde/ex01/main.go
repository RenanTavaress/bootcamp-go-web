package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"slices"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

type produto struct {
	Id            int     `json:"id"`
	Nome          string  `json:"nome"`
	Cor           string  `json:"cor"`
	Preco         float64 `json:"preco"`
	Estoque       int     `json:"estoque"`
	Codigo        string  `json:"codigo"`
	Publicacao    bool    `json:"publicacao"`
	DataDeCriacao string  `json:"data_de_criacao"`
}

func Filter(ctx *gin.Context) {
	sortArr := ctx.Query("sort")
	sortAscOrDesc := ctx.Query("sortDirection")
	var productMap []produto
	jsonPessoal, err := os.ReadFile("../../product.json")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": "nao foi possivel ler o arquvi JSON",
		})
		return
	}
	json.Unmarshal(jsonPessoal, &productMap)
	switch sortArr {
	case "id":
		sort.Slice(productMap, func(i, j int) bool { return productMap[i].Id < productMap[j].Id })
		fmt.Println(productMap)
	case "nome":
		sort.Slice(productMap, func(i, j int) bool { return productMap[i].Nome < productMap[j].Nome })
	case "preco":
		sort.Slice(productMap, func(i, j int) bool { return productMap[i].Preco < productMap[j].Preco })
	}

	switch sortAscOrDesc {
	case "desc":
		slices.Reverse(productMap)
	}

	//fmt.Println(productMap)
	ctx.JSON(http.StatusOK, productMap)
}

func FilterById(ctx *gin.Context) {
	productId := ctx.Param("id")
	var product []produto
	jsonPessoal, err := os.ReadFile("../../product.json")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": "nao foi possivel ler o arquvi JSONnnn",
		})
		return
	}

	json.Unmarshal(jsonPessoal, &product)

	marks, _ := strconv.Atoi(productId)

	for _, value := range product {
		if value.Id == marks {
			ctx.JSON(http.StatusOK, value)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"not found": "id not found",
	})

}

func main() {
	router := gin.Default()

	router.GET("/produtos/:id", FilterById)
	router.GET("/produtos", Filter)

	router.Run(":8080")
}

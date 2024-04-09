package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type produto struct {
	Id            int       `json:id`
	Nome          string    `json:nome`
	Cor           string    `json:cor`
	Preco         float64   `json:preco`
	Estoque       int       `json:estoque`
	Codigo        string    `json:codigo`
	Publicacao    bool      `json:publicacao`
	dataDeCriacao time.Time `json:data_de_criacao`
}

func GetAll(ctx *gin.Context) {
	jsonPessoal, err := os.ReadFile("../../product.json")
	var list []produto
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": "nao foi possivel ler o arquvi JSON",
		})
		return
	}

	json.Unmarshal(jsonPessoal, &list)

	ctx.JSON(200, list)
}

func main() {
	router := gin.Default()

	router.GET("/produtos", GetAll)

	router.Run(":8080")
}

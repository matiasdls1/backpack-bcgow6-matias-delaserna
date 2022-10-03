package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Transaction struct {
	ID       int     `json:"id"`
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Date     string  `json:"date"`
}

func GetAll(c *gin.Context) {
	transactions := []Transaction{}
	data, err := ioutil.ReadFile("transactions.json")
	if err != nil {
		c.JSON(500, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
	}
	err = json.Unmarshal(data, &transactions)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"status": "ok",
			"data":   transactions,
		})
	}
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Matias",
		})
	})

	router.GET("/transactions", GetAll)
	router.Run()
}

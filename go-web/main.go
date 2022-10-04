package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
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
	data, err := os.ReadFile("transactions.json")
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

func GetFiltered(c *gin.Context) {
	transactions := []Transaction{}
	filtered := []*Transaction{}
	data, err := os.ReadFile("transactions.json")
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
	}
	for _, tx := range transactions {
		if c.Query("id") == strconv.Itoa(tx.ID) && c.Query("code") == tx.Code && c.Query("currency") == tx.Currency && c.Query("amount") == fmt.Sprintf("%.2f", tx.Amount) && c.Query("sender") == tx.Sender && c.Query("receiver") == tx.Receiver && c.Query("date") == tx.Date {
			filtered = append(filtered, &tx)
		}
	}
	c.JSON(200, gin.H{
		"status": "ok",
		"data":   filtered,
	})
}

func GetByID(c *gin.Context) {
	transactions := []Transaction{}
	ok := false
	data, err := os.ReadFile("transactions.json")
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
		for _, tx := range transactions {
			if c.Param("id") == strconv.Itoa(tx.ID) {
				c.JSON(200, gin.H{
					"status": "ok",
					"data":   tx,
				})
				ok = true
			}
		}
		if !ok {
			c.JSON(404, gin.H{
				"status": "error",
				"error":  "no encontrada",
			})
		}
	}
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Matias",
		})
	})
	//router.GET("/transactions", GetAll)
	router.GET("/transactions", GetFiltered)
	router.GET("/transactions/:id", GetByID)
	router.Run()
}

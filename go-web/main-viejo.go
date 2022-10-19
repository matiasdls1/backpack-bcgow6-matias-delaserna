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
	Code     string  `json:"code" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
	Sender   string  `json:"sender" binding:"required"`
	Receiver string  `json:"receiver" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}

var transactionsMemory []Transaction
var lastID int

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

func ValidateFields(req Transaction) ([]string, bool) {
	fields := []string{}
	if req.ID == 0 {
		fields = append(fields, "ID")
	}
	if req.Code == "" {
		fields = append(fields, "Code")
	}
	if req.Currency == "" {
		fields = append(fields, "Currency")
	}
	if req.Amount == 0.0 {
		fields = append(fields, "Amount")
	}
	if req.Sender == "" {
		fields = append(fields, "Sender")
	}
	if req.Receiver == "" {
		fields = append(fields, "Receiver")
	}
	if req.Date == "" {
		fields = append(fields, "Date")
	}
	if len(fields) == 0 {
		return fields, false
	} else {
		return fields, true
	}
}
func Create(c *gin.Context) {
	var req Transaction
	token := c.GetHeader("token")
	if token != "mdls" {
		c.JSON(401, gin.H{
			"status": "error",
			"error":  "no tiene permisos para realizar la peticion solicitada",
		})
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(404, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	// Validate fields
	fields, err := ValidateFields(req)
	if err {
		message, _ := fmt.Printf("el/los campo/s %v es/son requerido/s", fields)
		c.JSON(400, gin.H{
			"status": "error",
			"error":  message,
		})
		return
	}

	// Check if there are tx in memory
	if len(transactionsMemory) == 0 { // If no tx in memory, get lastID from database
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
		}
		lastID = transactions[len(transactions)-1].ID + 1
		req.ID = lastID
	} else { // If there are tx in memory, get the incremental ID
		lastID++
		req.ID = lastID
	}
	transactionsMemory = append(transactionsMemory, req)
	c.JSON(200, gin.H{
		"status": "ok",
		"data":   req,
	})

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
	router.POST("/transactions", Create)
	router.Run()
}

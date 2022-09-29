package main

import "github.com/matiasdls1/backpack-bcgow6-matias-delaserna/hackaton/internal/service"

func main() {
	var tickets []service.Ticket

	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
}

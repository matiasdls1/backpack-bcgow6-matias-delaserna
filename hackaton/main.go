package main

import "hackaton/internal/service"

func main() {
	var tickets []service.Ticket

	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
}

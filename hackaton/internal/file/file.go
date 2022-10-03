package file

import (
	"fmt"
	"hackaton/internal/service"
	"os"
)

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	csvFile, err := os.Open("../../tickets.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	return nil, nil
}

func (f *File) Write(service.Ticket) error {
	return nil
}

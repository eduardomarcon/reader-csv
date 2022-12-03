package usecases

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"reader/internal/configs"
)

func readFile(chanUser chan User) (*os.File, error) {
	fileConfig := configs.GetFile()
	pathFile := fmt.Sprintf("%s", fileConfig.Name)

	file, err := os.Open(pathFile)
	if err != nil {
		return file, err
	}

	reader := csv.NewReader(file)
	go func() {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			user := User{
				Name:  record[0],
				Email: record[1],
			}
			fmt.Println("read line - user", user.Name)
			chanUser <- user
		}

		close(chanUser)
	}()

	return file, nil
}

package order

import (
	"fmt"

	"GoTemplate/internal/repository"
)

func Create(data map[string]interface{}) error {
	_, err := repository.Create("orders", data)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

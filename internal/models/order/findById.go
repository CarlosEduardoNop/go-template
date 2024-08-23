package order

import (
	"fmt"

	"GoTemplate/internal/repository"
)

func FindById(id int) (*Order, error) {
	res, err := repository.FindById("orders", id)

	if err != nil {
		return nil, err
	}

	var order Order

	err = res.Scan(&order.ID, &order.Name, &order.Price)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &order, nil
}

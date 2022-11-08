package entity

import "errors"

const (
	INVALID_ID    = "Invalid ID."
	INVALID_TAX   = "Invalid Tax."
	INVALID_PRICE = "Invalid Price."
)

type OrderRepositoryInterface interface {
	Save(order *Order) error
}

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:         id,
		Price:      price,
		Tax:        tax,
		FinalPrice: price + tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New(INVALID_ID)
	}
	if o.Price <= 0 {
		return errors.New(INVALID_PRICE)
	}
	if o.Tax < 0 {
		return errors.New(INVALID_TAX)
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}

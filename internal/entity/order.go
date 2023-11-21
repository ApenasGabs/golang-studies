package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func (o *Order) CalculateFinalPrice() {
	o.FinalPrice = o.Price + o.Tax
}
func NewOrder(id string, price, tax float64) *Order {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.Validate()
	if err != nil {
		return nil
	}
	order.CalculateFinalPrice()
	return order
}

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("the ID field can not be empty")
	}
	if o.Price <= 0 {
		return errors.New("the Price field must be greater than zero")
	}
	if o.Tax < 0 {
		return errors.New("the Tax field can not be negative")
	}
	return nil
}

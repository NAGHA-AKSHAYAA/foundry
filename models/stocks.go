package models

import "foundry/initialisers"

type Stocks struct {
	ID       int     `gorm:"primaryKey" json:"id"`
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
}

func (stock *Stocks) Create() (*Stocks, error) {
	results := initialisers.DB.Create(&stock).Error

	if results != nil {
		return nil, results
	}

	return stock, nil
}

func Find(stockName string) (*Stocks, error) {
	var stock Stocks

	err := initialisers.DB.Where("name = ?", stockName).First(&stock).Error

	if err != nil {
		return nil, err
	}

	return &stock, nil
}

func (stock *Stocks) Update() (*Stocks, error) {
	results := initialisers.DB.Save(&stock).Error

	if results != nil {
		return nil, results
	}

	return stock, nil
}

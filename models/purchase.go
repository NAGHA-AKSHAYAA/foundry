package models

import "foundry/initialisers"

type Purchase struct {
	ID                  int     `gorm:"primaryKey" json:"id"`
	Name                string  `json:"name"`
	SupplierName        string  `json:"supplier_name"`
	Movement            bool    `json:"movement"`
	Date                string  `json:"date"`
	GrossWeight         float32 `json:"gross_weight"`
	BagWeight           float32 `json:"bag_weight"`
	MiscellaneousWeight float32 `json:"misc_weight"`
	NetWeight           float32 `json:"net_weight"`
}

func (purchase *Purchase) Create() (*Purchase, error) {
	results := initialisers.DB.Create(&purchase).Error

	if results != nil {
		return nil, results
	}

	return purchase, nil
}

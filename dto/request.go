package dto

type UserDetails struct {
	Role     string `json:"role"`
	Password string `json:"password"`
}

type ScrapeInput struct {
	Quantity            float32 `json:"quantity"`
	Name                string  `json:"name"`
	SupplierName        string  `json:"supplier_name"`
	Movement            bool    `json:"movement"`
	Date                string  `json:"date"`
	BagWeight           float32 `json:"bag_weight"`
	MiscellaneousWeight float32 `json:"misc_weight"`
	NetWeight           float32 `json:"net_weight"`
}

type ProductionLineInput struct {
}

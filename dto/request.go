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

type ProductionProcess struct {
	PlateCount       int       `json:"plate_count"` //6.3 weight
	CircleSize       float32   `json:"circle_size"`
	ScrapeNames      []string  `json:"scrape_names"`
	ScrapeQuantities []float32 `json:"scrape_quantities"`
}

type ProductionLineInput struct {
	CruisibleCount      int                 `json:"cruisible_count"` //470 kgs
	ProcessManipulation []ProductionProcess `json:"process_manip"`
}

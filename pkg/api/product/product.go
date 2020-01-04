package product

type Product struct {
	ID       string  `json:"_id" bson:"_id,omitempty"`
	Locale   string  `json:"locale" bson:"locale"`
	Name     string  `json:"name" bson:"name"`
	Price    float64 `json:"price" bson:"price"`
	Location string  `json:"location" bson:"location"`
}

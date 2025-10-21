package Model

type ProductHistory struct {
	ID        int     `json:"id"`
	ProductID int     `json:"product_id"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	ChangedAt string  `json:"changed_at"`
}

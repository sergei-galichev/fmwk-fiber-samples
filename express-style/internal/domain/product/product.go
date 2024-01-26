package product

type Product struct {
	ID          int64  `json:"id,omitempty"`
	Amount      int64  `json:"amount"`
	Name        string `json:"product_name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

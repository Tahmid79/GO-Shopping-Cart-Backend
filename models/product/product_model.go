package product

type Product struct {
	ID          string `json:"_id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

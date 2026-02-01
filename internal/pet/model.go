package pet

type Pet struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Species string  `json:"species"`
	Age     int     `json:"age"`
	Price   float64 `json:"price"`
	Status  string  `json:"status"`
}

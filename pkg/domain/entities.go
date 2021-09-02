package domain

type Product struct {
	ID    int     `json: id`
	Name  string  `json: prod_name`
	Value float64 `json: prod_value`
}

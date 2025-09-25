package domain

type RequestBook struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Isbn     string `json:"isbn"`
	Quantity int    `json:"quantity"`
	Category string `json:"category"`
}

type ResponseBook struct {
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
	TotalPage   int    `json:"total_page"`
	CurrentPage int    `json:"current_page"`
	Books       []Book `json:"books"`
}

type Book struct {
	Id       uint   `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Isbn     string `json:"isbn"`
	Quantity int    `json:"quantity"`
	Category string `json:"category"`
}

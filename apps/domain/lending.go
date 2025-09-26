package domain

type RequestLending struct {
	BookId uint `json:"book_id"`
}

type RequestReturnBook struct {
	BookId uint `json:"book_id"`
}
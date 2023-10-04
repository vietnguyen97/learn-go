package dto

type CreateBookInput struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

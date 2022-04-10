package dto

type AuthorizedUser struct {
    ID uint64 `json:"id"`
    Password string `json:"password"`
    Email string `json:"email"`
} 

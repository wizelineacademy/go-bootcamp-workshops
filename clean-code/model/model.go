package model

// Todo struct (Model)
type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Status    string `json:"status"`
	IsDeleted bool   `json:"isDeleted"`
}

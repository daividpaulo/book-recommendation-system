package entities

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Subject     string `json:"subject"`
	Area        string `json:"area"`
	Description string `json:"description"`
}

type User struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Age              int      `json:"age"`
	Profession       string   `json:"profession"`
	InterestAreas    []string `json:"interest_areas"`
	PurchaseCount    int      `json:"purchase_count,omitempty"`
	PurchasedBookIDs []string `json:"purchased_book_ids,omitempty"`
}

type Purchase struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	BookID   string `json:"book_id"`
	Quantity int    `json:"quantity"`
}

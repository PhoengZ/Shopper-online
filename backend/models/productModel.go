package models

type Product struct {
	ID          string   `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      string   `json:"uid" bson:"uid"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Price       int      `json:"price" bson:"price"`
	Quantity    int      `json:"quantity" bson:"quantity"`
	Category    []string `json:"category" bson:"category"`
	URL         string   `json:"url" bson:"url"`
}

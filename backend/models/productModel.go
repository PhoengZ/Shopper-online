package models

type Product struct {
	ID          string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Price       int      `json:"price" bson:"price"`
	Category    []string `json:"category" bson:"category"`
	URL         string   `json:"url" bson:"url"`
}

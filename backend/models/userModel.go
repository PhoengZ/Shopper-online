package models

type User struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	CartList []Item `json:"cartlist" bson:"cartlist"`
	Coin     int    `json:"coin" bson:"coin"`
}
type Item struct {
	ID       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Price    int    `json:"price" bson:"price"`
	URL      string `json:"url" bson:"url"`
	Quantity int    `json:"quantity" bson:"quantity"`
}

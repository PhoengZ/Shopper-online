package models

type User struct {
	ID       string                   `json:"id,omitempty" bson:"_id,omitempty"`
	Username string                   `json:"username" bson:"username"`
	Password string                   `json:"password" bson:"password"`
	CartList []map[string]interface{} `json:"cartlist" bson:"cartlist"`
	Coin     int                      `json:"coin" bson:"coin"`
}

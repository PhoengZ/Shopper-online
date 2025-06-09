package models

type Transaction struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	UserID string `json:"uid" bson:"uid"`
	State  string `json:"state" bson:"state"`
	Money  int    `json:"money" bson:"money"`
}

package model

type User struct {
	Name    string `json:"name" bson:"name"`
	Email   string `json:"email" bson:"email"`
	Picture string `json:"picture" bson:"picture"`
}

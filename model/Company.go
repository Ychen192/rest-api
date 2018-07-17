package model

type Company struct {
	Name  string `bson:"name" json:"name"`
	Tel   string `bson:"tel" json:"tel"`
	Email string `bson:"email" json:"email"`
}
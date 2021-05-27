package models

type User struct {
	name     string `bson:"nombre"`
	lastName string `bson:"apellido"`
	email    string `bson:"email"`
	dni      string `bson:"dni"`
	phone    string `bson:"numero"`
	state    string `bson:"estado"`
	nickname string `bson:"apodo"`
}

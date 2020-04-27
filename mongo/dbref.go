package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

// DBRef for mongodb
type DBRef struct {
	Ref	string	`bson:"$ref"`
	ID	primitive.ObjectID	`bson:"$id"`
	DB	string	`bson:"$db"`
}

// Image for mongo
type Image struct {
	ID      string   `json:"id" bson:"_id"`
	Project DBRef `json:"project" bson:"project"`
	File    DBRef    `json:"file" bson:"file"`
}
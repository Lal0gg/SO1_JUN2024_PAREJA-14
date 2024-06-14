package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Data struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Percent   string             `json:"percent"`
	Timestamp primitive.DateTime `json:"timestamp"`
}

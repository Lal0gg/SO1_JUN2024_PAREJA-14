package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Data struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Percent   string             `json:"percent"`
	Timestamp primitive.DateTime `json:"timestamp"`
}

type ProcessData struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PID       int                `json:"pid"`
	Name      string             `json:"name"`
	State     string             `json:"state"`
	PIDPadre  int                `json:"pidPadre"`
	Timestamp primitive.DateTime `json:"timestamp"`
}

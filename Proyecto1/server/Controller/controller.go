package Controller

import (
	"context"
	"log"
	"server/Instance"
	"server/Model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertData(nameCol string, dataParam string, timestamp primitive.DateTime) {
	collection := Instance.Mg.Db.Collection(nameCol)
	doc := Model.Data{
		ID:        primitive.NewObjectID(),
		Percent:   dataParam,
		Timestamp: timestamp,
	}

	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
}

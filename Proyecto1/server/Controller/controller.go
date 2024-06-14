package Controller

import (
	"context"
	"log"
	"server/Instance"
	"server/Model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertData(nameCol string, dataParam string) {
	collection := Instance.Mg.Db.Collection(nameCol)
	doc := Model.Data{
		ID:        primitive.NewObjectID(), // Generar un nuevo ObjectID
		Percent:   dataParam,
		Timestamp: primitive.NewDateTimeFromTime(time.Now()), // Añadir el timestamp actual
	}

	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
}

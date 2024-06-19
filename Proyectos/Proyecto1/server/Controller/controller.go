package Controller

import (
	"context"
	"log"
	"server/Instance"
	"server/Model"
	"time"

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
		log.Println("Error al insertar datos:", err)
	} else {
		log.Println("Datos insertados correctamente en la colección:", nameCol)
	}
}

func InsertProcessData(process Model.ProcessData) {
	collection := Instance.Mg.Db.Collection("process")

	process.ID = primitive.NewObjectID()
	process.Timestamp = primitive.NewDateTimeFromTime(time.Now())

	if process.PIDPadre == 0 {
		process.PIDPadre = 0
	}

	_, err := collection.InsertOne(context.TODO(), process)
	if err != nil {
		log.Fatal(err)
	}
}

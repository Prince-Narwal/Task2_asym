package controller

import (
	"context"
	"fmt"
	"log"

	model "func.go/model02"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring = "mongodb+srv://narwalprince19:prince19@cluster0.0fpdf15.mongodb.net/"
const dbname = "AppData"
const colname = "Appname"

func Insert_Mongodb(data []model.NodeData) {
	// most important
	var collection *mongo.Collection

	//client options
	clientoptions := options.Client().ApplyURI(connectionstring)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientoptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection success")
	collection = client.Database(dbname).Collection(colname)

	//bsondata, err := bson.Marshal(data)
	/*var bsondata bson.M
	  err1 := bson.UnmarshalExtJSON(jsondata, true, &bsondata)
	  if err1 != nil {
	  	log.Fatal(err)
	  }*/
	type bsonData struct {
		ID         primitive.ObjectID `json:"_id,omitempty" bson:""`
		AppName    string             `json:"appname,omitempty"`
		AppId      string             `json:"appid,omitempty"`
		Version    string             `json:"version,omitempty"`
		State      int64              `json:"state,omitempty"`
		Devid      string             `json:"devid,omitempty"`
		VolumeRefs string             `json:"volumeRefs,omitempty"`
	}
	node := bsonData{
		AppName:    data[0].AppName,
		AppId:      data[0].AppId,
		Version:    data[0].Version,
		State:      data[0].State,
		Devid:      data[0].Devid,
		VolumeRefs: data[0].VolumeRefs,
	}
	inserted1, err := collection.InsertOne(context.Background(), node)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie in db with id: ", inserted1.InsertedID)

}

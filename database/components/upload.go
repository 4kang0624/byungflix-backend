package components

import (
	"byungflix-backend/database"
	"byungflix-backend/database/connection"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateSeries(series database.Series) {
	client, err := connection.ConnectMongo()
	if err != nil {
		return
	}
	defer connection.DisconnectMongo(client, err)

	collection := client.Database("byungflix").Collection("series")
	filter := bson.D{{"title", series.Title}}
	num, _ := collection.CountDocuments(context.TODO(), filter)
	data, _ := bson.Marshal(series)
	if num == 0 {
		_, err := collection.InsertOne(context.TODO(), data)
		if err != nil {
			fmt.Println("created series")
		}
	}
	return
}

func UploadVideo(video database.Video) {
	client, err := connection.ConnectMongo()
	if err != nil {
		return
	}
	defer connection.DisconnectMongo(client, err)

	collection := client.Database("byungflix").Collection("video")
	filter := bson.D{{"title", video.Title}}
	num, _ := collection.CountDocuments(context.TODO(), filter)
	data, _ := bson.Marshal(video)

	if num == 0 {
		_, err := collection.InsertOne(context.TODO(), data)
		if err != nil {
			fmt.Println("uploaded video")
		}
	}
	return
}

func UploadSubtitle(language string, path string, videoTitle string) {
	client, err := connection.ConnectMongo()
	if err != nil {
		return
	}
	defer connection.DisconnectMongo(client, err)

	collection := client.Database("byungflix").Collection("video")
	filter := bson.D{{"title", videoTitle}}
	num, _ := collection.CountDocuments(context.TODO(), filter)

	if num != 0 {
		update := bson.D{{"$set", bson.D{{"subtitlepath." + language, path}}}}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			fmt.Println("uploaded subtitle")
		}
	}
	return
}

package components

import (
	"byungflix-backend/database"
	"byungflix-backend/database/connection"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSeriesList(title string) []database.Series {
	client, err := connection.ConnectMongo()
	if err != nil {
		return nil
	}
	defer connection.DisconnectMongo(client, err)

	collection := client.Database("byungflix").Collection("series")

	filter := bson.D{{"$or", bson.A{bson.D{{"title", primitive.Regex{Pattern: title, Options: "i"}}}, bson.D{{"titlekor", primitive.Regex{Pattern: title, Options: "i"}}}}}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	var seriesList []database.Series
	for cursor.Next(context.Background()) {
		var series database.Series
		err := cursor.Decode(&series)
		if err != nil {
			return nil
		}
		seriesList = append(seriesList, series)
	}
	return seriesList
}

func GetVideoListBySeriesTitle(title string) []database.Video {
	client, err := connection.ConnectMongo()
	if err != nil {
		return nil
	}
	defer connection.DisconnectMongo(client, err)

	collection := client.Database("byungflix").Collection("video")

	filter := bson.D{{"seriestitle", title}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	var videoList []database.Video
	for cursor.Next(context.Background()) {
		var video database.Video
		err := cursor.Decode(&video)
		if err != nil {
			return nil
		}
		videoList = append(videoList, video)
	}
	return videoList
}

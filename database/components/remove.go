package components

import (
	"byungflix-backend/database/connection"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

func RemoveSeries(SeriesName string) {
	// 시리즈 이름으로 시리즈를 찾아서 삭제한다.
	client, err := connection.ConnectMongo()
	if err != nil {
		return
	}
	defer connection.DisconnectMongo(client, err)

	filter := bson.M{"title": SeriesName}
	SeriesCollection := client.Database("byungflix").Collection("series")
	SeriesCollection.DeleteOne(context.TODO(), filter)

	os.RemoveAll("contents/" + SeriesName)
	return
}

func RemoveVideo(SeriesName string, EpisodeCount int) {
	client, err := connection.ConnectMongo()
	if err != nil {
		return
	}
	defer connection.DisconnectMongo(client, err)

	filter := bson.M{"seriestitle": SeriesName, "episodecount": EpisodeCount}
	VideoCollection := client.Database("byungflix").Collection("video")
	VideoCollection.DeleteOne(context.TODO(), filter)

	remdirerr := os.RemoveAll("contents/" + SeriesName + "/" + fmt.Sprintf("%d", EpisodeCount))
	if remdirerr != nil {
		fmt.Println(remdirerr)
	}
	return
}

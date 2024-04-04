package components

import (
	"byungflix-backend/database"
	"byungflix-backend/database/connection"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

func UpdadeSeries(oldTitle string, series database.Series) {
	client, err := connection.ConnectMongo()
	if err != nil {
		return
	}
	defer connection.DisconnectMongo(client, err)

	collection := client.Database("byungflix").Collection("series")

	filterSeries := bson.M{"title": oldTitle}
	updateSeries := bson.M{"$set": bson.M{
		"title":       series.Title,
		"titlekor":    series.TitleKor,
		"description": series.Description,
	}}

	/*
		비디오 컬렉션에서의 시리즈 타이틀 변경, UpdateSeriesTitleInVideo 함수 구현해서 searchVideoBySeriesTitle로 찾아서 반복문 돌리기
		filterVideo := bson.M{"seriestitle": oldTitle}
		updateVideo := bson.M{"$set": bson.M{
			"seriestitle": series.Title,
		}
	*/

	err = os.Rename("contents/"+oldTitle, "contents/"+series.Title)
	if err != nil {
		return
	}

	_, err = collection.UpdateOne(context.TODO(), filterSeries, updateSeries)
	if err != nil {
		return
	}
	return
}

func UpdateSeriesTitleInVideo(oldSeriesTitle string, video database.Video) {
	// title, seriseTitle, videoPath, videoPathHls 변경, subtitlePath 입력값이 있다면 subtitlePath도 변경해야 함
	// mkv 및 m3u8 파일 제목 변경해야 함
}

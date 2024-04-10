package components

import (
	"byungflix-backend/database"
	"byungflix-backend/database/connection"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"strconv"
	"strings"
)

func UpdadeSeries(oldTitle string, series database.Series) {
	client, err := connection.ConnectMongo()
	if err != nil {
		return
	}
	defer connection.DisconnectMongo(client, err)

	collectionSeries := client.Database("byungflix").Collection("series")
	collectionVideo := client.Database("byungflix").Collection("video")

	filterSeries := bson.M{"title": oldTitle}
	updateSeries := bson.M{"$set": bson.M{
		"title":       series.Title,
		"titlekor":    series.TitleKor,
		"description": series.Description,
	}}

	_, errSeries := collectionSeries.UpdateOne(context.TODO(), filterSeries, updateSeries)
	if errSeries != nil {
		return
	}

	newVideoDBList := UpdateTitleInFileSystem(oldTitle, series.Title)
	for _, video := range newVideoDBList {
		filterVideo := bson.M{
			"title":        video.Title,
			"episodecount": video.EpisodeCount,
		}

		updateVideo := bson.M{"$set": bson.M{
			"title":        video.Title,
			"seriestitle":  series.Title,
			"episodecount": video.EpisodeCount,
			"releasedate":  video.ReleaseDate,
			"uploaddate":   video.UploadDate,
			"videopath":    video.VideoPath,
			"videopathhls": video.VideoPathHls,
			"subtitlepath": video.SubtitlePath,
		}}

		_, errVideo := collectionVideo.UpdateOne(context.TODO(), filterVideo, updateVideo)
		if errVideo != nil {
			return
		}
	}

	return
}

func UpdateTitleInFileSystem(oldTitle string, newTitle string) []database.Video {
	TitleInDB := GetVideoListBySeriesTitle(oldTitle)
	err := os.Rename("contents/"+oldTitle, "contents/"+newTitle)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for i := range TitleInDB {
		for key, value := range TitleInDB[i].SubtitlePath {
			TitleInDB[i].SubtitlePath[key] = strings.Replace(value, "/"+oldTitle+"/", "/"+newTitle+"/", -1)
		}
		TitleInDB[i].VideoPath = strings.Replace(TitleInDB[i].VideoPath, "/"+oldTitle+"/", "/"+newTitle+"/", -1)
		TitleInDB[i].VideoPathHls = strings.Replace(TitleInDB[i].VideoPathHls, "/"+oldTitle+"/", "/"+newTitle+"/", -1)
	}

	for i := range TitleInDB {
		for key, value := range TitleInDB[i].SubtitlePath {
			err := os.Rename(value, strings.Replace(value, "/"+oldTitle+"_"+strconv.Itoa(TitleInDB[i].EpisodeCount), "/"+newTitle+"_"+strconv.Itoa(TitleInDB[i].EpisodeCount), -1))
			TitleInDB[i].SubtitlePath[key] = strings.Replace(value, "/"+oldTitle+"_"+strconv.Itoa(TitleInDB[i].EpisodeCount), "/"+newTitle+"_"+strconv.Itoa(TitleInDB[i].EpisodeCount), -1)
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}

		err := os.Rename(TitleInDB[i].VideoPath, strings.Replace(TitleInDB[i].VideoPath, "/"+oldTitle+"_"+strconv.Itoa(TitleInDB[i].EpisodeCount), "/"+newTitle+"_"+strconv.Itoa(TitleInDB[i].EpisodeCount), -1))
		TitleInDB[i].VideoPath = strings.Replace(TitleInDB[i].VideoPath, "/"+oldTitle+"_"+strconv.Itoa(TitleInDB[i].EpisodeCount), "/"+newTitle+"_"+strconv.Itoa(TitleInDB[i].EpisodeCount), -1)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		err = os.Rename(TitleInDB[i].VideoPathHls, strings.Replace(TitleInDB[i].VideoPathHls, "/"+oldTitle+"_"+strconv.Itoa(TitleInDB[i].EpisodeCount), "/"+newTitle+"_"+strconv.Itoa(TitleInDB[i].EpisodeCount), -1))
		TitleInDB[i].VideoPathHls = strings.Replace(TitleInDB[i].VideoPathHls, "/"+oldTitle+"_"+strconv.Itoa(TitleInDB[i].EpisodeCount), "/"+newTitle+"_"+strconv.Itoa(TitleInDB[i].EpisodeCount), -1)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	return TitleInDB
}

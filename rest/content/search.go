package content

import (
	"byungflix-backend/database/components"
	"fmt"
	"net/http"
)

func SearchSeries(rw http.ResponseWriter, r *http.Request) {
	fmt.Print("Series title: ")
	var title string
	fmt.Scanln(&title)

	seriesList := components.GetSeriesList(title)

	for _, series := range seriesList {
		fmt.Println(series)
	}
}

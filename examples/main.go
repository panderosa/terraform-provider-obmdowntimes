package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/panderosa/obmprovider/downtime"
	"github.com/panderosa/obmprovider/obmsdk"
)

func main() {

	downtimeId := "01d9a07112e248d68976411305e07442"
	address := obmsdk.GetEnv("OBM_BASE_URL")
	basepath := obmsdk.GetEnv("OBM_DOWNTIME_PATH")
	username := obmsdk.GetEnv("OBM_BA_USER")
	password := obmsdk.GetEnv("OBM_BA_PASSWORD")

	client, err := obmsdk.NewClient(address, basepath, username, password)
	if err != nil {
		log.Fatal(err)
	}

	dnt, err := client.Downtimes.Read(downtimeId)
	if err != nil {
		log.Fatal(err)
	}
	cis2 := downtime.Flatten2CIs(dnt.SelectedCIs)
	fmt.Print(cis2)

	cis3 := downtime.Flatten3CIs(dnt.SelectedCIs)
	fmt.Print(cis3)

}

func PrintTime() {
	dt := time.Now()
	layout := "2006-01-02T15:04:05-07:00"
	datatime := fmt.Sprint(dt.Format(layout))
	fmt.Printf("%s\n", datatime)
	loc := dt.Location()
	fmt.Printf("loc: %v\n", loc.String())
	loc1, err := time.LoadLocation("Europe/Warsaw")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("loc: %v\n", loc1)
}

func flattenCIs(data []obmsdk.Ci) string {
	array := make([]string, 0, len(data))
	for i := range data {
		array = append(array, data[i].ID)
	}
	return strings.Join(array, ",")
}

func mapCIs(data string) []obmsdk.Ci {
	array := strings.Split(data, ",")
	cis := make([]obmsdk.Ci, 0, len(array))
	for i := range array {
		cis = append(cis, obmsdk.Ci{
			ID: array[i],
		})
	}
	return cis
}

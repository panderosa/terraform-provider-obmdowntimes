package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/panderosa/obmprovider/obmsdk"
)

func main() {
	PrintTime()
	/*item := obmsdk.Ci{
		ID: "123",
	}
	// works
	//items := []obmsdk.Ci{item}
	items := make([]obmsdk.Ci, 0, 1)
	items = append(items, item)

	options := obmsdk.DowntimeCreateOptions{
		UserId:      "1",
		Planned:     "true",
		Name:        "dada",
		Description: "opis",
		Approver:    "Dariusz Malinoski",
		Category:    "2",
		SelectedCIs: mapCIs("1234"),
	}
	options.Schedule.StartDate = "2022-02-25T14:40:00+01:00"
	options.Schedule.EndDate = "2022-02-25T14:40:00+01:00"
	options.Action.Type = "OS Monitoring"

	fmt.Println(flattenCIs(options.SelectedCIs))*/

	/*data, err := xml.MarshalIndent(options, " ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))*/
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

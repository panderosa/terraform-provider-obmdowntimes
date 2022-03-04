package main

import (
	"flag"
	"log"

	sdk "github.com/panderosa/obmprovider/obmsdk"
)

func main() {

	requestPtr := flag.String("request", "", "action on OBM Downtime Service: create, read, delete, update, filter")
	downtimePtr := flag.String("id", "", "OBM downtime ID")
	filenamePtr := flag.String("filename", "", "XML file with options to create and update OBM Downtime, or JSON file with filters to search")

	flag.Parse()

	address := sdk.GetEnv("OBM_BASE_URL")
	basepath := sdk.GetEnv("OBM_DOWNTIME_PATH")
	username := sdk.GetEnv("OBM_BA_USER")
	password := sdk.GetEnv("OBM_BA_PASSWORD")

	client, err := sdk.NewClient(address, basepath, username, password)
	if err != nil {
		log.Fatalf(err.Error())
	}

	request := *requestPtr
	downtimeID := *downtimePtr
	filename := *filenamePtr

	switch request {
	case "create":
		options, err := loadCreateXML(filename)
		if err != nil {
			log.Fatalln(err)
		}
		create(client, options)
	case "update":
		options, err := loadUpdateXML(filename)
		if err != nil {
			log.Fatalln(err)
		}
		options.ID = downtimeID
		update(client, downtimeID, options)
	case "delete":
		if downtimeID == "" {
			panic("Downtime ID is empty")
		}
		delete(client, downtimeID)
	case "filter":
		v := loadFilters(filename)
		filter(client, v)
	case "read":
		if downtimeID == "" {
			panic("Downtime ID is empty")
		}
		read(client, downtimeID)
	default:
		log.Fatal("Invalid request")
	}
}

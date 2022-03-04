package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"

	sdk "github.com/panderosa/obmprovider/obmsdk"
)

func create(client *sdk.Client, options *sdk.DowntimeCreateOptions) {
	created, err := client.Downtimes.Create(*options)
	if err != nil {
		log.Fatal(err)
	}

	bibi, err := xml.MarshalIndent(created, " ", "   ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Downtime %v was created\n", created.ID)
	fmt.Println(string(bibi))
}

func update(client *sdk.Client, downtimeID string, options *sdk.Downtime) {
	err := client.Downtimes.Update(downtimeID, *options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated ...")
}

func read(client *sdk.Client, downtimeID string) {
	dnt, err := client.Downtimes.Read(downtimeID)
	if err != nil {
		log.Fatal(err)
	}
	barray, err := xml.MarshalIndent(dnt, " ", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(barray))
}

func filter(client *sdk.Client, v interface{}) {
	dnts, err := client.Downtimes.Search(v)
	if err != nil {
		log.Fatal(err)
	}
	barray, err := xml.MarshalIndent(dnts, " ", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(barray))
}

func delete(client *sdk.Client, downtimeID string) {
	err := client.Downtimes.Delete(downtimeID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted ...")
}

func loadCreateXML(filename string) (*sdk.DowntimeCreateOptions, error) {
	v := &sdk.DowntimeCreateOptions{}
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("loadCreateXML: %v", err)
	}

	se := xml.Unmarshal(data, v)
	if se != nil {
		return nil, fmt.Errorf("loadCreateXML: %v", err)
	}
	return v, nil
}

func loadUpdateXML(filename string) (*sdk.Downtime, error) {
	v := &sdk.Downtime{}
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("loadCreateXML: %v", err)
	}

	se := xml.Unmarshal(data, v)
	if se != nil {
		return nil, fmt.Errorf("loadCreateXML: %v", err)
	}
	return v, nil
}

func loadFilters(filename string) interface{} {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	//v := make(map[string]interface{})
	v := new(interface{})
	err = json.Unmarshal(data, v)
	vlist := (*v).([]interface{})
	queryMap := make(map[string]string)
	for _, e := range vlist {
		item := e.(map[string]interface{})
		queryMap[item["name"].(string)] = item["value"].(string)
	}
	if err != nil {
		log.Fatal(err)
	}
	return queryMap
}

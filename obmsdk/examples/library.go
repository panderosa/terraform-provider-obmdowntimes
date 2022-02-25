package main

import (
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

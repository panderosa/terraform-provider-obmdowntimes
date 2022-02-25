package obmsdk

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"net/url"
)

var ctx = context.Background()

// Downtimes describes all the OBM downtime related methods that the OBM Downtime Service supports.
type Downtimes interface {
	Create(options DowntimeCreateOptions) (*Downtime, error)
	Read(downtimeID string) (*Downtime, error)
	Delete(downtimeID string) error
	Update(downtimeID string, options Downtime) error
}

// downtime implements Downtimes
type downtimes struct {
	client *Client
}

type Recipient struct {
	XMLName xml.Name `xml:"recipient"`
	ID      string   `xml:"id,attr"`
}

type Action struct {
	XMLName xml.Name `xml:"action"`
	Type    string   `xml:"name,attr"`
}

type Ci struct {
	XMLName xml.Name `xml:"ci"`
	ID      string   `xml:"id"`
}

type Schedule struct {
	XMLName   xml.Name `xml:"schedule"`
	Type      string   `xml:"type"`
	StartDate string   `xml:"startDate"`
	EndDate   string   `xml:"endDate"`
	TimeZone  string   `xml:"timeZone"`
}

// Structure to store XML Dowtime
type Downtime struct {
	XMLName      xml.Name    `xml:"downtime"`
	ID           string      `xml:"id,attr"`
	Planned      string      `xml:"planned,attr"`
	Name         string      `xml:"name"`
	Description  string      `xml:"description"`
	Action       Action      `xml:"action"`
	Approver     string      `xml:"approver"`
	Category     string      `xml:"category"`
	SelectedCIs  []Ci        `xml:"selectedCIs>ci"`
	Notification []Recipient `xml:"notification>recipients>recipient"`
	Schedule     Schedule    `xml:"schedule"`
}

type DowntimeCreateOptions struct {
	XMLName      xml.Name  `xml:"downtime"`
	UserId       string    `xml:"userId,attr"`
	Planned      string    `xml:"planned,attr"`
	Name         string    `xml:"name"`
	Description  string    `xml:"description"`
	Action       Action    `xml:"action"`
	Approver     string    `xml:"approver"`
	Category     string    `xml:"category"`
	SelectedCIs  []Ci      `xml:"selectedCIs>ci"`
	Notification Recipient `xml:"notification>recipients>recipient"`
	Schedule     Schedule  `xml:"schedule"`
}

func (s *downtimes) Create(options DowntimeCreateOptions) (*Downtime, error) {
	fmt.Println("Creating OBM Downtime")
	req, err := s.client.newRequest("POST", "", options)
	if err != nil {
		return nil, err
	}
	ent := &Downtime{}
	err = s.client.do(ctx, req, ent)
	if err != nil {
		return nil, err
	}

	return ent, nil
}

func (s *downtimes) Update(downtimeID string, options Downtime) error {
	fmt.Println("Updating OBM Downtime")
	u := fmt.Sprintf("/%s", url.QueryEscape(downtimeID))
	req, err := s.client.newRequest("PUT", u, options)
	if err != nil {
		return fmt.Errorf("create request failed: %v", err)
	}
	return s.client.do(ctx, req, nil)
}

func (s *downtimes) Read(downtimeID string) (*Downtime, error) {
	if !validStringID(&downtimeID) {
		return nil, errors.New("invalid value for downtimeID")
	}
	fmt.Printf("Reading OBM Downtime %v\n", downtimeID)
	u := fmt.Sprintf("/%s", url.QueryEscape(downtimeID))
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	ent := &Downtime{}
	err = s.client.do(ctx, req, ent)
	if err != nil {
		return nil, err
	}
	return ent, nil
}

func (s *downtimes) Delete(downtimeID string) error {
	if !validStringID(&downtimeID) {
		return errors.New("invalid value for downtimeID")
	}
	u := fmt.Sprintf("/%s", url.QueryEscape(downtimeID))
	req, err := s.client.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	return s.client.do(ctx, req, nil)
}

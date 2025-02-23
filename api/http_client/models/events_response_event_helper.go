package models

import (
	"encoding/json"
	"github.com/MinterTeam/minter-go-sdk/v2/api"
)

// EventItem is the structure of the EventsResponse view of list items
type EventItem struct {
	Type  string
	Value map[string]string
}

// NewEventItem returns an EventItem from the EventsResponse list item
func NewEventItem(i interface{}) (*EventItem, error) {
	marshal, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	var event EventItem
	err = json.Unmarshal(marshal, &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// ConvertStructToEvent returns api.Event
func ConvertStructToEvent(data interface{}) (api.Event, error) {
	str, err := NewEventItem(data)
	if err != nil {
		return nil, err
	}

	value, err := json.Marshal(str.Value)
	if err != nil {
		return nil, err
	}

	event, err := api.ConvertToEvent(api.EventType(str.Type), value)
	if err != nil {
		return nil, err
	}

	return event, nil
}

package event

import "time"

type Event struct {
	Timestamp time.Time
	Message   string
	Fields    map[string]interface{}
	Tags      []string
}

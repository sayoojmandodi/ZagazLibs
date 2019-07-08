package eventpush

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//Trace struct holds details of trace
type Trace struct {
	Host string
	Port int
}

var trace Trace

// Events ...
type Events struct {
	EventArray []EventFrame `json:"events"   binding:"required,dive"`
}

// EventFrame ...
type EventFrame struct {
	EventLogTime         time.Time              `json:"event_time_stamp"`
	AppName              string                 `json:"app_name"   binding:"required"`
	AppVersion           string                 `json:"app_version" binding:"required"`
	EventType            string                 `json:"event_type"`
	Event                string                 `json:"event"`
	EventID              string                 `json:"event_id"`
	EventMessage         string                 `json:"event_message"`
	User                 string                 `json:"user" binding:"required"`
	AdditionalAttributes map[string]interface{} `json:"additional_attributes"`
}

//ConnectToTrace ...
func ConnectToTrace(t Trace) *Trace {
	trace = t
	return &trace
}

//PushEventToTrace pushes events to trace
func PushEventToTrace(events *Events) error {
	jsonInput, err := json.Marshal(events)
	if err != nil {
		return err
	}

	bytesInput := bytes.NewBuffer(jsonInput)
	client := http.Client{}
	request, err := http.NewRequest("POST", fmt.Sprintf("http://%s:%d/v1/event/log_event", trace.Host, trace.Port), bytesInput)
	if err != nil {
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		return err
	}
}

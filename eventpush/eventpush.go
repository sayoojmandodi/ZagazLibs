package eventpush

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	EventLogTime         time.Time              `json:"event_time_stamp" binding:"required"`
	AppName              string                 `json:"app_name"   binding:"required"`
	AppVersion           string                 `json:"app_version" binding:"required"`
	RequestID            string                 `json:"request_id" binding:"required"`
	EventType            string                 `json:"event_type" binding:"required"`
	Event                string                 `json:"event" binding:"required"`
	EventID              string                 `json:"event_id" binding:"required"`
	EventMessage         string                 `json:"event_message"`
	User                 string                 `json:"user"`
	AdditionalAttributes map[string]interface{} `json:"additional_attributes"`
}

//ConnectToTrace ...
func ConnectToTrace(t Trace) *Trace {
	trace = t
	return &trace
}

//PushEventToTrace pushes events to trace
func PushEventToTrace(events *Events) error {
	var eventObj Events
	jsonInput, err := json.Marshal(events)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(jsonInput, &eventObj); err != nil {
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

	return nil
}

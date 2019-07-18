package errorhandler

import (
	"encoding/json"
	"io/ioutil"
)

type errorStructure struct {
	Errorcode     int    `json:"error_code"`
	Httpcode      int    `json:"http_code"`
	Internalerror string `json:"internal_error"`
	Publicerror   string `json:"public_error"`
	Status        bool   `json:"status"`
}

//Error ...
var Error map[int]errorStructure

// ReadErrorsFromFile : read all errors from json file to Error map
func ReadErrorsFromFile(filepath string) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	var e []errorStructure
	if err = json.Unmarshal(data, &e); err != nil {
		return err
	}
	Error = make(map[int]errorStructure)
	for k := range e {
		Error[e[k].Errorcode] = e[k]
	}
	return nil
}

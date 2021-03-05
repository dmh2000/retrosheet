package main

import (
	"encoding/json"
	"io/ioutil"
)

var idFields = []string {
  "ID",
  "Last",
  "First",
  "PlayDebut",
  "MgrDebut",
  "CoachDebut",
  "UmpDebut",
}

type id struct {
	ID         string `json:"ID,omitempty"`
	Last       string `json:"Last,omitempty"`
	First      string `json:"First,omitempty"`
	PlayDebut  string  `json:"PlayDebut,omitempty"`
	MgrDebut   string `json:"MgrDebut,omitempty"`
	CoachDebut string `json:"CoachDebut,omitempty"`
	UmpDebut   string  `json:"UmpDebut,omitempty"`
}

func getIDNames() []string {
	return []string {
		"ID",
		"Last",
		"First",
		"PlayDebut",
		"MgrDebut",
		"CoachDebut",
		"UmpDebut",
	  }
}

func loadIDs(fname string) ([]id, error) {
	jsonBlob, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	var ids []id

	err = json.Unmarshal(jsonBlob, &ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
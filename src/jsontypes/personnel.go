package jsontypes

import (
	"encoding/json"
	"io/ioutil"
)

// Person - one instance of the data from a personnel json file
type Person struct {
	ID         string `json:"ID,omitempty"`
	Last       string `json:"Last,omitempty"`
	First      string `json:"First,omitempty"`
	PlayDebut  string  `json:"PlayDebut,omitempty"`
	MgrDebut   string `json:"MgrDebut,omitempty"`
	CoachDebut string `json:"CoachDebut,omitempty"`
	UmpDebut   string  `json:"UmpDebut,omitempty"`
}

// GetPersonFields get the fieldnamesf for an instance of Person
// can be used in database queries
func GetPersonFields() []string {
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

// LoadPersonnel - reads the retrosheet personnel json file and returns a slice of 
// Person's
func LoadPersonnel(fname string) ([]Person, error) {
	jsonBlob, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	var ids []Person

	err = json.Unmarshal(jsonBlob, &ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
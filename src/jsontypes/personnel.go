package jsontypes

import (
	"encoding/json"
	"io/ioutil"
)

// Person - one instance of the data from a personnel json file
type Person struct {
	ID             string `json:"ID,omitempty"`
	Last           string `json:"Last,omitempty"`
	First          string `json:"First,omitempty"`
	PlayerDebut    string `json:"PlayerDebut,omitempty"`
	ManagerDebut   string `json:"ManagerDebut,omitempty"`
	CoachDebut     string `json:"CoachDebut,omitempty"`
	UmpireDebut    string `json:"UmpireDebut,omitempty"`
}

// GetPersonFields get the fieldnamesf for an instance of Person
// can be used in database queries
func GetPersonFields() []string {
	return []string {
		"ID",
		"Last",
		"First",
		"PlayerDebut",
		"ManagerDebut",
		"CoachDebut",
		"UmpireDebut",
	  }
}

// ReadPersonnel - reads the retrosheet personnel json file and returns a slice of 
// Person's
func ReadPersonnel(fname string) ([]Person, error) {
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
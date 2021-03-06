package json_interface

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// The GO struct to be populated by unmarshal
type team struct {
	Abbr        string `json:"Abbr,omitempty"`
	League      string `json:"League,omitempty"`
	City        string `json:"City,omitempty"`
	Nickname    string `json:"Nickname,omitempty"`
	FirstYear   string `json:"FirstYear,omitempty"`
	LastYear    string `json:"LastYear,omitempty"`
}


func unmarshalTeam(line []byte) (team, error) {
	var g team
	err := json.Unmarshal(line, &g)
	fmt.Println(g)
	return g, err

}

func loadTeams(fname string) ([]team, error) {
	jsonBlob, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	var teams []team

	err = json.Unmarshal(jsonBlob, &teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

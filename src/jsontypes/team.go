package jsontypes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Team - The GO struct to be populated by unmarshal
type Team struct {
	Abbr        string `json:"Abbr,omitempty"`
	League      string `json:"League,omitempty"`
	City        string `json:"City,omitempty"`
	Nickname    string `json:"Nickname,omitempty"`
	FirstYear   string `json:"FirstYear,omitempty"`
	LastYear    string `json:"LastYear,omitempty"`
}

func unmarshalTeam(line []byte) (Team, error) {
	var g Team
	err := json.Unmarshal(line, &g)
	fmt.Println(g)
	return g, err

}

// LoadTeams - load team data from the specified file
// returns a slice of Team data or an error
func LoadTeams(fname string) ([]Team, error) {
	jsonBlob, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	var teams []Team

	err = json.Unmarshal(jsonBlob, &teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

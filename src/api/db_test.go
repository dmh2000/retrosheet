package api

import (
	"testing"
)

var Mongodb = "mongodb://172.24.32.1:27017" //  os.Getenv("RETROSHEET_MONGO")

func TestTeamByNickname(t *testing.T) {
	t.Log("TestTeamByNickname")
	t.Log(Mongodb)
	team, err := GetTeamByNickname(Mongodb, "Orioles")
	if err != nil {
		t.Errorf("GetTeamByNickname failed with error: %s", err)
	}
	t.Log(team)
}

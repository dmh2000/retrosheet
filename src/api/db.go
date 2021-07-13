package api

import (
	"fmt"

	jt "github.com/dmh2000/retrosheet/src/jsontypes"
	mq "github.com/dmh2000/retrosheet/src/query"
)

func GetTeamByNickname(uri string, nickname string) ([]jt.Team, error) {
	t, err := mq.GetTeamByKey(uri, "nickname", nickname)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return t, nil
}

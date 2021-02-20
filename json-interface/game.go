package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type game struct {
	Date                     string `json:",omitempty"`
	GameNumber               string `json:",omitempty"`
	DayOfWeek                string `json:",omitempty"`
	VisitorTeam              string `json:",omitempty"`
	VisitorLeague            string `json:",omitempty"`
	VisitorGameNumber        int `json:",omitempty"`
	HomeTeam                 string `json:",omitempty"`
	HomeLeague               string `json:",omitempty"`
	HomeGameNumber           int `json:",omitempty"`
	VisitorScore             int `json:",omitempty"`
	HomeScore                int `json:",omitempty"`
	LengthInOuts             int `json:",omitempty"`
	DayOrNight               string `json:",omitempty"`
	Completion               string `json:",omitempty"`
	Forfeit                  string `json:",omitempty"`
	Protest                  string `json:",omitempty"`
	ParkID                   string `json:",omitempty"`
	Attendance               int `json:",omitempty"`
	Duration                 int `json:",omitempty"`
	VisitorLine              string `json:",omitempty"`
	HomeLine                 string `json:",omitempty"`
	VisitorAtBats            int `json:",omitempty"`
	VisitorHits              int `json:",omitempty"`
	VisitorDbls              int `json:",omitempty"`
	VisitorTriples           int `json:",omitempty"`
	VisitorHomers            int `json:",omitempty"`
	VisitorRBI               int `json:",omitempty"`
	VisitorSacHits           int `json:",omitempty"`
	VisitorSacFlies          int `json:",omitempty"`
	VisitorHitByPitch        int `json:",omitempty"`
	VisitorWalks             int `json:",omitempty"`
	VisitorIntentionalWalks  int `json:",omitempty"`
	VisitorStrikeouts        int `json:",omitempty"`
	VisitorSteals            int `json:",omitempty"`
	VisitorStealCaught       int `json:",omitempty"`
	VisitorGrounded          int `json:",omitempty"`
	VisitorCatcherIfc        int `json:",omitempty"`
	VisitorLeftOnBase        int `json:",omitempty"`
	VisitorPitchers          int `json:",omitempty"`
	VisitorIndiERs           int `json:",omitempty"`
	VisitorTeamERs           int `json:",omitempty"`
	VisitorWilds             int `json:",omitempty"`
	VisitorBalks             int `json:",omitempty"`
	VisitorPutouts           int `json:",omitempty"`
	VisitorAssists           int `json:",omitempty"`
	VisitorErrors            int `json:",omitempty"`
	VisitorPassedBalls       int `json:",omitempty"`
	VisitorDoublePlays       int `json:",omitempty"`
	VisitorTriplePlays       int `json:",omitempty"`
	HomeAtBats               int `json:",omitempty"`
	HomeHits                 int `json:",omitempty"`
	HomeDbls                 int `json:",omitempty"`
	HomeTriples              int `json:",omitempty"`
	HomeHomers               int `json:",omitempty"`
	HomeRBI                  int `json:",omitempty"`
	HomeSacHits              int `json:",omitempty"`
	HomeSacFlies             int `json:",omitempty"`
	HomeHitByPitch           int `json:",omitempty"`
	HomeWalks                int `json:",omitempty"`
	HomeIntentionalWalks     int `json:",omitempty"`
	HomeStrikeouts           int `json:",omitempty"`
	HomeSteals               int `json:",omitempty"`
	HomeStealCaught          int `json:",omitempty"`
	HomeGrounded             int `json:",omitempty"`
	HomeCatcherIfc           int `json:",omitempty"`
	HomeLeftOnBase           int `json:",omitempty"`
	HomePitchers             int `json:",omitempty"`
	HomeIndiERs              int `json:",omitempty"`
	HomeTeamERs              int `json:",omitempty"`
	HomeWilds                int `json:",omitempty"`
	HomeBalks                int `json:",omitempty"`
	HomePutouts              int `json:",omitempty"`
	HomeAssists              int `json:",omitempty"`
	HomeErrors               int `json:",omitempty"`
	HomePassedBalls          int `json:",omitempty"`
	HomedoublePlays          int `json:",omitempty"`
	HomeTriplePlays          int `json:",omitempty"`
	UmpireHomeID             string `json:",omitempty"`
	UmpireHomeName           string `json:",omitempty"`
	UmpFirstBaseID           string `json:",omitempty"`
	UmpFirstBaseName         string `json:",omitempty"`
	UmpSecondBaseID          string `json:",omitempty"`
	UmpSecondBaseName        string `json:",omitempty"`
	UmpThirdBaseID           string `json:",omitempty"`
	UmpThirdBaseName         string `json:",omitempty"`
	UmpLeftFieldID           string `json:",omitempty"`
	UmpLeftFieldName         string `json:",omitempty"`
	UmpRightFieldID          string `json:",omitempty"`
	UmpRightFieldName        string `json:",omitempty"`
	VisitorManagerID         string `json:",omitempty"`
	VisitorManagerName       string `json:",omitempty"`
	HomeManagerID            string `json:",omitempty"`
	HomeMangerName           string `json:",omitempty"`
	WinningPitcherID         string `json:",omitempty"`
	WinningPitcherName       string `json:",omitempty"`
	LosingPitcherID          string `json:",omitempty"`
	LosingPitcherName        string `json:",omitempty"`
	SavingPitcherID          string `json:",omitempty"`
	SavingPitcherName        string `json:",omitempty"`
	RbiBatterID              string `json:",omitempty"`
	RbiBatterName            string `json:",omitempty"`
	VisitorStarterID         string `json:",omitempty"`
	VisitorStarterName       string `json:",omitempty"`
	HomeStarterID            string `json:",omitempty"`
	HomeStarterName          string `json:",omitempty"`
	VisitorLineup1ID         string `json:",omitempty"`
	VisitorLineup1Name       string `json:",omitempty"`
	VisitorLineup1Pos        int `json:",omitempty"`
	VisitorLineup2ID         string `json:",omitempty"`
	VisitorLineup2Name       string `json:",omitempty"`
	VisitorLineup2Pos        int `json:",omitempty"`
	VisitorLineup3ID         string `json:",omitempty"`
	VisitorLineup3Name       string `json:",omitempty"`
	VisitorLineup3Pos        int `json:",omitempty"`
	VisitorLineup4ID         string `json:",omitempty"`
	VisitorLineup4Name       string `json:",omitempty"`
	VisitorLineup4Pos        int `json:",omitempty"`
	VisitorLineup5ID         string `json:",omitempty"`
	VisitorLineup5Name       string `json:",omitempty"`
	VisitorLineup5Pos        int `json:",omitempty"`
	VisitorLineup6ID         string `json:",omitempty"`
	VisitorLineup6Name       string `json:",omitempty"`
	VisitorLineup6Pos        int `json:",omitempty"`
	VisitorLineup7ID         string `json:",omitempty"`
	VisitorLineup7Name       string `json:",omitempty"`
	VisitorLineup7Pos        int `json:",omitempty"`
	VisitorLineup8ID         string `json:",omitempty"`
	VisitorLineup8Name       string `json:",omitempty"`
	VisitorLineup8Pos        int `json:",omitempty"`
	VisitorLineup9ID         string `json:",omitempty"`
	VisitorLineup9Name       string `json:",omitempty"`
	VisitorLineup9Pos        int `json:",omitempty"`
	HomeLineup1ID            string `json:",omitempty"`
	HomeLineup1Name          string `json:",omitempty"`
	HomeLineup1Pos           int `json:",omitempty"`
	HomeLineup2ID            string `json:",omitempty"`
	HomeLineup2Name          string `json:",omitempty"`
	HomeLineup2Pos           int `json:",omitempty"`
	HomeLineup3ID            string `json:",omitempty"`
	HomeLineup3Name          string `json:",omitempty"`
	HomeLineup3Pos           int `json:",omitempty"`
	HomeLineup4ID            string `json:",omitempty"`
	HomeLineup4Name          string `json:",omitempty"`
	HomeLineup4Pos           int `json:",omitempty"`
	HomeLineup5ID            string `json:",omitempty"`
	HomeLineup5Name          string `json:",omitempty"`
	HomeLineup5Pos           int `json:",omitempty"`
	HomeLineup6ID            string `json:",omitempty"`
	HomeLineup6Name          string `json:",omitempty"`
	HomeLineup6Pos           int `json:",omitempty"`
	HomeLineup7ID            string `json:",omitempty"`
	HomeLineup7Name          string `json:",omitempty"`
	HomeLineup7Pos           int `json:",omitempty"`
	HomeLineup8ID            string `json:",omitempty"`
	HomeLineup8Name          string `json:",omitempty"`
	HomeLineup8Pos           string `json:",omitempty"`
	HomeLineup9ID            string `json:",omitempty"`
	HomeLineup9Name          string `json:",omitempty"`
	HomeLineup9Pos           string `json:",omitempty"`
	AdditionalInfo           string `json:",omitempty"`
	AcquisitionInfo          string `json:",omitempty"`
}

var line = []byte(`{"date":"20100404","gameNumber":"0","dayOfWeek":"Sun","visitorTeam":"NYA","visitorLeague":"AL","visitorGameNumber":"1","homeTeam":"BOS","homeLeague":"AL","homeGameNumber":"1","visitorScore":"7","homeScore":"9","lengthInOuts":"51","dayOrNight":"N","completion":"","forfeit":"","protest":"","parkID":"BOS07","attendance":"37440","duration":"226","visitorLine":"020300200","homeLine":"01001331x","visitorAtBats":"37","visitorHits":"12","visitorDbls":"2","visitorTriples":"0","visitorHomers":"2","visitorRBI":"6","visitorSacHits":"0","visitorSacFlies":"0","visitorHitByPitch":"0","visitorWalks":"6","visitorIntentionalWalks":"0","visitorStrikeouts":"2","visitorSteals":"2","visitorStealCaught":"0","visitorGrounded":"2","visitorCatcherIfc":"0","visitorLeftOnBase":"9","visitorPitchers":"5","visitorIndiERs":"8","visitorTeamERs":"8","visitorWilds":"1","visitorBalks":"0","visitorPutouts":"24","visitorAssists":"9","visitorErrors":"1","visitorPassedBalls":"1","visitorDoublePlays":"1","visitorTriplePlays":"0","homeAtBats":"34","homeHits":"12","homeDbls":"3","homeTriples":"1","homeHomers":"1","homeRBI":"8","homeSacHits":"0","homeSacFlies":"1","homeHitByPitch":"0","homeWalks":"4","homeIntentionalWalks":"0","homeStrikeouts":"5","homeSteals":"0","homeStealCaught":"0","homeGrounded":"0","homeCatcherIfc":"0","homeLeftOnBase":"6","homePitchers":"6","homeIndiERs":"7","homeTeamERs":"7","homeWilds":"1","homeBalks":"0","homePutouts":"27","homeAssists":"15","homeErrors":"0","homePassedBalls":"0","homedoublePlays":"2","homeTriplePlays":"0","umpireHomeID":"westj901","umpireHomeName":"Joe West","umpFirstBaseID":"herna901","umpFirstBaseName":"Angel Hernandez","umpSecondBaseID":"schrp901","umpSecondBaseName":"Paul Schrieber","umpThirdBaseID":"drakr901","umpThirdBaseName":"Rob Drake","umpLeftFieldID":"","umpLeftFieldName":"(none)","umpRightFieldID":"","umpRightFieldName":"(none)","visitorManagerID":"giraj001","visitorManagerName":"Joe Girardi","homeManagerID":"frant001","homeMangerName":"Terry Francona","winningPitcherID":"okajh001","winningPitcherName":"HIDeki Okajima","losingPitcherID":"parkc002","losingPitcherName":"Chan Ho Park","savingPitcherID":"papej001","savingPitcherName":"Jonathan Papelbon","rbiBatterID":"","rbiBatterName":"(none)","visitorStarterID":"sabac001","visitorStarterName":"CC Sabathia","homeStarterID":"beckj002","homeStarterName":"Josh Beckett","visitorLineup1ID":"jeted001","visitorLineup1Name":"Derek Jeter","visitorLineup1Pos":"6","visitorLineup2ID":"johnn001","visitorLineup2Name":"Nick Johnson","visitorLineup2Pos":"10","visitorLineup3ID":"teixm001","visitorLineup3Name":"Mark Teixeira","visitorLineup3Pos":"3","visitorLineup4ID":"rodra001","visitorLineup4Name":"Alex Rodriguez","visitorLineup4Pos":"5","visitorLineup5ID":"canor001","visitorLineup5Name":"Robinson Cano","visitorLineup5Pos":"4","visitorLineup6ID":"posaj001","visitorLineup6Name":"Jorge Posada","visitorLineup6Pos":"2","visitorLineup7ID":"granc001","visitorLineup7Name":"Curtis Granderson","visitorLineup7Pos":"8","visitorLineup8ID":"swisn001","visitorLineup8Name":"Nick Swisher","visitorLineup8Pos":"9","visitorLineup9ID":"gardb001","visitorLineup9Name":"Brett Gardner","visitorLineup9Pos":"7","homeLineup1ID":"ellsj001","homeLineup1Name":"Jacoby Ellsbury","homeLineup1Pos":"7","homeLineup2ID":"pedrd001","homeLineup2Name":"Dustin Pedroia","homeLineup2Pos":"4","homeLineup3ID":"martv001","homeLineup3Name":"Victor Martinez","homeLineup3Pos":"2","homeLineup4ID":"youkk001","homeLineup4Name":"Kevin Youkilis","homeLineup4Pos":"3","homeLineup5ID":"ortID001","homeLineup5Name":"DavID Ortiz","homeLineup5Pos":"10","homeLineup6ID":"belta001","homeLineup6Name":"Adrian Beltre","homeLineup6Pos":"5","homeLineup7ID":"drewj001","homeLineup7Name":"J.D. Drew","homeLineup7Pos":"9","homeLineup8ID":"camem001","homeLineup8Name":"Mike Cameron","homeLineup8Pos":"8","homeLineup9ID":"scutm001","homeLineup9Name":"Marco Scutaro","homeLineup9Pos":"6","additionalInfo":"","acquisitionInfo":"Y"}`)

func unmarshalGame(line []byte) (game, error) {
	var g game
	err := json.Unmarshal(line, &g)
	fmt.Println(g)
	return g, err

}

func loadGames(fname string) ([]game, error) {
	jsonBlob, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	var games []game

	err = json.Unmarshal(jsonBlob, &games)
	if err != nil {
		return nil, err
	}

	return games, nil
}


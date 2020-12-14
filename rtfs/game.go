package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type game1 struct {
	Date                    string
	GameNumber              string
	DayOfWeek               string
	VisitorTeam             string
	VisitorLeague           string
	VisitorGameNumber       string
	HomeTeam                string
	HomeLeague              string
	HomeGameNumber          string
	VisitorScore            string
	HomeScore               string
	LengthInOuts            string
	DayOrNight              string
	Completion              string
	Forfeit                 string
	Protest                 string
	ParkID                  string
	Attendance              string
	Duration                string
	VisitorLine             string
	HomeLine                string
	VisitorAtBats           string
	VisitorHits             string
	VisitorDbls             string
	VisitorTriples          string
	VisitorHomers           string
	VisitorRBI              string
	VisitorSacHits          string
	VisitorSacFlies         string
	VisitorHitByPitch       string
	VisitorWalks            string
	VisitorIntentionalWalks string
	VisitorStrikeouts       string
	VisitorSteals           string
	VisitorStealCaught      string
	VisitorGrounded         string
	VisitorCatcherIfc       string
	VisitorLeftOnBase       string
	VisitorPitchers         string
	VisitorIndiERs          string
	VisitorTeamERs          string
	VisitorWilds            string
	VisitorBalks            string
	visitorPutouts          string
	VisitorAssists          string
	VisitorErrors           string
	VisitorPassedBalls      string
	VisitorDoublePlays      string
	VisitorTriplePlays      string
	HomeAtBats              string
	HomeHits                string
	HomeDbls                string
	HomeTriples             string
	HomeHomers              string
	HomeRBI                 string
	HomeSacHits             string
	HomeSacFlies            string
	HomeHitByPitch          string
	HomeWalks               string
	HomeIntentionalWalks    string
	HomeStrikeouts          string
	HomeSteals              string
	HomeStealCaught         string
	HomeGrounded            string
	HomeCatcherIfc          string
	HomeLeftOnBase          string
	HomePitchers            string
	HomeIndiERs             string
	HomeTeamERs             string
	HomeWilds               string
	HomeBalks               string
	HomePutouts             string
	HomeAssists             string
	HomeErrors              string
	HomePassedBalls         string
	HomedoublePlays         string
	HomeTriplePlays         string
	UmpireHomeID            string
	UmpireHomeName          string
	UmpFirstBaseID          string
	UmpFirstBaseName        string
	UmpSecondBaseID         string
	UmpSecondBaseName       string
	UmpThirdBaseID          string
	UmpThirdBaseName        string
	UmpLeftFieldID          string
	UmpLeftFieldName        string
	UmpRightFieldID         string
	UmpRightFieldName       string
	VisitorManagerID        string
	VisitorManagerName      string
	HomeManagerID           string
	HomeMangerName          string
	WinningPitcherID        string
	WinningPitcherName      string
	LosingPitcherID         string
	LosingPitcherName       string
	SavingPitcherID         string
	SavingPitcherName       string
	RbiBatterID             string
	RbiBatterName           string
	VisitorStarterID        string
	VisitorStarterName      string
	HomeStarterID           string
	HomeStarterName         string
	VisitorLineup1ID        string
	VisitorLineup1Name      string
	VisitorLineup1Pos       string
	VisitorLineup2ID        string
	VisitorLineup2Name      string
	VisitorLineup2Pos       string
	VisitorLineup3ID        string
	VisitorLineup3Name      string
	VisitorLineup3Pos       string
	VisitorLineup4ID        string
	VisitorLineup4Name      string
	VisitorLineup4Pos       string
	VisitorLineup5ID        string
	VisitorLineup5Name      string
	VisitorLineup5Pos       string
	VisitorLineup6ID        string
	VisitorLineup6Name      string
	VisitorLineup6Pos       string
	VisitorLineup7ID        string
	VisitorLineup7Name      string
	VisitorLineup7Pos       string
	VisitorLineup8ID        string
	VisitorLineup8Name      string
	VisitorLineup8Pos       string
	VisitorLineup9ID        string
	VisitorLineup9Name      string
	VisitorLineup9Pos       string
	HomeLineup1ID           string
	HomeLineup1Name         string
	HomeLineup1Pos          string
	HomeLineup2ID           string
	HomeLineup2Name         string
	HomeLineup2Pos          string
	HomeLineup3ID           string
	HomeLineup3Name         string
	HomeLineup3Pos          string
	HomeLineup4ID           string
	HomeLineup4Name         string
	HomeLineup4Pos          string
	HomeLineup5ID           string
	HomeLineup5Name         string
	HomeLineup5Pos          string
	HomeLineup6ID           string
	HomeLineup6Name         string
	HomeLineup6Pos          string
	HomeLineup7ID           string
	HomeLineup7Name         string
	HomeLineup7Pos          string
	HomeLineup8ID           string
	HomeLineup8Name         string
	HomeLineup8Pos          string
	HomeLineup9ID           string
	HomeLineup9Name         string
	HomeLineup9Pos          string
	AdditionalInfo          string
	AcquisitionInfo         string
}

type game struct {
	Date                    string
	GameNumber              string
	DayOfWeek               string
	VisitorTeam             string
	VisitorLeague           string
	VisitorGameNumber       int
	HomeTeam                string
	HomeLeague              string
	HomeGameNumber          int
	VisitorScore            int
	HomeScore               int
	LengthInOuts            int
	DayOrNight              string
	Completion              string
	Forfeit                 string
	Protest                 string
	ParkID                  string
	Attendance              int
	Duration                int
	VisitorLine             string
	HomeLine                string
	VisitorAtBats           int
	VisitorHits             int
	VisitorDbls             int
	VisitorTriples          int
	VisitorHomers           int
	VisitorRBI              int
	VisitorSacHits          int
	VisitorSacFlies         int
	VisitorHitByPitch       int
	VisitorWalks            int
	VisitorIntentionalWalks int
	VisitorStrikeouts       int
	VisitorSteals           int
	VisitorStealCaught      int
	VisitorGrounded         int
	VisitorCatcherIfc       int
	VisitorLeftOnBase       int
	VisitorPitchers         int
	VisitorIndiERs          int
	VisitorTeamERs          int
	VisitorWilds            int
	VisitorBalks            int
	visitorPutouts          int
	VisitorAssists          int
	VisitorErrors           int
	VisitorPassedBalls      int
	VisitorDoublePlays      int
	VisitorTriplePlays      int
	HomeAtBats              int
	HomeHits                int
	HomeDbls                int
	HomeTriples             int
	HomeHomers              int
	HomeRBI                 int
	HomeSacHits             int
	HomeSacFlies            int
	HomeHitByPitch          int
	HomeWalks               int
	HomeIntentionalWalks    int
	HomeStrikeouts          int
	HomeSteals              int
	HomeStealCaught         int
	HomeGrounded            int
	HomeCatcherIfc          int
	HomeLeftOnBase          int
	HomePitchers            int
	HomeIndiERs             int
	HomeTeamERs             int
	HomeWilds               int
	HomeBalks               int
	HomePutouts             int
	HomeAssists             int
	HomeErrors              int
	HomePassedBalls         int
	HomedoublePlays         int
	HomeTriplePlays         int
	UmpireHomeID            string
	UmpireHomeName          string
	UmpFirstBaseID          string
	UmpFirstBaseName        string
	UmpSecondBaseID         string
	UmpSecondBaseName       string
	UmpThirdBaseID          string
	UmpThirdBaseName        string
	UmpLeftFieldID          string
	UmpLeftFieldName        string
	UmpRightFieldID         string
	UmpRightFieldName       string
	VisitorManagerID        string
	VisitorManagerName      string
	HomeManagerID           string
	HomeMangerName          string
	WinningPitcherID        string
	WinningPitcherName      string
	LosingPitcherID         string
	LosingPitcherName       string
	SavingPitcherID         string
	SavingPitcherName       string
	RbiBatterID             string
	RbiBatterName           string
	VisitorStarterID        string
	VisitorStarterName      string
	HomeStarterID           string
	HomeStarterName         string
	VisitorLineup1ID        string
	VisitorLineup1Name      string
	VisitorLineup1Pos       int
	VisitorLineup2ID        string
	VisitorLineup2Name      string
	VisitorLineup2Pos       int
	VisitorLineup3ID        string
	VisitorLineup3Name      string
	VisitorLineup3Pos       int
	VisitorLineup4ID        string
	VisitorLineup4Name      string
	VisitorLineup4Pos       int
	VisitorLineup5ID        string
	VisitorLineup5Name      string
	VisitorLineup5Pos       int
	VisitorLineup6ID        string
	VisitorLineup6Name      string
	VisitorLineup6Pos       int
	VisitorLineup7ID        string
	VisitorLineup7Name      string
	VisitorLineup7Pos       int
	VisitorLineup8ID        string
	VisitorLineup8Name      string
	VisitorLineup8Pos       int
	VisitorLineup9ID        string
	VisitorLineup9Name      string
	VisitorLineup9Pos       int
	HomeLineup1ID           string
	HomeLineup1Name         string
	HomeLineup1Pos          int
	HomeLineup2ID           string
	HomeLineup2Name         string
	HomeLineup2Pos          int
	HomeLineup3ID           string
	HomeLineup3Name         string
	HomeLineup3Pos          int
	HomeLineup4ID           string
	HomeLineup4Name         string
	HomeLineup4Pos          int
	HomeLineup5ID           string
	HomeLineup5Name         string
	HomeLineup5Pos          int
	HomeLineup6ID           string
	HomeLineup6Name         string
	HomeLineup6Pos          int
	HomeLineup7ID           string
	HomeLineup7Name         string
	HomeLineup7Pos          int
	HomeLineup8ID           string
	HomeLineup8Name         string
	HomeLineup8Pos          string
	HomeLineup9ID           string
	HomeLineup9Name         string
	HomeLineup9Pos          string
	AdditionalInfo          string
	AcquisitionInfo         string
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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// The GO struct to be populated by unmarshal
type game struct {
	Date                     string `json:"date,omitempty"`
	GameNumber               string `json:"gameNumber,omitempty"`
	DayOfWeek                string `json:"dayOfWeek,omitempty"`
	VisitorTeam              string `json:"visitorTeam,omitempty"`
	VisitorLeague            string `json:"visitorLeague,omitempty"`
	VisitorGameNumber        int `json:"visitorGameNumber,omitempty"`
	HomeTeam                 string `json:"homeTeam,omitempty"`
	HomeLeague               string `json:"homeLeague,omitempty"`
	HomeGameNumber           int `json:"homeGameNumber,omitempty"`
	VisitorScore             int `json:"visitorScore,omitempty"`
	HomeScore                int `json:"homeScore,omitempty"`
	LengthInOuts             int `json:"lengthInOuts,omitempty"`
	DayOrNight               string `json:"dayOrNight,omitempty"`
	Completion               string `json:"completion,omitempty"`
	Forfeit                  string `json:"forfeit,omitempty"`
	Protest                  string `json:"protest,omitempty"`
	ParkID                   string `json:"parkId,omitempty"`
	Attendance               int `json:"attendance,omitempty"`
	Duration                 int `json:"duration,omitempty"`
	VisitorLine              string `json:"visitorLine,omitempty"`
	HomeLine                 string `json:"homeLine,omitempty"`
	VisitorAtBats            int `json:"visitorAtBats,omitempty"`
	VisitorHits              int `json:"visitorHits,omitempty"`
	VisitorDbls              int `json:"visitorDbls,omitempty"`
	VisitorTriples           int `json:"visitorTriples,omitempty"`
	VisitorHomers            int `json:"visitorHomers,omitempty"`
	VisitorRBI               int `json:"visitorRBI,omitempty"`
	VisitorSacHits           int `json:"visitorSacHits,omitempty"`
	VisitorSacFlies          int `json:"visitorSacFlies,omitempty"`
	VisitorHitByPitch        int `json:"visitorHitByPitch,omitempty"`
	VisitorWalks             int `json:"visitorWalks,omitempty"`
	VisitorIntentionalWalks  int `json:"visitorIntentionalWalks,omitempty"`
	VisitorStrikeouts        int `json:"visitorStrikeouts,omitempty"`
	VisitorSteals            int `json:"visitorSteals,omitempty"`
	VisitorStealCaught       int `json:"visitorStealCaught,omitempty"`
	VisitorGrounded          int `json:"visitorGrounded,omitempty"`
	VisitorCatcherIfc        int `json:"visitorCatcherIfc,omitempty"`
	VisitorLeftOnBase        int `json:"visitorLeftOnBase,omitempty"`
	VisitorPitchers          int `json:"visitorPitchers,omitempty"`
	VisitorIndiERs           int `json:"visitorIndiERs,omitempty"`
	VisitorTeamERs           int `json:"visitorTeamERs,omitempty"`
	VisitorWilds             int `json:"visitorWilds,omitempty"`
	VisitorBalks             int `json:"visitorBalks,omitempty"`
	VisitorPutouts           int `json:"visitorPutouts,omitempty"`
	VisitorAssists           int `json:"visitorAssists,omitempty"`
	VisitorErrors            int `json:"visitorErrors,omitempty"`
	VisitorPassedBalls       int `json:"visitorPassedBalls,omitempty"`
	VisitorDoublePlays       int `json:"visitorDoublePlays,omitempty"`
	VisitorTriplePlays       int `json:"visitorTriplePlays,omitempty"`
	HomeAtBats               int `json:"homeAtBats,omitempty"`
	HomeHits                 int `json:"homeHits,omitempty"`
	HomeDbls                 int `json:"homeDbls,omitempty"`
	HomeTriples              int `json:"homeTriples,omitempty"`
	HomeHomers               int `json:"homeHomers,omitempty"`
	HomeRBI                  int `json:"homeRBI,omitempty"`
	HomeSacHits              int `json:"homeSacHits,omitempty"`
	HomeSacFlies             int `json:"homeSacFlies,omitempty"`
	HomeHitByPitch           int `json:"homeHitByPitch,omitempty"`
	HomeWalks                int `json:"homeWalks,omitempty"`
	HomeIntentionalWalks     int `json:"homeIntentionalWalks,omitempty"`
	HomeStrikeouts           int `json:"homeStrikeouts,omitempty"`
	HomeSteals               int `json:"homeSteals,omitempty"`
	HomeStealCaught          int `json:"homeStealCaught,omitempty"`
	HomeGrounded             int `json:"homeGrounded,omitempty"`
	HomeCatcherIfc           int `json:"homeCatcherIfc,omitempty"`
	HomeLeftOnBase           int `json:"homeLeftOnBase,omitempty"`
	HomePitchers             int `json:"homePitchers,omitempty"`
	HomeIndiERs              int `json:"homeIndiERs,omitempty"`
	HomeTeamERs              int `json:"homeTeamERs,omitempty"`
	HomeWilds                int `json:"homeWilds,omitempty"`
	HomeBalks                int `json:"homeBalks,omitempty"`
	HomePutouts              int `json:"homePutouts,omitempty"`
	HomeAssists              int `json:"homeAssists,omitempty"`
	HomeErrors               int `json:"homeErrors,omitempty"`
	HomePassedBalls          int `json:"homePassedBalls,omitempty"`
	HomedoublePlays          int `json:"homedoublePlays,omitempty"`
	HomeTriplePlays          int `json:"homeTriplePlays,omitempty"`
	UmpireHomeID             string `json:"umpireHomeId,omitempty"`
	UmpireHomeName           string `json:"umpireHomeName,omitempty"`
	UmpFirstBaseID           string `json:"umpFirstBaseId,omitempty"`
	UmpFirstBaseName         string `json:"umpFirstBaseName,omitempty"`
	UmpSecondBaseID          string `json:"umpSecondBaseId,omitempty"`
	UmpSecondBaseName        string `json:"umpSecondBaseName,omitempty"`
	UmpThirdBaseID           string `json:"umpThirdBaseId,omitempty"`
	UmpThirdBaseName         string `json:"umpThirdBaseName,omitempty"`
	UmpLeftFieldID           string `json:"umpLeftFieldId,omitempty"`
	UmpLeftFieldName         string `json:"umpLeftFieldName,omitempty"`
	UmpRightFieldID          string `json:"umpRightFieldId,omitempty"`
	UmpRightFieldName        string `json:"umpRightFieldName,omitempty"`
	VisitorManagerID         string `json:"visitorManagerId,omitempty"`
	VisitorManagerName       string `json:"visitorManagerName,omitempty"`
	HomeManagerID            string `json:"homeManagerId,omitempty"`
	HomeMangerName           string `json:"homeMangerName,omitempty"`
	WinningPitcherID         string `json:"winningPitcherId,omitempty"`
	WinningPitcherName       string `json:"winningPitcherName,omitempty"`
	LosingPitcherID          string `json:"losingPitcherId,omitempty"`
	LosingPitcherName        string `json:"losingPitcherName,omitempty"`
	SavingPitcherID          string `json:"savingPitcherId,omitempty"`
	SavingPitcherName        string `json:"savingPitcherName,omitempty"`
	RbiBatterID              string `json:"rbiBatterId,omitempty"`
	RbiBatterName            string `json:"rbiBatterName,omitempty"`
	VisitorStarterID         string `json:"visitorStarterId,omitempty"`
	VisitorStarterName       string `json:"visitorStarterName,omitempty"`
	HomeStarterID            string `json:"homeStarterId,omitempty"`
	HomeStarterName          string `json:"homeStarterName,omitempty"`
	VisitorLineup1ID         string `json:"visitorLineup1Id,omitempty"`
	VisitorLineup1Name       string `json:"visitorLineup1Name,omitempty"`
	VisitorLineup1Pos        int `json:"visitorLineup1Pos,omitempty"`
	VisitorLineup2ID         string `json:"visitorLineup2Id,omitempty"`
	VisitorLineup2Name       string `json:"visitorLineup2Name,omitempty"`
	VisitorLineup2Pos        int `json:"visitorLineup2Pos,omitempty"`
	VisitorLineup3ID         string `json:"visitorLineup3Id,omitempty"`
	VisitorLineup3Name       string `json:"visitorLineup3Name,omitempty"`
	VisitorLineup3Pos        int `json:"visitorLineup3Pos,omitempty"`
	VisitorLineup4ID         string `json:"visitorLineup4Id,omitempty"`
	VisitorLineup4Name       string `json:"visitorLineup4Name,omitempty"`
	VisitorLineup4Pos        int `json:"visitorLineup4Pos,omitempty"`
	VisitorLineup5ID         string `json:"visitorLineup5Id,omitempty"`
	VisitorLineup5Name       string `json:"visitorLineup5Name,omitempty"`
	VisitorLineup5Pos        int `json:"visitorLineup5Pos,omitempty"`
	VisitorLineup6ID         string `json:"visitorLineup6Id,omitempty"`
	VisitorLineup6Name       string `json:"visitorLineup6Name,omitempty"`
	VisitorLineup6Pos        int `json:"visitorLineup6Pos,omitempty"`
	VisitorLineup7ID         string `json:"visitorLineup7Id,omitempty"`
	VisitorLineup7Name       string `json:"visitorLineup7Name,omitempty"`
	VisitorLineup7Pos        int `json:"visitorLineup7Pos,omitempty"`
	VisitorLineup8ID         string `json:"visitorLineup8Id,omitempty"`
	VisitorLineup8Name       string `json:"visitorLineup8Name,omitempty"`
	VisitorLineup8Pos        int `json:"visitorLineup8Pos,omitempty"`
	VisitorLineup9ID         string `json:"visitorLineup9Id,omitempty"`
	VisitorLineup9Name       string `json:"visitorLineup9Name,omitempty"`
	VisitorLineup9Pos        int `json:"visitorLineup9Pos,omitempty"`
	HomeLineup1ID            string `json:"homeLineup1Id,omitempty"`
	HomeLineup1Name          string `json:"homeLineup1Name,omitempty"`
	HomeLineup1Pos           int `json:"homeLineup1Pos,omitempty"`
	HomeLineup2ID            string `json:"homeLineup2Id,omitempty"`
	HomeLineup2Name          string `json:"homeLineup2Name,omitempty"`
	HomeLineup2Pos           int `json:"homeLineup2Pos,omitempty"`
	HomeLineup3ID            string `json:"homeLineup3Id,omitempty"`
	HomeLineup3Name          string `json:"homeLineup3Name,omitempty"`
	HomeLineup3Pos           int `json:"homeLineup3Pos,omitempty"`
	HomeLineup4ID            string `json:"homeLineup4Id,omitempty"`
	HomeLineup4Name          string `json:"homeLineup4Name,omitempty"`
	HomeLineup4Pos           int `json:"homeLineup4Pos,omitempty"`
	HomeLineup5ID            string `json:"homeLineup5Id,omitempty"`
	HomeLineup5Name          string `json:"homeLineup5Name,omitempty"`
	HomeLineup5Pos           int `json:"homeLineup5Pos,omitempty"`
	HomeLineup6ID            string `json:"homeLineup6Id,omitempty"`
	HomeLineup6Name          string `json:"homeLineup6Name,omitempty"`
	HomeLineup6Pos           int `json:"homeLineup6Pos,omitempty"`
	HomeLineup7ID            string `json:"homeLineup7Id,omitempty"`
	HomeLineup7Name          string `json:"homeLineup7Name,omitempty"`
	HomeLineup7Pos           int `json:"homeLineup7Pos,omitempty"`
	HomeLineup8ID            string `json:"homeLineup8Id,omitempty"`
	HomeLineup8Name          string `json:"homeLineup8Name,omitempty"`
	HomeLineup8Pos           string `json:"homeLineup8Pos,omitempty"`
	HomeLineup9ID            string `json:"homeLineup9Id,omitempty"`
	HomeLineup9Name          string `json:"homeLineup9Name,omitempty"`
	HomeLineup9Pos           string `json:"homeLineup9Pos,omitempty"`
	AdditionalInfo           string `json:"additionalInfo,omitempty"`
	AcquisitionInfo          string `json:"acquisitionInfo,omitempty"`
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

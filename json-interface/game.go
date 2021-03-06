package json_interface

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// The GO struct to be populated by unmarshal
type game struct {
	Date                     string `json:"Date,omitempty"`
	GameNumber               string `json:"GameNumber,omitempty"`
	DayOfWeek                string `json:"DayOfWeek,omitempty"`
	VisitorTeam              string `json:"VisitorTeam,omitempty"`
	VisitorLeague            string `json:"VisitorLeague,omitempty"`
	VisitorGameNumber        int `json:"VisitorGameNumber,omitempty"`
	HomeTeam                 string `json:"HomeTeam,omitempty"`
	HomeLeague               string `json:"HomeLeague,omitempty"`
	HomeGameNumber           int `json:"HomeGameNumber,omitempty"`
	VisitorScore             int `json:"VisitorScore,omitempty"`
	HomeScore                int `json:"HomeScore,omitempty"`
	LengthInOuts             int `json:"LengthInOuts,omitempty"`
	DayOrNight               string `json:"DayOrNight,omitempty"`
	Completion               string `json:"Completion,omitempty"`
	Forfeit                  string `json:"Forfeit,omitempty"`
	Protest                  string `json:"Protest,omitempty"`
	ParkID                   string `json:"ParkId,omitempty"`
	Attendance               int `json:"Attendance,omitempty"`
	Duration                 int `json:"Duration,omitempty"`
	VisitorLine              string `json:"VisitorLine,omitempty"`
	HomeLine                 string `json:"HomeLine,omitempty"`
	VisitorAtBats            int `json:"VisitorAtBats,omitempty"`
	VisitorHits              int `json:"VisitorHits,omitempty"`
	VisitorDbls              int `json:"VisitorDbls,omitempty"`
	VisitorTriples           int `json:"VisitorTriples,omitempty"`
	VisitorHomers            int `json:"VisitorHomers,omitempty"`
	VisitorRBI               int `json:"VisitorRBI,omitempty"`
	VisitorSacHits           int `json:"VisitorSacHits,omitempty"`
	VisitorSacFlies          int `json:"VisitorSacFlies,omitempty"`
	VisitorHitByPitch        int `json:"VisitorHitByPitch,omitempty"`
	VisitorWalks             int `json:"VisitorWalks,omitempty"`
	VisitorIntentionalWalks  int `json:"VisitorIntentionalWalks,omitempty"`
	VisitorStrikeouts        int `json:"VisitorStrikeouts,omitempty"`
	VisitorSteals            int `json:"VisitorSteals,omitempty"`
	VisitorStealCaught       int `json:"VisitorStealCaught,omitempty"`
	VisitorGrounded          int `json:"VisitorGrounded,omitempty"`
	VisitorCatcherIfc        int `json:"VisitorCatcherIfc,omitempty"`
	VisitorLeftOnBase        int `json:"VisitorLeftOnBase,omitempty"`
	VisitorPitchers          int `json:"VisitorPitchers,omitempty"`
	VisitorIndiERs           int `json:"VisitorIndiERs,omitempty"`
	VisitorTeamERs           int `json:"VisitorTeamERs,omitempty"`
	VisitorWilds             int `json:"VisitorWilds,omitempty"`
	VisitorBalks             int `json:"VisitorBalks,omitempty"`
	VisitorPutouts           int `json:"VisitorPutouts,omitempty"`
	VisitorAssists           int `json:"VisitorAssists,omitempty"`
	VisitorErrors            int `json:"VisitorErrors,omitempty"`
	VisitorPassedBalls       int `json:"VisitorPassedBalls,omitempty"`
	VisitorDoublePlays       int `json:"VisitorDoublePlays,omitempty"`
	VisitorTriplePlays       int `json:"VisitorTriplePlays,omitempty"`
	HomeAtBats               int `json:"HomeAtBats,omitempty"`
	HomeHits                 int `json:"HomeHits,omitempty"`
	HomeDbls                 int `json:"HomeDbls,omitempty"`
	HomeTriples              int `json:"HomeTriples,omitempty"`
	HomeHomers               int `json:"HomeHomers,omitempty"`
	HomeRBI                  int `json:"HomeRBI,omitempty"`
	HomeSacHits              int `json:"HomeSacHits,omitempty"`
	HomeSacFlies             int `json:"HomeSacFlies,omitempty"`
	HomeHitByPitch           int `json:"HomeHitByPitch,omitempty"`
	HomeWalks                int `json:"HomeWalks,omitempty"`
	HomeIntentionalWalks     int `json:"HomeIntentionalWalks,omitempty"`
	HomeStrikeouts           int `json:"HomeStrikeouts,omitempty"`
	HomeSteals               int `json:"HomeSteals,omitempty"`
	HomeStealCaught          int `json:"HomeStealCaught,omitempty"`
	HomeGrounded             int `json:"HomeGrounded,omitempty"`
	HomeCatcherIfc           int `json:"HomeCatcherIfc,omitempty"`
	HomeLeftOnBase           int `json:"HomeLeftOnBase,omitempty"`
	HomePitchers             int `json:"HomePitchers,omitempty"`
	HomeIndiERs              int `json:"HomeIndiERs,omitempty"`
	HomeTeamERs              int `json:"HomeTeamERs,omitempty"`
	HomeWilds                int `json:"HomeWilds,omitempty"`
	HomeBalks                int `json:"HomeBalks,omitempty"`
	HomePutouts              int `json:"HomePutouts,omitempty"`
	HomeAssists              int `json:"HomeAssists,omitempty"`
	HomeErrors               int `json:"HomeErrors,omitempty"`
	HomePassedBalls          int `json:"HomePassedBalls,omitempty"`
	HomedoublePlays          int `json:"HomedoublePlays,omitempty"`
	HomeTriplePlays          int `json:"HomeTriplePlays,omitempty"`
	UmpireHomeID            string `json:"UmpireHomeId,omitempty"`
	UmpireHomeName           string `json:"UmpireHomeName,omitempty"`
	UmpFirstBaseID           string `json:"UmpFirstBaseId,omitempty"`
	UmpFirstBaseName         string `json:"UmpFirstBaseName,omitempty"`
	UmpSecondBaseID          string `json:"UmpSecondBaseId,omitempty"`
	UmpSecondBaseName        string `json:"UmpSecondBaseName,omitempty"`
	UmpThirdBaseID           string `json:"UmpThirdBaseId,omitempty"`
	UmpThirdBaseName         string `json:"UmpThirdBaseName,omitempty"`
	UmpLeftFieldID           string `json:"UmpLeftFieldId,omitempty"`
	UmpLeftFieldName         string `json:"UmpLeftFieldName,omitempty"`
	UmpRightFieldID          string `json:"UmpRightFieldId,omitempty"`
	UmpRightFieldName        string `json:"UmpRightFieldName,omitempty"`
	VisitorManagerID         string `json:"VisitorManagerId,omitempty"`
	VisitorManagerName       string `json:"VisitorManagerName,omitempty"`
	HomeManagerID            string `json:"HomeManagerId,omitempty"`
	HomeMangerName           string `json:"HomeMangerName,omitempty"`
	WinningPitcherID         string `json:"WinningPitcherId,omitempty"`
	WinningPitcherName       string `json:"WinningPitcherName,omitempty"`
	LosingPitcherID          string `json:"LosingPitcherId,omitempty"`
	LosingPitcherName        string `json:"LosingPitcherName,omitempty"`
	SavingPitcherID          string `json:"SavingPitcherId,omitempty"`
	SavingPitcherName        string `json:"SavingPitcherName,omitempty"`
	RbiBatterID              string `json:"RbiBatterId,omitempty"`
	RbiBatterName            string `json:"RbiBatterName,omitempty"`
	VisitorStarterID         string `json:"VisitorStarterId,omitempty"`
	VisitorStarterName       string `json:"VisitorStarterName,omitempty"`
	HomeStarterID            string `json:"HomeStarterId,omitempty"`
	HomeStarterName          string `json:"HomeStarterName,omitempty"`
	VisitorLineup1Id         string `json:"VisitorLineup1Id,omitempty"`
	VisitorLineup1Name       string `json:"VisitorLineup1Name,omitempty"`
	VisitorLineup1Pos        int `json:"VisitorLineup1Pos,omitempty"`
	VisitorLineup2Id         string `json:"VisitorLineup2Id,omitempty"`
	VisitorLineup2Name       string `json:"VisitorLineup2Name,omitempty"`
	VisitorLineup2Pos        int `json:"VisitorLineup2Pos,omitempty"`
	VisitorLineup3Id         string `json:"VisitorLineup3Id,omitempty"`
	VisitorLineup3Name       string `json:"VisitorLineup3Name,omitempty"`
	VisitorLineup3Pos        int `json:"VisitorLineup3Pos,omitempty"`
	VisitorLineup4Id         string `json:"VisitorLineup4Id,omitempty"`
	VisitorLineup4Name       string `json:"VisitorLineup4Name,omitempty"`
	VisitorLineup4Pos        int `json:"VisitorLineup4Pos,omitempty"`
	VisitorLineup5Id         string `json:"VisitorLineup5Id,omitempty"`
	VisitorLineup5Name       string `json:"VisitorLineup5Name,omitempty"`
	VisitorLineup5Pos        int `json:"VisitorLineup5Pos,omitempty"`
	VisitorLineup6Id         string `json:"VisitorLineup6Id,omitempty"`
	VisitorLineup6Name       string `json:"VisitorLineup6Name,omitempty"`
	VisitorLineup6Pos        int `json:"VisitorLineup6Pos,omitempty"`
	VisitorLineup7Id         string `json:"VisitorLineup7Id,omitempty"`
	VisitorLineup7Name       string `json:"VisitorLineup7Name,omitempty"`
	VisitorLineup7Pos        int `json:"VisitorLineup7Pos,omitempty"`
	VisitorLineup8Id         string `json:"VisitorLineup8Id,omitempty"`
	VisitorLineup8Name       string `json:"VisitorLineup8Name,omitempty"`
	VisitorLineup8Pos        int `json:"VisitorLineup8Pos,omitempty"`
	VisitorLineup9Id         string `json:"VisitorLineup9Id,omitempty"`
	VisitorLineup9Name       string `json:"VisitorLineup9Name,omitempty"`
	VisitorLineup9Pos        int `json:"VisitorLineup9Pos,omitempty"`
	HomeLineup1Id            string `json:"HomeLineup1Id,omitempty"`
	HomeLineup1Name          string `json:"HomeLineup1Name,omitempty"`
	HomeLineup1Pos           int `json:"HomeLineup1Pos,omitempty"`
	HomeLineup2Id            string `json:"HomeLineup2Id,omitempty"`
	HomeLineup2Name          string `json:"HomeLineup2Name,omitempty"`
	HomeLineup2Pos           int `json:"HomeLineup2Pos,omitempty"`
	HomeLineup3Id            string `json:"HomeLineup3Id,omitempty"`
	HomeLineup3Name          string `json:"HomeLineup3Name,omitempty"`
	HomeLineup3Pos           int `json:"HomeLineup3Pos,omitempty"`
	HomeLineup4Id            string `json:"HomeLineup4Id,omitempty"`
	HomeLineup4Name          string `json:"HomeLineup4Name,omitempty"`
	HomeLineup4Pos           int `json:"HomeLineup4Pos,omitempty"`
	HomeLineup5Id            string `json:"HomeLineup5Id,omitempty"`
	HomeLineup5Name          string `json:"HomeLineup5Name,omitempty"`
	HomeLineup5Pos           int `json:"HomeLineup5Pos,omitempty"`
	HomeLineup6Id            string `json:"HomeLineup6Id,omitempty"`
	HomeLineup6Name          string `json:"HomeLineup6Name,omitempty"`
	HomeLineup6Pos           int `json:"HomeLineup6Pos,omitempty"`
	HomeLineup7Id            string `json:"HomeLineup7Id,omitempty"`
	HomeLineup7Name          string `json:"HomeLineup7Name,omitempty"`
	HomeLineup7Pos           int `json:"HomeLineup7Pos,omitempty"`
	HomeLineup8Id            string `json:"HomeLineup8Id,omitempty"`
	HomeLineup8Name          string `json:"HomeLineup8Name,omitempty"`
	HomeLineup8Pos           string `json:"HomeLineup8Pos,omitempty"`
	HomeLineup9Id            string `json:"HomeLineup9Id,omitempty"`
	HomeLineup9Name          string `json:"HomeLineup9Name,omitempty"`
	HomeLineup9Pos           string `json:"HomeLineup9Pos,omitempty"`
	AdditionalInfo           string `json:"AdditionalInfo,omitempty"`
	AcquisitionInfo          string `json:"AcquisitionInfo,omitempty"`
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

func printGame(g game) {

}
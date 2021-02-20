package main

import "fmt"

// golang requires fieldnames in the interface def to be upper case
// when using json.Unmarshal
var glfields = []string{
	// Field(s)  Meaning
	"Date",       //    1     Date in the form"yyyymmdd"
	"GameNumber", //    2     Number of game:
	//   "0" -- a single game
	//   "1" -- the first game of a double (or triple) header
	//                     including seperate admission doubleheaders
	//   "2" -- the second game of a double (or triple) header
	//                     including seperate admission doubleheaders
	//   "3" -- the third game of a triple-header
	//   "A" -- the first game of a double-header involving 3 teams
	//   "B" -- the second game of a double-header involving 3 teams
	"DayOfWeek",         //            Day of week  ("Sun","Mon","Tue","Wed","Thu","Fri","Sat")
	"VisitorTeam",       //     4     Visiting team and league
	"VisitorLeague",     //     5
	"VisitorGameNumber", //     6     Visiting team game number
	//           For this and the home team game number, ties are counted as
	//           games and suspended games are counted from the starting
	//           rather than the ending date.
	"HomeTeam",       //     7     Home team and league
	"HomeLeague",     //     8
	"HomeGameNumber", //     9     Home team game number
	"VisitorScore",   //    10     Visiting and home team score (unquoted)
	"HomeScore",      //    11
	"LengthInOuts",   //    12     Length of game in outs (unquoted).  A full 9-inning game would
	//           have a 54 in this field.  If the home team won without batting
	//           in the bottom of the ninth, this field would contain a 51.
	"DayOrNight", //    13     Day/night indicator ("D" or"N")
	"Completion", //    14     Completion information.  If the game was completed at a
	//           later date (either due to a suspension or an upheld protest)
	//           this field will include:
	//   "yyyymmdd,park,vs,hs,len" Where
	//           yyyymmdd -- the date the game was completed
	//           park -- the park ID where the game was completed
	//           vs -- the visitor score at the time of interruption
	//           hs -- the home score at the time of interruption
	//           len -- the length of the game in outs at time of interruption
	//           All the rest of the information in the record refers to the
	//           entire game.
	"Forfeit", //    15     Forfeit information:
	//   "V" -- the game was forfeited to the visiting team
	//   "H" -- the game was forfeited to the home team
	//   "T" -- the game was ruled a no-decision
	"Protest", //    16     Protest information:
	//   "P" -- the game was protested by an unidentified team
	//   "V" -- a disallowed protest was made by the visiting team
	//   "H" -- a disallowed protest was made by the home team
	//   "X" -- an upheld protest was made by the visiting team
	//   "Y" -- an upheld protest was made by the home team
	//           Note: two of these last four codes can appear in the field
	//           (if both teams protested the game).
	"ParkID",      //    17     Park ID
	"Attendance",  //    18     Attendance (unquoted)
	"Duration",    //    19     Time of game in minutes (unquoted)
	"VisitorLine", //    20       Visiting and home line scores.  For example:
	"HomeLine",    //    21
	//   "010000(10)0x"
	//           Would indicate a game where the home team scored a run in
	//           the second inning, ten in the seventh and didn't bat in the
	//           bottom of the ninth.
	// VISITOR STATS
	// 22-38    Visiting team offensive statistics (unquoted) (in order):
	"VisitorAtBats",  // 22             at-bats
	"VisitorHits",    // 23             hits
	"VisitorDbls",    // 24             doubles
	"VisitorTriples", // 25             triples
	"VisitorHomers",  // 26             homeruns
	"VisitorRBI",     // 27             RBI
	"VisitorSacHits", // 29             sacrifice hits.  This may include sacrifice flies for years
	//                 prior to 1954 when sacrifice flies were allowed.
	"VisitorSacFlies",         // 29             sacrifice flies (since 1954)
	"VisitorHitByPitch",       // 30             hit-by-pitch
	"VisitorWalks",            // 31             walks
	"VisitorIntentionalWalks", // 32             intentional walks
	"VisitorStrikeouts",       // 33             strikeouts
	"VisitorSteals",           // 34             stolen bases
	"VisitorStealCaught",      // 35             caught stealing
	"VisitorGrounded",         // 36             grounded into double plays
	"VisitorCatcherIfc",       // 37             awarded first on catcher's interference
	"VisitorLeftOnBase",       // 39             left on base
	// 39-43     Visiting team pitching statistics (unquoted)(in order):
	"VisitorPitchers", // 39            pitchers used ( 1 means it was a complete game )
	"VisitorIndiERs",  // 40             individual earned runs
	"VisitorTeamERs",  // 41             team earned runs
	"VisitorWilds",    //  42            wild pitches
	"VisitorBalks",    // 43             balks
	// 44-49    Visiting team defensive statistics (unquoted) (in order):
	"VisitorPutouts", // 44            putouts.  Note: prior to 1931, this may not equal 3 times
	//                 the number of innings pitched.  Prior to that, no
	//                 putout was awarded when a runner was declared out for
	//                 being hit by a batted ball.
	"VisitorAssists",     //  45           assists
	"VisitorErrors",      //  46            errors
	"VisitorPassedBalls", //  47            passed balls
	"VisitorDoublePlays", //  48            double plays
	"VisitorTriplePlays", //  49            triple plays

	// HOME STATS
	// 50-66    Home team offensive statistics
	// 67-71    Home team pitching statistics
	// 72-77    Home team defensive statistics
	// 78-79    Home plate umpire ID and name
	"HomeAtBats",  // 50             at-bats
	"HomeHits",    // 51            hits
	"HomeDbls",    // 52             doubles
	"HomeTriples", // 53            triples
	"HomeHomers",  // 54             homeruns
	"HomeRBI",     // 55             RBI
	"HomeSacHits", // 56             sacrifice hits.  This may include sacrifice flies for years
	//                 prior to 1954 when sacrifice flies were allowed.
	"HomeSacFlies",         // 57             sacrifice flies (since 1954)
	"HomeHitByPitch",       // 58             hit-by-pitch
	"HomeWalks",            // 59             walks
	"HomeIntentionalWalks", // 60             intentional walks
	"HomeStrikeouts",       // 61             strikeouts
	"HomeSteals",           // 62             stolen bases
	"HomeStealCaught",      // 63             caught stealing
	"HomeGrounded",         // 64             grounded into double plays
	"HomeCatcherIfc",       // 65             awarded first on catcher's interference
	"HomeLeftOnBase",       // 66             left on base
	//          home team pitching statistics (unquoted)(in order):
	"HomePitchers", // 67            pitchers used ( 1 means it was a complete game )
	"HomeIndiERs",  // 69             individual earned runs
	"HomeTeamERs",  // 69             team earned runs
	"HomeWilds",    // 70            wild pitches
	"HomeBalks",    // 71             balks
	//          home team defensive statistics (unquoted) (in order):
	"HomePutouts", // 72            putouts.  Note: prior to 1931, this may not equal 3 times
	//                 the number of innings pitched.  Prior to that, no
	//                 putout was awarded when a runner was declared out for
	//                 being hit by a batted ball.
	"HomeAssists",     //  73           assists
	"HomeErrors",      //  74            errors
	"HomePassedBalls", //  75            passed balls
	"HomedoublePlays", //  76            double plays
	"HomeTriplePlays", //  77            triple plays

	// UMPIRES
	"UmpireHomeID",      // 78
	"UmpireHomeName",    // 79
	"UmpFirstBaseID",    // 80       1B umpire ID and name
	"UmpFirstBaseName",  // 81
	"UmpSecondBaseID",   // 82       2B umpire ID and name
	"UmpSecondBaseName", // 83
	"UmpThirdBaseID",    // 84       3B umpire ID and name
	"UmpThirdBaseName",  // 85
	"UmpLeftFieldID",    // 86        LF umpire ID and name
	"UmpLeftFieldName",  // 86
	"UmpRightFieldID",   // 88        RF umpire ID and name
	"UmpRightFieldName", // 89
	//          If any umpire positions were not filled for a particular game
	//           the fields will be"","(none)".
	// MANAGERS
	"VisitorManagerID",   // 90        Visiting team manager ID and name
	"VisitorManagerName", // 91
	"HomeManagerID",      // 92        Home team manager ID and name
	"HomeMangerName",     // 93
	// PITCHERS
	"WinningPitcherID",   // 94        Winning pitcher ID and name
	"WinningPitcherName", // 95
	"LosingPitcherID",    // 96        Losing pitcher ID and name
	"LosingPitcherName",  // 97
	"SavingPitcherID",    // 98        Saving pitcher ID and name--"","(none)" if none awarded
	"SavingPitcherName",  // 99
	"RbiBatterID",        // 100       Game Winning RBI batter ID and name--"","(none)" if none
	"RbiBatterName",      // 101
	"VisitorStarterID",   // 102       Visiting starting pitcher ID and name
	"VisitorStarterName", // 103
	"HomeStarterID",      // 104       Home starting pitcher ID and name
	"HomeStarterName",    // 105

	// 106-132   Visiting starting players ID, name and defensive position,
	//           listed in the order (1-9) they appeared in the batting order.
	"VisitorLineup1ID",   // 106
	"VisitorLineup1Name", // 107
	"VisitorLineup1Pos",  // 108
	"VisitorLineup2ID",   // 109
	"VisitorLineup2Name", // 110
	"VisitorLineup2Pos",  // 111
	"VisitorLineup3ID",   // 112
	"VisitorLineup3Name", // 113
	"VisitorLineup3Pos",  // 114
	"VisitorLineup4ID",   // 115
	"VisitorLineup4Name", // 116
	"VisitorLineup4Pos",  // 117
	"VisitorLineup5ID",   // 118
	"VisitorLineup5Name", // 119
	"VisitorLineup5Pos",  // 120
	"VisitorLineup6ID",   // 121
	"VisitorLineup6Name", // 122
	"VisitorLineup6Pos",  // 123
	"VisitorLineup7ID",   // 124
	"VisitorLineup7Name", // 125
	"VisitorLineup7Pos",  // 126
	"VisitorLineup8ID",   // 127
	"VisitorLineup8Name", // 128
	"VisitorLineup8Pos",  // 129
	"VisitorLineup9ID",   // 130
	"VisitorLineup9Name", // 131
	"VisitorLineup9Pos",  // 132

	// 133-159   Home starting players ID, name and defensive position
	//           listed in the order (1-9) they appeared in the batting order.
	"HomeLineup1ID",   // 133
	"HomeLineup1Name", // 134
	"HomeLineup1Pos",  // 135
	"HomeLineup2ID",   // 136
	"HomeLineup2Name", // 137
	"HomeLineup2Pos",  // 138
	"HomeLineup3ID",   // 139
	"HomeLineup3Name", // 140
	"HomeLineup3Pos",  // 141
	"HomeLineup4ID",   // 142
	"HomeLineup4Name", // 143
	"HomeLineup4Pos",  // 144
	"HomeLineup5ID",   // 145
	"HomeLineup5Name", // 146
	"HomeLineup5Pos",  // 147
	"HomeLineup6ID",   // 148
	"HomeLineup6Name", // 149
	"HomeLineup6Pos",  // 150
	"HomeLineup7ID",   // 151
	"HomeLineup7Name", // 152
	"HomeLineup7Pos",  // 153
	"HomeLineup8ID",   // 154
	"HomeLineup8Name", // 155
	"HomeLineup8Pos",  // 156
	"HomeLineup9ID",   // 157
	"HomeLineup9Name", // 158
	"HomeLineup9Pos",  // 159

	"AdditionalInfo", //   160     Additional information.  This is a grab-bag of informational
	//           items that might not warrant a field on their own.  The field
	//           is alpha-numeric. Some items are represented by tokens such as:
	//   "HTBF" -- home team batted first.
	//              Note: if"HTBF" is specified it would be possible to see
	//              something like"01002000x" in the visitor's line score.
	//           Changes in umpire positions during a game will also appear in
	//           this field.  These will be in the form:
	//              umpchange,inning,umpPosition,umpid with the latter three
	//              repeated for each umpire.
	//           These changes occur with umpire injuries, late arrival of
	//           umpires or changes from completion of suspended games. Details
	//           of suspended games are in field 14.

	"AcquisitionInfo", //   161     Acquisition information:
	//   "Y" -- we have the complete game
	//   "N" -- we don't have any portion of the game
	//   "D" -- the game was derived from box score and game story
	//   "P" -- we have some portion of the game.  We may be missing
	//                     innings at the beginning, middle and end of the game.

	// Missing fields will be NULL.
	//
}

// matching types for each of those fields
var gltypes = []string{
	"string", "string", "string", "string", "string", "int", "string", "string", "int", "int", "int", "int", "string", "string", "string", "string", "string",
	"int", "int", "string", "string", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int",
	"int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int",
	"int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int", "int",
	"int", "int", "int", "int", "int", "int", "int", "int", "int", "string", "string", "string", "string", "string", "string", "string", "string",
	"string", "string", "string", "string", "string", "string", "string", "string", "string", "string", "string", "string", "string", "string", "string", "string", "string",
	"string", "string", "string", "string", "string", "int", "string", "string", "int", "string", "string", "int", "string", "string", "int", "string", "string",
	"int", "string", "string", "int", "string", "string", "int", "string", "string", "int", "string", "string", "int", "string", "string", "int", "string",
	"string", "int", "string", "string", "int", "string", "string", "int", "string", "string", "int", "string", "string", "int", "string", "string", "int",
	"string", "string", "string", "string", "string", "string", "string", "string",
}

func makeGameStruct() {
	fmt.Println("type game struct {")
	for i := 0; i < len(glfields); i++ {
		fmt.Printf("\t%-24s %s `json:\",omitempty\"`\n", glfields[i], gltypes[i])
	}
	fmt.Println("}")
}
